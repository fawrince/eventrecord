package http

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/fawrince/eventrecord/application"
	"github.com/fawrince/eventrecord/dto"
	"github.com/fawrince/eventrecord/logger"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"io/ioutil"
	"net/http"
	"time"
)

type ServerNetHttp struct {
	logger      *logger.Logger
	server      *http.Server
	application *application.App
}

func NewServer(logger *logger.Logger, app *application.App) *ServerNetHttp {
	router := mux.NewRouter()

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", Port),
		Handler: router,
	}

	return &ServerNetHttp{
		logger:      logger,
		server:      srv,
		application: app,
	}
}

func (server *ServerNetHttp) Start() {
	server.logger.Infof("Starting the http server at address: %s...", server.server.Addr)

	server.mapHandlers(server.server.Handler.(*mux.Router))

	go func() {
		err := server.server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			server.logger.Fatal(fmt.Errorf("couldnt start the http server: %w", err))
		}
	}()

	server.logger.Infof("Server started")
}

func (server *ServerNetHttp) Stop() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer func() {
		// extra handling here
		cancel()
	}()

	if err := server.server.Shutdown(ctx); err != nil {
		server.logger.Fatal(fmt.Errorf("server shutdown failed: %w", err))
	}
	server.logger.Infof("Server stopped")
}

func (server *ServerNetHttp) mapHandlers(router *mux.Router) {
	router.HandleFunc("/", server.indexHandler)
	router.HandleFunc("/send", server.pushCoordinatesHandler).Methods("POST")
	router.HandleFunc("/recv", server.pullCoordinatesHandler).Methods("GET")

	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

	// Prometheus endpoint
	router.Path("/prometheus").Handler(promhttp.Handler())

	router.Use(server.buildMiddleware())
	//router.Use(monitor.PrometheusMiddleware)
}

func (server *ServerNetHttp) buildMiddleware() mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			server.logger.Tracef("Request intercepted: %s", r.RequestURI)
			next.ServeHTTP(w, r)
		})
	}
}

func (server *ServerNetHttp) indexHandler(w http.ResponseWriter, r *http.Request) {
	fileBytes, err := ioutil.ReadFile("static/index.html")
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/html")
	w.Write(fileBytes)
}

// pushCoordinatesHandler receives a new coordinates data from the client and produces a message to the broker
func (server *ServerNetHttp) pushCoordinatesHandler(w http.ResponseWriter, r *http.Request) {
	var coordinates dto.Coordinates
	err := json.NewDecoder(r.Body).Decode(&coordinates)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	server.application.Produce.ProduceInput() <- coordinates

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("{ \"message\": \"ok\"}"))
}

// pullCoordinatesHandler sends consumed coordinates over long living server-sent-events connection
func (server *ServerNetHttp) pullCoordinatesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	for {
		select {
		case <-r.Context().Done():
			server.logger.Infof("SSE connection closed")
			return
		case coord := <-server.application.Consume.ConsumeOutput():
			time.Sleep(time.Millisecond * 10)
			var buf bytes.Buffer
			enc := json.NewEncoder(&buf)
			enc.Encode(coord)
			fmt.Fprintf(w, "data: %v\n\n", buf.String())
			if f, ok := w.(http.Flusher); ok {
				f.Flush()
			}
			server.logger.Infof("Data sent to SSE response: %v", coord)
		}
	}
}

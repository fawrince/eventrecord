package http

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/fawrince/eventrecord/application"
	"github.com/fawrince/eventrecord/broker"
	logger "github.com/fawrince/eventrecord/logger"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"time"
)

type Server struct {
	logger      *logger.Logger
	server      *http.Server
	application *application.App
}

func NewServer(logger *logger.Logger, app *application.App) *Server {
	router := mux.NewRouter()

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", Port),
		Handler: router,
	}

	return &Server{
		logger:      logger,
		server:      srv,
		application: app,
	}
}

func (server *Server) Start() {
	server.logger.Infof("Starting the http server at address: %s...", server.server.Addr)

	server.mapHandlers(server.server.Handler.(*mux.Router))

	go func() {
		err := server.server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			server.logger.Fatal("goerr", err, fmt.Sprintf("Couldnt start the http server: %s", err))
		}
	}()

	server.logger.Infof("Server started")
}

func (server *Server) Stop() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer func() {
		// extra handling here
		cancel()
	}()

	if err := server.server.Shutdown(ctx); err != nil {
		server.logger.Fatal("goerr", err, "Server shutdown failed:%+v", err)
	}
	server.logger.Infof("Server stopped")
}

func (server *Server) mapHandlers(router *mux.Router) {
	router.HandleFunc("/", server.indexHandler)
	router.HandleFunc("/send", server.pushCoordinatesHandler).Methods("POST")
	router.HandleFunc("/recv", server.pullCoordinatesHandler).Methods("GET")
	router.PathPrefix("/static").Handler(http.FileServer(http.Dir("./static")))
	router.Use(server.buildMiddleware())
}

func (server *Server) buildMiddleware() mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			server.logger.Trace("url", r.RequestURI, "Request intercepted")
			next.ServeHTTP(w, r)
		})
	}
}

func (server *Server) indexHandler(w http.ResponseWriter, r *http.Request) {
	fileBytes, err := ioutil.ReadFile("static/index.html")
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/html")
	w.Write(fileBytes)
}

// pushCoordinatesHandler receives a new coordinates data from the client and produces a message to the broker
func (server *Server) pushCoordinatesHandler(w http.ResponseWriter, r *http.Request) {
	var coordinates broker.Coordinates
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
func (server *Server) pullCoordinatesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	for {
		select {
		case coord := <-server.application.Consume.ConsumeOutput():
			var buf bytes.Buffer
			enc := json.NewEncoder(&buf)
			enc.Encode(coord)
			fmt.Fprintf(w, "data: %v\n\n", buf.String())
			if f, ok := w.(http.Flusher); ok {
				f.Flush()
			}
			server.logger.Infof("Data sent to sse: %v (%v)", coord, buf)
		}
	}
}

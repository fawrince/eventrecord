package http

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/ansrivas/fiberprometheus/v2"
	"github.com/fawrince/eventrecord/application"
	"github.com/fawrince/eventrecord/dto"
	"github.com/fawrince/eventrecord/logger"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"time"
)

type Server struct {
	logger      *logger.Logger
	server      *fiber.App
	application *application.App
}

func NewServer2(logger *logger.Logger, app *application.App) *Server {
	server := fiber.New()

	return &Server{
		logger:      logger,
		server:      server,
		application: app,
	}
}

func (server *Server) Start() {
	server.logger.Infof("Starting the http server at address: :%s...", Port)

	server.mapHandlers()

	go func() {
		err := server.server.Listen(fmt.Sprintf(":%s", Port))
		if err != nil {
			server.logger.Fatal(err)
		}
	}()

	server.logger.Infof("Server started")
}

func (server *Server) Stop() {
	server.server.Shutdown()
}

func (server *Server) mapHandlers() {
	srv := server.server

	srv.Use(func(c *fiber.Ctx) error {
		server.logger.Tracef("Request intercepted: %s", c.OriginalURL())
		return c.Next()
	})

	srv.Use("/ws", func(c *fiber.Ctx) error {
		// IsWebSocketUpgrade returns true if the client
		// requested upgrade to the WebSocket protocol.
		if websocket.IsWebSocketUpgrade(c) {
			server.logger.Infof("WebSocket upgrade requested")
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})

	server.registerPrometheusHandler()

	srv.Static("/", "./static/index.html")
	srv.Get("/ws/:client", server.buildPushCoordinatesSocketHandler())
	srv.Get("/recv", server.buildPullCoordinatesHandler())
}

// buildPushCoordinatesSocketHandler receives the coordinates from the client via WebSocket-connection and produces a message to the broker.
func (server *Server) buildPushCoordinatesSocketHandler() func(c *fiber.Ctx) error {
	return websocket.New(func(c *websocket.Conn) {
		defer func() {
			c.Close()
		}()

		var (
			mt  int
			msg []byte
			err error
		)
		for {
			if mt, msg, err = c.ReadMessage(); err != nil {
				if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
					server.logger.Error(err)
				}
				return // Calls the deferred function, i.e. closes the connection on error
			}

			if mt == websocket.TextMessage {
				var coordinates dto.Coordinates
				if err := json.Unmarshal(msg, &coordinates); err != nil {
					server.logger.Error(err)
				}
				server.application.Produce.ProduceInput() <- coordinates
			} else {
				server.logger.Infof("WebSocket message received of type %v", mt)
			}
		}
	})
}

// buildPullCoordinatesHandler consumes events from the broker and sends it the client via SSE.
func (server *Server) buildPullCoordinatesHandler() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		c.Set("Content-Type", "text/event-stream")
		c.Set("Cache-Control", "no-cache")
		c.Set("Connection", "keep-alive")
		c.Set("Transfer-Encoding", "chunked")

		doneCh := c.Context().Done()

		c.Context().SetBodyStreamWriter(func(w *bufio.Writer) {
			for {
				select {
				case <-doneCh:
					server.logger.Infof("Sse connection closed...")
					return

				case coord := <-server.application.Consume.ConsumeOutput():
					time.Sleep(time.Millisecond * 10)
					var buf bytes.Buffer
					enc := json.NewEncoder(&buf)
					enc.Encode(coord)
					fmt.Fprintf(w, "data: %v\n\n", buf.String())
					if err := w.Flush(); err != nil {
						server.logger.Infof("Sse connection closed (err)...")
						return
					}
				}
			}
		})

		return nil
	}
}

func (server *Server) registerPrometheusHandler() {
	prometheus := fiberprometheus.New("event-record")
	prometheus.RegisterAt(server.server, "/prometheus")
	server.server.Use(prometheus.Middleware)
}

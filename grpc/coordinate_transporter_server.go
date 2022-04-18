// Package grpc contains an implementation of grpc-tool generated server interface
package grpc

import (
	"fmt"
	"github.com/fawrince/eventrecord/dto"
	"github.com/fawrince/eventrecord/logger"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
)

type ProduceHandler interface {
	ProduceInput() chan<- dto.Coordinates
}

type coordinateTransporterServer struct {
	UnimplementedCoordinateTransporterServer
	producer   ProduceHandler
	logger     *logger.Logger
	grpcServer *grpc.Server
}

func NewCoordinateTransporterServer(producer ProduceHandler, logger *logger.Logger) *coordinateTransporterServer {
	return &coordinateTransporterServer{
		producer: producer,
		logger:   logger,
	}
}

func (server *coordinateTransporterServer) Start() {
	server.logger.Infof("Start the grpc server at address %s...", Port)

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%s", Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	var opts []grpc.ServerOption
	server.grpcServer = grpc.NewServer(opts...)
	RegisterCoordinateTransporterServer(server.grpcServer, server)
	go func() {
		err := server.grpcServer.Serve(lis)
		if err != nil {
			server.logger.Fatal("goerr", err, fmt.Sprintf("Couldnt start the grpc server: %s", err))
		}
	}()
}

func (server *coordinateTransporterServer) Stop() {
	server.grpcServer.Stop()
}

func (server *coordinateTransporterServer) PostCoordinates(stream CoordinateTransporter_PostCoordinatesServer) error {
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		coordinates := dto.Coordinates{
			X:      int64(in.X),
			Y:      int64(in.Y),
			Client: in.ClientId,
		}

		server.producer.ProduceInput() <- coordinates

		if err := stream.Send(&PostCoordinateResponse{Ok: true}); err != nil {
			return err
		}
	}
}

package grpc

import "flag"

var (
	Port = ""
)

func init() {
	flag.StringVar(&Port, "grpcport", "2001", "Grpc server port")
}

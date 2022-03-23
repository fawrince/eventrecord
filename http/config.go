package http

import "flag"

// Http server configuration option
var (
	Port = ""
)

func init() {
	flag.StringVar(&Port, "httpport", "1001", "Http server port")
}

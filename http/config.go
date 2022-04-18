package http

import (
	"flag"
)

var (
	Port = ""
)

func init() {
	flag.StringVar(&Port, "httpport", "1001", "Http server port")
}

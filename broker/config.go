package broker

import "flag"

// Broker configuration options
var (
	Brokers = ""
	Version = ""
	Group   = ""
	Topics  = ""
	Oldest  = true
	Verbose = false
)

func init() {
	flag.StringVar(&Brokers, "brokers", "broker:9092", "Kafka bootstrap brokers to connect to, as a comma separated list")
	flag.StringVar(&Group, "group", "g1", "Kafka consumer group definition")
	flag.StringVar(&Version, "version", "2.8.1", "Kafka cluster version")
	flag.StringVar(&Topics, "topics", "coordinates", "Kafka topics to be consumed, as a comma separated list")
	flag.BoolVar(&Oldest, "oldest", true, "Kafka consumer consume initial offset from oldest")
	flag.BoolVar(&Verbose, "verbose", false, "Sarama logging")
}

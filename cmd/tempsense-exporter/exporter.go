package main

import (
	"flag"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
)

var (
	flagAddr = flag.String("address", ":9181", "The address to listen on for HTTP requests.")
)

func main() {
	flag.Parse()

	prometheus.MustRegister(NewTempsenseCollector())
	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(*flagAddr, nil))
}

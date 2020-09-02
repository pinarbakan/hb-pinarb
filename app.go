package main

import (
"fmt"
"log"
"net/http"
"time"
"github.com/prometheus/client_golang/prometheus"
"github.com/prometheus/client_golang/prometheus/promauto"
"github.com/prometheus/client_golang/prometheus/promhttp"
)
func recordMetrics() {
go func() {
for {
opsProcessed.Inc()
time.Sleep(2 * time.Second)
}
}()
}

var (
opsProcessed = promauto.NewCounter(prometheus.CounterOpts{
Name: "myapp_processed_ops_total",
Help: "The total number of processed events",
})
)

func main() {
recordMetrics()

 http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
fmt.Fprintf(w, "Hello Hepsiburada from @pinarbakan")
})

 http.Handle("/metrics", promhttp.Handler())
log.Fatal(http.ListenAndServe(":11130", nil))

}
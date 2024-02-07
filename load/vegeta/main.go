package main

import (
	"fmt"
	"loadtest/config"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	vegeta "github.com/tsenart/vegeta/v12/lib"
	"github.com/tsenart/vegeta/v12/lib/prom"
)

func main() {

	cfg := config.NewConfig()

	name := cfg.TestName
	rate := vegeta.Rate{Freq: cfg.Freq, Per: time.Second}
	duration := time.Duration(cfg.Duration) * time.Second
	targeter := vegeta.NewStaticTargeter(vegeta.Target{
		Method: cfg.Method,
		URL:    cfg.URL,
	})
	promAddr := cfg.PromAddr

	var pm *prom.Metrics
	if promAddr != "" {
		pm = prom.NewMetrics()

		r := prometheus.NewRegistry()
		if err := pm.Register(r); err != nil {
			fmt.Errorf("error registering prometheus metrics: %s", err)
		}

		srv := http.Server{
			Addr:    promAddr,
			Handler: prom.NewHandler(r, time.Now().UTC()),
		}

		defer srv.Close()
		go srv.ListenAndServe()
	}

	attacker := vegeta.NewAttacker()

	var metrics vegeta.Metrics
	for res := range attacker.Attack(targeter, rate, duration, name) {
		metrics.Add(res)
		if res.Error != "" {
			fmt.Printf("Request error: %s\n", res.Error)
		}
		fmt.Printf("Request status code: %d\n", res.Code)
	}
	defer metrics.Close()

	fmt.Printf("Total Requests: %d\n", metrics.Requests)
	fmt.Printf("Success Ratio: %.2f%%\n", metrics.Success*100)
	fmt.Printf("Status Codes: %+v\n", metrics.StatusCodes)
	fmt.Printf("Mean Latency: %s\n", metrics.Latencies.Mean)
	fmt.Printf("50th Percentile Latency: %s\n", metrics.Latencies.P50)
	fmt.Printf("95th Percentile Latency: %s\n", metrics.Latencies.P95)
	fmt.Printf("99th percentile: %s\n", metrics.Latencies.P99)
}

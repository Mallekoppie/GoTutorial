package main

import (
	"fmt"
	"time"

	vegeta "github.com/tsenart/vegeta/lib"
)

func main() {
	rate := vegeta.Rate{Freq: 1, Per: time.Second}

	duration := 4 * time.Second
	targeter := vegeta.NewStaticTargeter(vegeta.Target{
		Method: "GET",
		URL:    "http://localhost:80/ConnectionTest",
	},
		vegeta.Target{
			Method: "GET",
			URL:    "http://localhost:80/",
		})

	attacker := vegeta.NewAttacker()
	//vegeta.DefaultTLSConfig.InsecureSkipVerify = true
	var metrics vegeta.Metrics
	for res := range attacker.Attack(targeter, rate, duration, "Big Bang!") {
		metrics.Add(res)
	}
	metrics.Close()

	fmt.Printf("99th percentile: %s\n", metrics.Latencies.P99)
}

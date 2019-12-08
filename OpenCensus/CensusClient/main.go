package main

import (
	"context"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	//	openzipkin "github.com/openzipkin/zipkin-go"
	//	zipkinHTTP "github.com/openzipkin/zipkin-go/reporter/http"

	"contrib.go.opencensus.io/exporter/prometheus"
	//	"contrib.go.opencensus.io/exporter/zipkin"
	"go.opencensus.io/plugin/ochttp"
	"go.opencensus.io/stats/view"
)

// Reference: https://opencensus.io/guides/http/go/net_http/client/

func main() {
	// Firstly, we'll register ochttp Client views
	if err := view.Register(ochttp.DefaultClientViews...); err != nil {
		log.Fatalf("Failed to register client views for HTTP metrics: %v", err)
	}

	// Enable observability to extract and examine traces and metrics.
	enableObservabilityAndExporters()

	// Create our HTTP client that uses the ochttp.Transport.
	client := &http.Client{Transport: &ochttp.Transport{}}
	i := uint64(0)

	// Then finally do the work every 5 seconds.
	for {
		i += 1
		log.Printf("Performing fetch #%d", i)
		doWork(context.Background(), client)

		<-time.After(500 * time.Millisecond)
	}
}

func doWork(ctx context.Context, client *http.Client) {
	req, _ := http.NewRequest("GET", "http://localhost:85/ConnectionTest", nil)

	// It is imperative that req.WithContext is used to
	// propagate context and use it in the request.
	//req = req.WithContext(ctx)

	// Now make the request to the remote end.
	res, err := client.Do(req)
	if err != nil {
		log.Printf("Failed to make the request: %v", err)
		return
	}

	// Consume the body and close it.
	io.Copy(ioutil.Discard, res.Body)
	_ = res.Body.Close()

}

func enableObservabilityAndExporters() {
	// Stats exporter: Prometheus
	pe, err := prometheus.NewExporter(prometheus.Options{
		Namespace: "ochttp_tutorial",
	})
	if err != nil {
		log.Fatalf("Failed to create the Prometheus stats exporter: %v", err)
	}

	go func() {
		mux := http.NewServeMux()
		mux.Handle("/metrics", pe)
		log.Fatal(http.ListenAndServe(":8888", mux))
	}()

}

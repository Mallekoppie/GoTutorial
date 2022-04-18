package main

import (
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"strings"
	"time"

	"contrib.go.opencensus.io/exporter/prometheus"
	"go.opencensus.io/plugin/ochttp"
	"go.opencensus.io/stats/view"
)

func main() {
	if err := view.Register(ochttp.DefaultServerViews...); err != nil {
		log.Fatalf("Failed to register server views for HTTP metrics: %v", err)
	}

	// Enable observability to extract and examine stats.
	enableObservabilityAndExporters()

	// The handler containing your business logic to process requests.
	originalHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Consume the request's body entirely.
		io.Copy(ioutil.Discard, r.Body)

		// Generate some payload of random length.
		res := strings.Repeat("a", rand.Intn(99971)+1)

		// Sleep for a random time to simulate a real server's operation.
		time.Sleep(time.Duration(rand.Intn(977)+1) * time.Millisecond)

		// Finally write the body to the response.
		w.Write([]byte("Hello, World! " + res))
	})
	och := &ochttp.Handler{
		Handler: originalHandler, // The handler you'd have used originally
	}
	cst := httptest.NewServer(och)
	defer cst.Close()

	client := &http.Client{}
	for {
		body := strings.NewReader(strings.Repeat("a", rand.Intn(777)+1))
		req, _ := http.NewRequest("POST", cst.URL, body)
		res, _ := client.Do(req)
		io.Copy(ioutil.Discard, res.Body)
		res.Body.Close()
		time.Sleep(979 * time.Millisecond)
	}
}

func enableObservabilityAndExporters() {
	// Stats exporter: Prometheus
	pe, err := prometheus.NewExporter(prometheus.Options{
		Namespace: "ochttp_tutorial",
	})
	if err != nil {
		log.Fatalf("Failed to create the Prometheus stats exporter: %v", err)
	}

	view.RegisterExporter(pe)
	go func() {
		mux := http.NewServeMux()
		mux.Handle("/metrics", pe)
		log.Fatal(http.ListenAndServe(":8888", mux))
	}()

	// Trace exporter: Zipkin
	//localEndpoint, err := openzipkin.NewEndpoint("ochttp_tutorial", "localhost:5454")
	//if err != nil {
	//	log.Fatalf("Failed to create the local zipkinEndpoint: %v", err)
	//}
	//reporter := zipkinHTTP.NewReporter("http://localhost:9411/api/v2/spans")
	//ze := zipkin.NewExporter(reporter, localEndpoint)
	//trace.RegisterExporter(ze)
	//trace.ApplyConfig(trace.Config{DefaultSampler: trace.AlwaysSample()})
}

package main

import (
	"fmt"
	"log"
	"net/http"

	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	// Create a span for a web request at the /user1 URL.
	span := tracer.StartSpan("web.request", tracer.ResourceName("/user1"))
	defer span.Finish()

	// Set a custom tag
	span.SetTag("user", "Vamshi")

	// Send a custom metric
	tracer.SetMetric(span, "custom.metric", 1)

	fmt.Fprint(w, "Hello, Vamshi")
	log.Printf("my log message %v User: Vamshi", span)
}

func customResponseHandler(w http.ResponseWriter, r *http.Request) {
	// Create a span for a web request at the /user2 URL.
	span := tracer.StartSpan("web.request", tracer.ResourceName("/user2"))
	defer span.Finish()

	// Set a custom tag
	span.SetTag("user", "Tara")

	// Send a custom metric
	tracer.SetMetric(span, "custom.metric", 1)

	fmt.Fprint(w, "Hello, Tara")
	log.Printf("my log message %v User: Tara", span)
}

func handler(w http.ResponseWriter, r *http.Request) {
	// Create a span for a web request at the /posts URL.
	span := tracer.StartSpan("web.request", tracer.ResourceName("/posts"))
	defer span.Finish()

	// Set a custom metric
	tracer.SetMetric(span, "custom.metric", 1)

	// Append span info to log messages:
	fmt.Fprint(w, "Hello, this is for tracing")
	log.Printf("my log message %v", span)
}

func main() {
	// Configure the Datadog tracer to communicate with the DataDog Agent in Kubernetes
	tracer.Start(tracer.WithAgentAddr("datadog-agent-service-name:8126"), tracer.WithRuntimeMetrics())

	// Register your handlers
	http.HandleFunc("/user1", helloWorldHandler)
	http.HandleFunc("/user2", customResponseHandler)
	http.HandleFunc("/posts", handler)

	port := 8080
	fmt.Printf("Server is running on port %d...\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		fmt.Println("Error:", err)
	}
}

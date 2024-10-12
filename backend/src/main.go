package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
)

type Response struct {
	Data    string `json:"data"`
	HTTPMux string `json:"http_mux"`
}

func main() {
	stdMux := http.NewServeMux()
	stdMux.HandleFunc("/api/v1/healthz", func(w http.ResponseWriter, r *http.Request) {
		resp := Response{Data: "healthz", HTTPMux: "net/http"}
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			fmt.Printf(err.Error())
		}
	})

	stdMux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		resp := Response{Data: "hello world!", HTTPMux: "net/http"}
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			fmt.Printf(err.Error())
		}
	})

	l, ok := os.LookupEnv("API_CONTAINER_PORT")

	if !ok {
		fmt.Print("No specified port number")
		return
	}

	listeningPortNumber, err := strconv.Atoi(l)

	if err != nil {
		fmt.Print("Invalid port number")
		return
	}

	server := http.Server{
		Addr:    fmt.Sprintf(":%d", listeningPortNumber),
		Handler: stdMux,
	}

	fmt.Printf("Listening on %s...\n", server.Addr)
	_ = server.ListenAndServe()

	shutdown := make(chan os.Signal, 1)
	fmt.Println("Press Ctrl+C to stop.")
	signal.Notify(shutdown, syscall.SIGTERM, syscall.SIGINT)
	<-shutdown
}

package main

import (
	"net/http"
	"time"

	"github.com/JanStanleyWatt/hukubiki-api/api"
)

func main() {
	server := http.Server{
		Addr:              ":8000",
		ReadTimeout:       500 * time.Millisecond,
		ReadHeaderTimeout: 500 * time.Millisecond,
		Handler:           http.TimeoutHandler(http.DefaultServeMux, 500*time.Millisecond, "Timeout"),
	}

	http.HandleFunc("/int64/n", api.Int64n)
	http.HandleFunc("/", api.Usage)

	server.ListenAndServe()
}

package main

import (
	"net/http"

	"github.com/JanStanleyWatt/hukubiki-api/api"
)

func main() {
	server := http.Server{
		Addr:    ":8000",
		Handler: nil,
	}

	http.HandleFunc("/int64/n", api.Int64n)
	http.HandleFunc("/", api.Usage)

	server.ListenAndServe()
}

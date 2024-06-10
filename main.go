package main

import (
	"net/http"

	"github.com/JanStanleyWatt/hukubiki-api/hukubiki"
)

func main() {
	server := http.Server{
		Addr:    ":8000",
		Handler: nil,
	}

	http.HandleFunc("/int64n", hukubiki.Int64n)
	http.HandleFunc("/", hukubiki.Usage)

	server.ListenAndServe()
}

package main

import (
	"net/http"

	"github.com/JanStanleyWatt/hukubiki-api/api"
	"github.com/JanStanleyWatt/hukubiki-api/route"
)

func main() {
	server := http.Server{
		Addr:    ":8000",
		Handler: nil,
	}

	http.HandleFunc("/int64n", api.Int64n)
	http.HandleFunc("/", route.ResponseUsage)

	server.ListenAndServe()
}

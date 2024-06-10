package main

import "net/http"

type Input struct {
	Seed int64 `json:"seed"`
	Max  int64 `json:"max"`
}

func main() {
	server := http.Server{
		Addr:    ":8000",
		Handler: nil,
	}

	http.HandleFunc("/int64n", func(w http.ResponseWriter, r *http.Request) {})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {})

	server.ListenAndServe()
}

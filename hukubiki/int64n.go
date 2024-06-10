package hukubiki

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
)

type jsonInt64N struct {
	RandomNumber int64 `json:"randomNumber"`
}

type request struct {
	Seed int64 `json:"seed"`
	Max  int64 `json:"max"`
}

func Int64n(w http.ResponseWriter, r *http.Request) {

	// Read json from request
	var in request

	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Generate random number
	j := jsonInt64N{
		RandomNumber: rand.New(rand.NewSource(in.Seed)).Int63n(in.Max),
	}

	res, err := json.Marshal(j)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "%s", string(res))
}

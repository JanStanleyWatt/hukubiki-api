package api

import (
	"encoding/json"
	"fmt"
	"math/rand/v2"
	"net/http"
)

type responseInt64n struct {
	RandomNumber int64 `json:"randomNumber"`
}

type requestInt64n struct {
	Seed [2]uint64 `json:"seed"`
	Max  int64     `json:"max"`
}

func Int64n(w http.ResponseWriter, r *http.Request) {

	// Read json from request
	var in requestInt64n

	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Generate random number
	j := responseInt64n{
		RandomNumber: rand.New(rand.NewPCG(in.Seed[0], in.Seed[1])).Int64N(in.Max),
	}

	res, err := json.Marshal(j)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "%s", string(res))
}

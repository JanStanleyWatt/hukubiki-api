package main

import (
	"encoding/json"
	"math/rand/v2"
)

type Output struct {
	RandomNumber int `json:"randomNumber"`
}

func main() {
	// Seed the random number generator with a random number
	n := rand.IntN(32)

	// Create a new Output struct
	output := Output{
		RandomNumber: n,
	}

	// Create json
	jsonOutput, err := json.Marshal(&output)
	if err != nil {
		println("Error marshalling to JSON:", err)
		return
	}

	// Print the json
	println(string(jsonOutput))
}

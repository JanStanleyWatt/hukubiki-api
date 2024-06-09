package main

import (
	"encoding/json"
	"math/rand"
	"os"
)

type Input struct {
	Seed int64 `json:"seed"`
	Max  int64 `json:"max"`
}

type Output struct {
	RandomNumber int64 `json:"randomNumber"`
}

func main() {
	// Read the json from the command line
	var input Input
	j := os.Args[1]

	// Unmarshal the json
	err := json.Unmarshal([]byte(j), &input)
	if err != nil {
		println("Error unmarshalling JSON:", err.Error())
		return
	}

	// Create a new random number generator
	r := rand.New(rand.NewSource(input.Seed))

	// Create a new Output struct
	output := Output{
		RandomNumber: r.Int63n(input.Max),
	}

	// Create json
	jsonOutput, err := json.Marshal(&output)
	if err != nil {
		println("Error marshalling to JSON:", err.Error())
		return
	}

	// Print the json
	println(string(jsonOutput))
}

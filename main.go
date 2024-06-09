package main

import (
	"encoding/json"
	"math/rand"
	"os"
	"strconv"
)

type Output struct {
	RandomNumber int64 `json:"randomNumber"`
}

func main() {
	// Seed the random number generator with a random number
	seed, err := strconv.Atoi(os.Args[1])
	if err != nil {
		println("Error converting seed to int:", err)
		return
	}

	// Get the max number
	max, err := strconv.Atoi(os.Args[2])
	if err != nil {
		println("Error converting seed to int:", err)
		return
	}

	// Create a new random number generator
	r := rand.New(rand.NewSource(int64(seed)))

	// Create a new Output struct
	output := Output{
		RandomNumber: r.Int63n(int64(max)),
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

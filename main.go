package main

import (
	"encoding/json"
	"os"

	"github.com/JanStanleyWatt/hukubiki-api/hukubiki"
)

type Input struct {
	Seed int64 `json:"seed"`
	Max  int64 `json:"max"`
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

	output := hukubiki.Int64n(input.Seed, input.Max)

	// Create json
	jsonOutput, err := json.Marshal(&output)
	if err != nil {
		println("Error marshalling to JSON:", err.Error())
		return
	}

	// Print the json
	println(string(jsonOutput))
}

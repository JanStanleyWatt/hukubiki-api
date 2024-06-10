package hukubiki

import "math/rand"

type Output struct {
	RandomNumber int64 `json:"randomNumber"`
}

func Int64n(seed, max int64) Output {
	r := rand.New(rand.NewSource(seed))
	return Output{
		RandomNumber: r.Int63n(max),
	}
}

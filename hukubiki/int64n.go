package hukubiki

import "math/rand"

type jsonInt64N struct {
	RandomNumber int64 `json:"randomNumber"`
}

func Int64n(seed, max int64) jsonInt64N {
	r := rand.New(rand.NewSource(seed))
	return jsonInt64N{
		RandomNumber: r.Int63n(max),
	}
}

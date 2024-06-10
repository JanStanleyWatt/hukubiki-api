package hukubiki

import "math/rand"

type JsonInt64N struct {
	RandomNumber int64 `json:"randomNumber"`
}

func Int64n(seed, max int64) JsonInt64N {
	r := rand.New(rand.NewSource(seed))
	return JsonInt64N{
		RandomNumber: r.Int63n(max),
	}
}

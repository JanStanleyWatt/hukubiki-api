package api

import (
	"errors"
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

	// リクエストのボディをデコード、エラーがあればエラーレスポンスをJsonで返す
	var in requestInt64n
	if err := jsonInput(r, &in); err != nil {
		jsonErrorOutput(w, http.StatusBadRequest, err)
		return
	}

	// [0, Max)の範囲で乱数を生成。Maxが0以下の場合はエラーレスポンスを返す
	if in.Max <= 0 {
		err := errors.New("max must be greater than 0")
		jsonErrorOutput(w, http.StatusBadRequest, err)
		return
	}
	out := responseInt64n{
		RandomNumber: rand.New(rand.NewPCG(in.Seed[0], in.Seed[1])).Int64N(in.Max),
	}

	// レスポンスを返す
	jsonOutput(w, out)
}

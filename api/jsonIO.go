package api

import (
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
)

type jsonError struct {
	StatusCode int    `json:"statusCode"`
	Error      string `json:"error"`
}

const CONTENT_TYPE_JSON = "application/json; charset=utf-8"

// Jsonの入力を受け取るヘルパ関数。リクエストのヘッダのうち、Content-Typeがapplication/jsonであることを要求する
func jsonInput(r *http.Request, in interface{}) error {
	// リクエストのヘッダをチェックし、要件に合わなければエラーレスポンスをJsonで返す
	if r.Header.Get("Content-Type") != CONTENT_TYPE_JSON {
		return errors.New("Content-Type must be application/json")
	}

	// リクエストのボディをデコード、エラーがあればエラーレスポンスをJsonで返す
	if err := json.NewDecoder(r.Body).Decode(in); err != nil {
		return err
	}

	slog.Debug("jsonInput", "header", r.Header, "body", in)
	return nil
}

// Jsonへの出力を行うヘルパ関数
func jsonOutput(w http.ResponseWriter, out interface{}) error {
	// レスポンスヘッダを設定
	w.Header().Set("Content-Type", CONTENT_TYPE_JSON)
	w.WriteHeader(http.StatusOK)

	// レスポンスボディをエンコード、エラーがあればエラーレスポンスを返す
	if err := json.NewEncoder(w).Encode(out); err != nil {
		return err
	}

	slog.Debug("jsonOutput", "header", w.Header(), "body", out)
	return nil
}

// Jsonのエラーレスポンスを返すヘルパ関数。APIのエラーはこの関数を使ってハンドリングする
func jsonErrorOutput(w http.ResponseWriter, statusCode int, err error) {
	// エラーレスポンスをJsonで返す
	w.Header().Set("Content-Type", CONTENT_TYPE_JSON)
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(jsonError{
		StatusCode: statusCode,
		Error:      err.Error(),
	})

	slog.Debug("error", "statusCode", statusCode, "desc", err)
}

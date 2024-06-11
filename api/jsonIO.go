package api

import (
	"encoding/json"
	"net/http"
)

type jsonError struct {
	StatusCode int    `json:"statusCode"`
	Error      string `json:"error"`
}

// Jsonの入力を受け取るヘルパ関数
func jsonInput(w http.ResponseWriter, r *http.Request, in interface{}) {
	// リクエストのヘッダをチェック
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Content-Type must be application/json", http.StatusBadRequest)
		return
	}

	// リクエストのボディをデコード、エラーがあればエラーレスポンスをJsonで返す
	if err := json.NewDecoder(r.Body).Decode(in); err != nil {
		jsonErrorOutput(w, http.StatusBadRequest, err)
		return
	}
}

// Jsonへの出力を行うヘルパ関数
func jsonOutput(w http.ResponseWriter, out interface{}) {
	// レスポンスヘッダを設定
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	// レスポンスボディをエンコード、エラーがあればエラーレスポンスを返す
	if err := json.NewEncoder(w).Encode(out); err != nil {
		jsonErrorOutput(w, http.StatusInternalServerError, err)
		return
	}

	// レスポンスを返す
	w.WriteHeader(http.StatusOK)
}

// Jsonのエラーレスポンスを返すヘルパ関数
func jsonErrorOutput(w http.ResponseWriter, statusCode int, err error) {
	// エラーレスポンスをJsonで返す
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(jsonError{
		StatusCode: statusCode,
		Error:      err.Error(),
	})
}

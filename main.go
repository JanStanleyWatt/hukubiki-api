package main

import (
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/JanStanleyWatt/hukubiki-api/api"
)

const TIMEOUT = 500 * time.Millisecond

func main() {
	// slog初期化
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	slog.SetDefault(logger)

	// サーバー初期化
	logger.Info("Starting server...")
	server := &http.Server{
		Addr:              ":8000",
		ReadTimeout:       TIMEOUT,
		ReadHeaderTimeout: TIMEOUT,
		Handler:           http.TimeoutHandler(http.DefaultServeMux, TIMEOUT, "Timeout!"),
	}
	slog.Debug("server initialized",
		"addr", server.Addr,
		"readTimeoutMs", TIMEOUT,
		"readHeaderTimeoutMs", TIMEOUT,
		"handlerTimeoutMs", TIMEOUT)

	// ルーティング定義
	slog.Info("Registering handlers...")
	http.HandleFunc("/int64/n", api.Int64n)
	http.HandleFunc("/", api.Usage)

	// サーバー起動
	slog.Info("Server started.")
	server.ListenAndServe()
}

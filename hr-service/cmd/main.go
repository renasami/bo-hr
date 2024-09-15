package main

import (
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/renasami/bo-hr/hr-service/internal/handler"
)

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 特定のオリジンのみを許可
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// OPTIONSメソッドへの対応（プリフライトリクエスト）
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

type loggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func newLoggingResponseWriter(w http.ResponseWriter) *loggingResponseWriter {
	// デフォルトステータスコードは200に設定
	return &loggingResponseWriter{w, http.StatusOK}
}

func (lrw *loggingResponseWriter) WriteHeader(code int) {
	lrw.statusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}

func loggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		method := r.Method
		url := r.RequestURI
		ctx := r.Context()

		slog.Info("Request Received",
			slog.String("method", method),
			slog.String("url", url),
			slog.Any("context", ctx),
		)

		lrw := newLoggingResponseWriter(w)

		next.ServeHTTP(w, r)

		duration := time.Since(start)
		statusCode := lrw.statusCode

		switch {
		case statusCode >= 200 && statusCode < 300:
			slog.Info("Request succeeded",
				slog.String("method", method),
				slog.String("url", url),
				slog.Int("status", statusCode),
				slog.Duration("duration", duration),
			)
		case statusCode >= 400 && statusCode < 500:
			slog.Warn("Client error occurred",
				slog.String("method", method),
				slog.String("url", url),
				slog.Int("status", statusCode),
				slog.Duration("duration", duration),
			)
		case statusCode >= 500:
			slog.Error("Server error occurred",
				slog.String("method", method),
				slog.String("url", url),
				slog.Int("status", statusCode),
				slog.Duration("duration", duration),
			)
		default:
			slog.Info("Unknown status",
				slog.String("method", method),
				slog.String("url", url),
				slog.Int("status", statusCode),
				slog.Duration("duration", duration),
			)
		}
	})
}

func main() {
	r := mux.NewRouter()

	r.Use(corsMiddleware)

	r.HandleFunc("/employees", handler.GetEmployees).Methods("GET")
	r.HandleFunc("/user_title", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		fmt.Fprintf(w, "<h1>Hello</h1>")
	}).Methods("GET", "OPTIONS")
	r.HandleFunc("/user_title", handler.CreateUserTitle).Methods("POST")

	log.Println("Starting HR service on port 8081...")
	log.Fatal(http.ListenAndServe("localhost:8081", r))
}

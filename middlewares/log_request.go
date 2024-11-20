package middlewares

import (
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/charmbracelet/log"
)

type responseWriter struct {
	http.ResponseWriter
	size       int
	statusCode int
}

func (w *responseWriter) WriteHeader(statusCode int) {
	w.statusCode = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}

func (w *responseWriter) Write(b []byte) (int, error) {
	size, err := w.ResponseWriter.Write(b)
	w.size += size
	return size, err
}

func LogRequest(next http.Handler) http.Handler {
	logger := log.NewWithOptions(os.Stdout, log.Options{
		ReportCaller:    false,
		ReportTimestamp: true,
		TimeFormat:      time.Kitchen,
		Prefix:          "fresh",
	})

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		rw := &responseWriter{ResponseWriter: w, statusCode: http.StatusOK}

		logger.Info("incoming request",
			"method", r.Method,
			"path", r.URL.Path,
			"ip", r.RemoteAddr,
			"ua", truncateUA(r.UserAgent()),
			"proto", r.Proto,
		)

		next.ServeHTTP(rw, r)

		logger.Info("request completed",
			"method", r.Method,
			"path", r.URL.Path,
			"status", rw.statusCode,
			"size", formatBytes(rw.size),
			"time", time.Since(start).String(),
			"ip", r.RemoteAddr,
		)
	})
}

func truncateUA(ua string) string {
	const maxLen = 50
	if len(ua) <= maxLen {
		return ua
	}
	return ua[:maxLen-3] + "..."
}

func formatBytes(bytes int) string {
	const unit = 1024
	if bytes < unit {
		return strconv.Itoa(bytes) + " B"
	}

	div, exp := int64(unit), 0
	for n := bytes / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}

	sizes := []string{"B", "KB", "MB", "GB", "TB", "PB", "EB"}
	if exp >= len(sizes) {
		return "TOO LARGE"
	}

	return strconv.FormatFloat(float64(bytes)/float64(div), 'f', 1, 64) + " " + sizes[exp]
}

package main

import (
	"fmt"
	"net/http"
)

// Middleware thêm các HTTP security headers vào phản hồi
func secureHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Ngăn chặn XSS, hạn chế tải tài nguyên bên ngoài
		w.Header().Set("Content-Security-Policy",
			"default-src 'self'; style-src 'self' https://fonts.googleapis.com; font-src https://fonts.gstatic.com")

		// Chỉ gửi referrer khi cùng origin hoặc khi điều hướng cross-origin
		w.Header().Set("Referrer-Policy", "origin-when-cross-origin")

		// Ngăn chặn MIME sniffing
		w.Header().Set("X-Content-Type-Options", "nosniff")

		// Ngăn chặn Clickjacking
		w.Header().Set("X-Frame-Options", "DENY")

		// Vô hiệu hóa bộ lọc XSS của trình duyệt (vì CSP đã xử lý)
		w.Header().Set("X-XSS-Protection", "0")

		// Bắt buộc sử dụng HTTPS (chỉ nên bật nếu dùng HTTPS)
		w.Header().Set("Strict-Transport-Security", "max-age=63072000; includeSubDomains; preload")

		next.ServeHTTP(w, r)
	})
}

func (app *application) logRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		app.infoLog.Printf("%s - %s %s %s", r.RemoteAddr, r.Proto, r.Method,
			r.URL.RequestURI())
		next.ServeHTTP(w, r)
	})
}

func (app *application) recoverPanic(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Create a deferred function (which will always be run in the event
		// of a panic as Go unwinds the stack).
		defer func() {
			// Use the builtin recover function to check if there has been a
			// panic or not. If there has...

			if err := recover(); err != nil {
				// Set a "Connection: close" header on the response.
				w.Header().Set("Connection", "close")
				// Call the app.serverError helper method to return a 500
				// Internal Server response.
				app.serverError(w, fmt.Errorf("%s", err))
			}
		}()
		next.ServeHTTP(w, r)
	})
}

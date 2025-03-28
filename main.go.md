// practice
package main

import (
	"fmt" // New import
	"log"
	"net/http"
	"strconv" // New import
)

// Handler cho trang chủ
// func home(w http.ResponseWriter, r *http.Request) {
// 	log.Println("Request:", r.Method, r.URL.Path)
// 	w.Write([]byte("Hello from Snippetbox"))
// }

// func home(w http.ResponseWriter, r *http.Request) {
// 	// Check if the current request URL path exactly matches "/". If it doesn't use
// 	// the http.NotFound() function to send a 404 response to the client.
// 	// Importantly, we then return from the handler. If we don't return the handler
// 	// would keep executing and also write the "Hello from SnippetBox" message.
// 	if r.URL.Path != "/" {
// 	http.NotFound(w, r)
// 	return
// 	}
// 	w.Write([]byte("Hello from Snippetbox"))
// 	}

func home(w http.ResponseWriter, r *http.Request) {
	// Nếu không phải "/", trả về 404
	if r.URL.Path != "/" {
		log.Printf("404 Not Found: %s\n", r.URL.Path)
		http.Error(w, "404 - Page Not Found", http.StatusNotFound)
		return
	}

	// Thiết lập Content-Type để tránh lỗi hiển thị trên trình duyệt
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("Hello from Snippetbox"))
}

// Xem snippet cụ thể
// func snippetView(w http.ResponseWriter, r *http.Request) {
// 	log.Println("Request:", r.Method, r.URL.Path)
// 	w.Write([]byte("Display a specific snippet..."))
// }

func snippetView(w http.ResponseWriter, r *http.Request) {
	// Extract the value of the id parameter from the query string and try to
	// convert it to an integer using the strconv.Atoi() function. If it can't
	// be converted to an integer, or the value is less than 1, we return a 404 page
	// not found response.
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	// Use the fmt.Fprintf() function to interpolate the id value with our response
	// and write it to the http.ResponseWriter.
	fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
}

// Tạo snippet mới, chỉ cho phép phương thức POST
// func snippetCreate(w http.ResponseWriter, r *http.Request) {
// 	if r.Method != http.MethodPost {
// 		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
// 		return
// 	}
// 	log.Println("Request:", r.Method, r.URL.Path)
// 	w.Write([]byte("Create a new snippet..."))
// }

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	// Use r.Method to check whether the request is using POST or not.
	if r.Method != "POST" {
		// If it's not, use the w.WriteHeader() method to send a 405 status
		// code and the w.Write() method to write a "Method Not Allowed"
		// response body. We then return from the function so that the
		// subsequent code is not executed.
		w.Header().Set("Allow", "POST")

		// Set a new cache-control header. If an existing "Cache-Control" header exists
		// it will be overwritten.
		w.Header().Set("Cache-Control", "public, max-age=31536000")
		// Thiết lập header bảo vệ chống XSS
		w.Header().Set("X-XSS-Protection", "1; mode=block")

		// In contrast, the Add() method appends a new "Cache-Control" header and can
		// be called multiple times.
		w.Header().Add("Cache-Control", "public")
		w.Header().Add("Cache-Control", "max-age=31536000")

		// Delete all values for the "Cache-Control" header.
		w.Header().Del("Cache-Control")

		// Retrieve the first value for the "Cache-Control" header.
		w.Header().Get("Cache-Control")
		// Retrieve a slice of all values for the "Cache-Control" header.
		w.Header().Values("Cache-Control")

		// w.WriteHeader(405)
		// w.Write([]byte("Method Not Allowed"))
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Write([]byte("Create a new snippet..."))

}

// Khởi tạo router
// func newRouter() *http.ServeMux {
// 	mux := http.NewServeMux()
// 	mux.HandleFunc("/", home)
// 	mux.HandleFunc("/snippet/view", snippetView)
// 	mux.HandleFunc("/snippet/create", snippetCreate)
// 	return mux
// }

// func main() {
// 	addr := ":4000"
// 	log.Println("Starting server on", addr)

// 	server := &http.Server{
// 		Addr:    addr,
// 		Handler: newRouter(),
// 	}

// 	err := server.ListenAndServe()
// 	if err != nil {
// 		log.Fatal("Server error:", err)
// 	}
// }

//khong nen dung
// func main() {
// 	http.HandleFunc("/", home)
// 	http.HandleFunc("/snippet/view", snippetView)
// 	http.HandleFunc("/snippet/create", snippetCreate)
// 	log.Println("Starting server on :4000")
// 	err := http.ListenAndServe(":4000", nil)
// 	log.Fatal(err)
// }

// Middleware: Logging request
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request: %s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

// Tạo router
func newRouter() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)
	return mux
}

func main() {
	addr := ":4000"
	log.Println("Starting server on", addr)

	mux := newRouter()
	handler := loggingMiddleware(mux) // Thêm middleware

	server := &http.Server{
		Addr:    addr,
		Handler: handler,
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("Server error:", err)
	}
}

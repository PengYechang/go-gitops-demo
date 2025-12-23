package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	// 获取版本号环境变量
	version := os.Getenv("VERSION")
	if version == "" {
		version = "v1.0.0"
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		message := fmt.Sprintf("Hello, Argo CD! This is UPDATED VERSION: %s !!!", version)
		fmt.Fprintf(w, message)
		fmt.Println("Request received:", message)
	})

	fmt.Println("Server starting on :8080...")
	http.ListenAndServe(":8080", nil)
}

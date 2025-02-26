package main

import (
	"fmt"
	"net/http"
	"io"
)

func main() {
	fmt.Println("🚀 Client is requesting the server...")

	resp, err := http.Get("http://localhost:8080")
	if err != nil {
		fmt.Println("❌ Error:", err)
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	fmt.Println("📡 Server Response:", string(body))
}

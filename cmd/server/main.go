package main

import (
	"github.com/osdeving/fc-go-api-example/configs"
)

func main() {
	config, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	println(config.DBDriver)

	// fmt.Println("🚀 Server is starting on port 8080...")
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintln(w, "Hello, World!")
	// })
	// http.ListenAndServe(":8080", nil)
}

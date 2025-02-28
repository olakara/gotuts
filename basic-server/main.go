package main

import (
	"fmt"
	"net/http"
)

func main() {

	const port = 8080
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})

	fmt.Println("Server is running on port: ", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
}

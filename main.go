package main

import (
	"fmt"
	"net/http"
)

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, Vamshi")
}

func customResponseHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, Tara")
}

func main() {
	http.HandleFunc("/user1", helloWorldHandler)
	http.HandleFunc("/user2", customResponseHandler)

	port := 8080
	fmt.Printf("Server is running on port %d...\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		fmt.Println("Error:", err)
	}
}

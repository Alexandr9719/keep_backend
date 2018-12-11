package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", authHandler)
	log.Println("Listening api port 3000...")
	http.ListenAndServe(":3000", nil)

}

func authHandler(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(response, "Hello. Status %d", http.StatusOK)
}

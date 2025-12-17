package main

import (
	"fmt"
	"log"
	"net/http"
	"sw2_hw1/helloworld"
)

func main() {
	port := 8080
	// for now its http://localhost:8080/helloworld
	http.HandleFunc("/helloworld", helloworld.HelloWorldHandler)
	http.HandleFunc("/goodbyeworld", helloworld.GoodbyeWorldHandler)
	http.HandleFunc("/ggoodbyeworld", helloworld.GracefulGoodbyeWorldHandler)

	log.Printf("Server starting on port: %v\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}

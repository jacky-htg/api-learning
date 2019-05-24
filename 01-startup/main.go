package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	// handler
	handler := http.HandlerFunc(helloworld)

	// start server listening
	if err := http.ListenAndServe("localhost:9000", handler); err != nil {
		log.Fatalf("error: listening and serving: %s", err)
	}
}

// helloworld: basic http handler with response hello world string
func helloworld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World!")
}

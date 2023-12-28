package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/combinedData", combineData)
	fmt.Println("Server listening on port : 8080")
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

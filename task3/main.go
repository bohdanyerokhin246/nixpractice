package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	MakeRequest()
}

func MakeRequest() {
	response, errorResponse := http.Get("https://jsonplaceholder.typicode.com/posts/")
	if errorResponse != nil {
		log.Fatalln(errorResponse)
	}

	body, errorRead := io.ReadAll(response.Body)
	if errorRead != nil {
		log.Fatalln(errorRead)
	}

	fmt.Println(string(body))
}

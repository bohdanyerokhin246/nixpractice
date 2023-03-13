package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
)

func main() {

	for i := 0; i < 101; i++ {
		go MakeRequest(i)
	}
	fmt.Scanln()
}

func MakeRequest(i int) {
	response, errorResponse := http.Get("https://jsonplaceholder.typicode.com/posts/" + strconv.Itoa(i))
	if errorResponse != nil {
		log.Fatalln(errorResponse)
	}

	body, errorRead := io.ReadAll(response.Body)
	if errorRead != nil {
		log.Fatalln(errorRead)
	}
	fmt.Println(string(body))

}

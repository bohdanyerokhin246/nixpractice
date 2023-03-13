package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
)

func main() {

	writeFiles()

}

func writeFiles() {

	directoryName := "storage/posts"

	errorMakeDirection := os.MkdirAll(directoryName, 077)
	if errorMakeDirection != nil {
		log.Fatalln(errorMakeDirection)
	}

	errorChangeDirection := os.Chdir(directoryName)
	if errorChangeDirection != nil {
		log.Fatalln(errorChangeDirection)
	}
	for i := 0; i < 101; i++ {
		go makeRequest(i)
	}
}

func makeRequest(i int) {
	response, errorResponse := http.Get("https://jsonplaceholder.typicode.com/posts/" + strconv.Itoa(i))
	if errorResponse != nil {
		log.Fatalln(errorResponse)
	}

	file, errorCreateFile := os.Create(strconv.Itoa(i) + ".txt")
	if errorCreateFile != nil {
		log.Fatalln(errorCreateFile)
	}

	defer file.Close()

	body, errorRead := io.ReadAll(response.Body)
	if errorRead != nil {
		log.Fatalln(errorRead)
	}
	file.WriteString(string(body))
}

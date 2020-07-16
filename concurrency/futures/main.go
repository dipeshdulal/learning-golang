package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type data struct {
	Body  []byte
	Error error
}

func getRequest(url string) chan data {
	c := make(chan data)
	go func() {
		resp, err := http.Get(url)
		if err != nil {
			c <- data{Error: err}
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		c <- data{Body: body, Error: err}
	}()
	return c
}

func main() {
	todos := getRequest("https://jsonplaceholder.typicode.com/todos/1")
	posts := getRequest("https://jsonplaceholder.typicode.com/posts/1")

	myData := data{Body: []byte("MyData"), Error: nil}
	printBody(myData)

	// some lengthy task
	time.Sleep(1 * time.Second)
	fmt.Println("After sleeping")

	printBody(<-todos)
	printBody(<-posts)
}

func printBody(b data) {
	fmt.Printf("Body: %v \n", string(b.Body))
	fmt.Printf("Error: %v \n", b.Error)
}

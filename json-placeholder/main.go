package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	apiUrl := "https://jsonplaceholder.typicode.com/posts/1"

	response, err := http.Get(apiUrl)
	if err != nil {
		fmt.Println("Error makeing GET request: ", err)
		return
	}
	defer response.Body.Close()

	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response body: ", err)
		return
	}

	fmt.Println("API Response:")
	fmt.Println(string(responseBody))
}

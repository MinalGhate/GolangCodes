package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	//url of the api endpoint
	apiUrl := "https://api.example.com/data"

	//send get request
	response, err := http.Get(apiUrl)

	if err != nil {
		fmt.Printf("Error making Get request: %s\n", err)
		return
	}

	defer response.Body.Close() //will execute after completion of function

	//read the response body
	body, err := io.ReadAll(response.Body)

	if err != nil {
		fmt.Printf("Error reading response body: %s\n", err)
		return
	}

	fmt.Println(string(body))
}

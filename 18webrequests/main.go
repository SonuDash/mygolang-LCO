package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://www.lco.dev/"

func main() {
	fmt.Println("This is a class on how to handle web requests")

	response, err := http.Get(url)

	checkNilErr(err)
	fmt.Printf("\nthe response type is: %T\n", response)
	fmt.Println("\nthe response is:", response)

	defer response.Body.Close()

	dataBytes, err := ioutil.ReadAll(response.Body)
	checkNilErr(err)
	content := string(dataBytes)
	fmt.Println(content)

}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}

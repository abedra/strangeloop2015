package main

import (
	"fmt"
	"net/http"
	"os"
	"io/ioutil"
)

func main() {
	response, err := http.Get("http://localhost:8080")

	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	} else {
		defer response.Body.Close()
		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Printf("%s", err)
			os.Exit(1)
		}
		fmt.Printf("%s\n", string(contents))
	}
}

package main

import (
        "fmt"
        "net/http"
        "os"
        "net/url"
        "bufio"
)


func main() {
        file, err := os.Open("dictionary.txt")
        if err != nil {
                fmt.Println(err)
        }
        defer file.Close()

        var lines []string
        scanner := bufio.NewScanner(file)
        for scanner.Scan() {
                lines = append(lines, scanner.Text())
        }


        data := url.Values{}
        data.Set("inputEmail", "admin@example.com")
        data.Set("inputPassword", "password")

        response, err := http.PostForm("http://localhost:8080/", data)
        if err != nil {
                fmt.Println(err)
        }
        defer response.Body.Close()

	fmt.Println(response.Status)
}

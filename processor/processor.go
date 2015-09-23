package main

import (
        "fmt"
        "os"
        "flag"
        "bufio"
        "strings"
        "github.com/fzzy/radix/redis"
)

type logEntry struct {
        Address string
        Method  string
        Uri     string
        ResponseCode string
}

func main() {
        logFilePtr := flag.String("file", "", "Path to log file")
        thresholdPtr := flag.Int("threshold", 10, "Threshold for failed logins")
        flag.Parse()

        if *logFilePtr != "" {
                file, err := os.Open(*logFilePtr)
                if err != nil {
                        fmt.Println("Couldn't open logfile:", err)
                        os.Exit(1)
                }
                defer file.Close()

                var lines []string
                scanner := bufio.NewScanner(file)
                for scanner.Scan() {
                        lines = append(lines, scanner.Text())
                }

                var entries map[string]int
                entries = make(map[string]int)

                for _, entry := range lines {
                        parts := strings.Split(entry, " ")
                        l := logEntry{
                                Address: parts[0],
                                Method: strings.Replace(parts[5], "\"", "", 1),
                                Uri: parts[6],
                                ResponseCode: parts[8],
                        }
                        if l.Method == "POST" && l.ResponseCode == "200" {
                                entries[l.Address] += 1
                        }
                }

                if len(entries) > 0 {
                        connection, err := redis.Dial("tcp", "localhost:6379")
                        if err != nil {
                                fmt.Println("Couldn't connect to Redis:", err)
                                os.Exit(1)
                        }

                        connection.Cmd("MULTI")
                        for k, v := range entries {
                                if v >= *thresholdPtr {
                                        fmt.Printf("Blacklisting %s. Threshold: %d, Actual: %d\n", k, *thresholdPtr, v)
                                        actorString := fmt.Sprintf("%s:repsheet:ip:blacklisted", k)
                                        connection.Cmd("SET", actorString, "Failed login processor")
                                }
                        }
                        connection.Cmd("EXEC")
                }
        }
}

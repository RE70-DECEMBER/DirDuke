package main

import (
	"bufio"
	"flag"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"
)

var (
	url      = flag.String("u", "http://localhost/", "Target URL")
	wordlist = flag.String("w", "directory-list-2.3-medium.txt", "Wordlist")
)

func main() {
	flag.Parse()

	fmt.Printf(` 
 _____  _      _____        _        
 |  __ \(_)    |  __ \      | |       
 | |  | |_ _ __| |  | |_   _| | _____ 
 | |  | | | '__| |  | | | | | |/ / _ \
 | |__| | | |  | |__| | |_| |   <  __/
 |_____/|_|_|  |_____/ \__,_|_|\_\___|

 Author: RE70-DECEMBER
 Github: https://github.com/RE70-DECEMBER                     
 target: %s
 wordlist: %s
 `, *url, *wordlist)

	file, _ := os.Open(*wordlist)

	defer file.Close()

	scanner := bufio.NewScanner(file)
	client := &http.Client{Timeout: 5 * time.Second}
	fmt.Println()

	for scanner.Scan() {
		line := scanner.Text()
		word := strings.TrimSpace(line)

		if word == "" || strings.HasPrefix(word, "#") {
			continue
		}

		target := strings.TrimRight(*url, "/") + "/" + word

		resp, err := client.Get(target)

		if err != nil {
			continue
		}

		resp.Body.Close()

		if resp.StatusCode == 200 {
			fmt.Printf("\033[32m[*]\033[0m Found: %v\n", target)
		} else if resp.StatusCode != 404 {
			fmt.Printf("\033[33m[?]\033[0m %s \t Status Code: %d\n", target, resp.StatusCode)
		}

	}

	fmt.Println()
	fmt.Println("[*] Done Scanning!")
}

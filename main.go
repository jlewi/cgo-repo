package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
)

func main() {
	url := flag.String("url", "https://www.google.com", "the url to fetch")
	flag.Parse()
	c := &http.Client{}
	resp, err := c.Get(*url)
	if err != nil {
		fmt.Printf("Get failed; returned error %v", err)
		return
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Request succeeded but error reading response; %v", err)
	}
	fmt.Printf("Success\nresponse\n%v", string(body))
}

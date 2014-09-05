package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	client := &http.Client{}

	req, err := http.NewRequest("GET", "http://twitpic.com/photos/harukasan.json", nil)
	req.Header.Add("User-Agent", "Mozilla/5.0")

	resp, err := client.Do(req)

	if err != nil {
		fmt.Printf("error: %s\n", err.Error())
		os.Exit(1)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Printf("error: %s\n", err.Error())
		os.Exit(1)
	}

	fmt.Printf("%s\n", body)
}

package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	resp, err := client.Get("http://twitpic.com/photos/harukasan.json")

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

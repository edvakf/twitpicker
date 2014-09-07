package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"./twitpic"
)

func main() {
	body, err := getHTTP("http://twitpic.com/photos/harukasan.json")

	if err != nil {
		fmt.Printf("error: %s\n", err.Error())
		os.Exit(1)
	}

	//fmt.Printf("%s\n", body)

	photos := twitpic.DecodePhotos(body)

	//fmt.Printf("%+v\n", photos)

	for i := 0; i < len(photos.Images); i++ {
		//fmt.Printf("%s\n", photos.Images[i].ShortID)
		fmt.Printf("%s\n", photos.Images[i].ToURL())
	}
}

func getHTTP(url string) (string, error) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return "", err
	}

	req.Header.Add("User-Agent", "Mozilla/5.0")
	resp, err := client.Do(req)

	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return "", err
	}

	return string(body[:]), nil
}

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"sync"

	"./twitpic"
)

var numDownloads = 2

func main() {
	body, err := getHTTP("http://twitpic.com/photos/harukasan.json")

	if err != nil {
		fmt.Printf("error: %s\n", err.Error())
		os.Exit(1)
	}

	//fmt.Printf("%s\n", body)

	photos := twitpic.DecodePhotos(body)

	//fmt.Printf("%+v\n", photos)

	downloadImages(photos)
}

func downloadImages(photos twitpic.Photos) {
	chs := make(chan twitpic.Image, len(photos.Images))

	var wg sync.WaitGroup
	defer wg.Wait()

	for n := 0; n < numDownloads; n++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for img := range chs {
				err := img.Download()
				if err != nil {
					fmt.Printf("error: %s\n", err.Error())
					os.Exit(1)
				}
			}
		}()
	}

	for _, img := range photos.Images {
		chs <- img
	}
	close(chs)
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

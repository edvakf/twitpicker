package twitpic

/*
{
id: "130707",
twitter_id: "6517782",
username: "harukasan",
name: "はるかさん",
...
images: [
	{
		id: "844034960",
		short_id: "dyiljk",
		user_id: "130707",
		source: "api",
		message: "はてブ見てるとときたまこうなってる",
		views: "0",
		width: "965",
		height: "585",
		size: "302418",
		type: "png",
		...
	},
*/

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"os"
)

type Image struct {
	ShortID string `json:"short_id"`
	Type    string `json:"type"`
}

type Photos struct {
	Images []Image
}

func DecodePhotos(phJson []byte) Photos {
	var p Photos
	json.Unmarshal(phJson, &p)
	return p
}

func (img Image) ToURL() string {
	return "http://twitpic.com/show/large/" + img.ShortID
}

func (img Image) Download() error {
	log.Println(">> starting download ", img.ShortID)

	var resp *http.Response
	var err error

	numRetry := 3
	for i := 0; i < numRetry; i++ {
		resp, err = http.Get(img.ToURL())
		if err != nil {
			return err
		}
		if resp.StatusCode != 200 {
			log.Println("bad http status ", resp.StatusCode)
			resp.Body.Close()
			if i == numRetry-1 {
				return errors.New("maxium number of retry reached")
			}
			continue
		}
		defer resp.Body.Close()
		break
	}

	name := img.ShortID + "." + img.Type
	f, err := os.Create(name)

	if err != nil {
		return err
	}

	written, err := io.Copy(f, resp.Body)

	if err != nil {
		return err
	}

	log.Printf("<< saved file %s (%d bytes)\n", name, written)
	return nil
}

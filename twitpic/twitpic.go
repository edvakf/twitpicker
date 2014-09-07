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
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

type Image struct {
	ShortID string `json:"short_id"`
	Type    string `json:"type"`
}

type Photos struct {
	Images []Image
}

func DecodePhotos(phJson string) Photos {
	dec := json.NewDecoder(strings.NewReader(phJson))
	var p Photos
	dec.Decode(&p)
	return p
}

func (img Image) ToURL() string {
	return "http://twitpic.com/show/large/" + img.ShortID
}

func (img Image) Download() error {
	log.Println(">> starting download")

	resp, err := http.Get(img.ToURL())

	if err != nil {
		return err
	}
	defer resp.Body.Close()

	name := img.ShortID + "." + img.Type
	f, err := os.Create(name)

	if err != nil {
		return err
	}

	written, err := io.Copy(f, resp.Body)

	if err != nil {
		return err
	}

	log.Printf("%d bytes written to %s\n", written, name)

	if resp.ContentLength > 0 && resp.ContentLength != written {
		log.Printf(
			"content-length (%d) does not match the file size (%d)\n",
			resp.ContentLength, written)
	}

	log.Printf("<< saved file %s\n", name)
	return nil
}

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
	"io/ioutil"
	"net/http"
	"strings"
)

type image struct {
	ShortID string `json:"short_id"`
	Type    string `json:"type"`
}

type photos struct {
	Images []image
}

func DecodePhotos(phJson string) photos {
	dec := json.NewDecoder(strings.NewReader(phJson))
	var p photos
	dec.Decode(&p)
	return p
}

func (img image) ToURL() string {
	return "http://twitpic.com/show/large/" + img.ShortID
}

func (img image) Download() error {
	resp, err := http.Get(img.ToURL())

	if err != nil {
		return err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	name := img.ShortID + "." + img.Type
	err = ioutil.WriteFile(name, data, 0666)
	if err != nil {
		return err
	}

	return nil
}

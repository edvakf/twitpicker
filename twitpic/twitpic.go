package twitpic

import (
	"encoding/json"
	"strings"
)

type image struct {
	ShortID string `json:"short_id"`
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

package spotify

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/ayberktandogan/melody/config"
	"golang.org/x/oauth2"
)

type explicitContent struct {
	FilterEnabled bool `json:"filter_enabled"`
	FilterLocked  bool `json:"filter_locked"`
}

type externalUrls struct {
	Spotify string `json:"spotify"`
}

type followers struct {
	Href  string `json:"href"`
	Total int    `json:"total"`
}

type image struct {
	Url    string `json:"url"`
	Height int    `json:"height"`
	Width  int    `json:"width"`
}

type SpotifyUser struct {
	Country         string          `json:"country"`
	DisplayName     string          `json:"display_name"`
	Email           string          `json:"email"`
	ExplicitContent explicitContent `json:"explicit_content"`
	ExternalUrls    externalUrls    `json:"external_urls"`
	Followers       followers       `json:"followers"`
	Href            string          `json:"href"`
	Id              string          `json:"id"`
	Images          []image         `json:"images"`
	Product         string          `json:"product"`
	Type            string          `json:"type"`
	Uri             string          `json:"uri"`
}

func (s *SpotifyClient) GetMe(uc oauth2.Token) (*SpotifyUser, error) {
	req, err := http.NewRequest(http.MethodGet, config.Config.BaseAPIUri+"/me", nil)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	req.Header.Add("Authorization", "Bearer "+uc.AccessToken)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	var d SpotifyUser
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&d); err != nil {
		log.Fatal(err)
		return nil, err
	}

	return &d, nil
}

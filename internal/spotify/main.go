package spotify

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/ayberktandogan/melody/config"
)

func (s *SpotifyClient[T]) Request(path string) (*T, error) {
	req, err := http.NewRequest(http.MethodGet, config.Config.BaseAPIUri+path, nil)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	req.Header.Add("Authorization", "Bearer "+s.Auth.AccessToken)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	var d T
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&d); err != nil {
		log.Fatal(err)
		return nil, err
	}

	return &d, nil
}

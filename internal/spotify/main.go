package spotify

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/ayberktandogan/melody/config"
	"golang.org/x/oauth2"
)

type SpotifyClient struct {
	Auth        oauth2.Token
	Body        []byte
	QueryParams map[string]string
}

func Request[T any](s *SpotifyClient, method string, path string) (*T, *http.Response, error) {
	req, err := http.NewRequest(method, config.Config.BaseAPIUri+path, bytes.NewBuffer(s.Body))
	if err != nil {
		log.Fatal(err)
		return nil, nil, err
	}

	req.Header.Add("Authorization", "Bearer "+s.Auth.AccessToken)
	req.Header.Add("Content-Type", "application/json; charset=UTF-8")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
		return nil, res, err
	}

	var d T
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&d); err != nil {
		if err == io.EOF {
			return nil, res, nil
		}
		log.Fatal(err)
		return nil, res, err
	}

	return &d, res, nil
}

func ParseBody[K any](s *SpotifyClient, b *K) {
	enc, err := json.Marshal(b)
	if err != nil {
		log.Fatal(err)
	}

	s.Body = enc
}

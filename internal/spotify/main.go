package spotify

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/ayberktandogan/melody/config"
	"github.com/ayberktandogan/melody/internal/utils"
	"golang.org/x/oauth2"
)

type SpotifyClient struct {
	Auth        oauth2.Token
	Body        []byte
	QueryParams map[string]string
}

func Request[T any](s *SpotifyClient, method string, path string) (*T, *http.Response, error) {
	url := config.Config.BaseAPIUri + path
	if len(s.QueryParams) != 0 {
		url = utils.AppendQueries(url, s.QueryParams)
	}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(s.Body))
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

	if rerr := getError(res); rerr != nil {
		return nil, res, rerr
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

func getError(res *http.Response) error {
	if res.StatusCode == 401 {
		return fmt.Errorf("unauthorized, %d", res.StatusCode)
	}

	if res.StatusCode == 403 {
		return fmt.Errorf("forbidden, %d", res.StatusCode)
	}

	if res.StatusCode == 404 {
		return fmt.Errorf("not found, %d", res.StatusCode)
	}

	if res.StatusCode > 299 || res.StatusCode < 200 {
		fmt.Printf("We tried to hit %s, but had a dead end\n", res.Request.URL)
		return fmt.Errorf("we had a problem, %d", res.StatusCode)
	}

	return nil
}

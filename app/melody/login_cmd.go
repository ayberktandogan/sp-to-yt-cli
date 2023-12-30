package app

import (
	"log"

	"github.com/ayberktandogan/melody/internal/spotify"
	"golang.org/x/oauth2"
)

type loginCmd struct {
}

func (i *loginCmd) Run() error {
	sc := spotify.SpotifyClient[oauth2.Token]{}
	res, err := sc.Login()
	if err != nil {
		log.Fatal(err)
		return err
	}

	UserConfig.Spotify = *res
	SaveUserConfig()

	return nil
}

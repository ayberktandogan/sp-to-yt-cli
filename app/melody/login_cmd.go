package app

import (
	"log"
)

type loginCmd struct {
}

func (i *loginCmd) Run(clients *Clients, userConfig *userConfig) error {
	res, err := clients.Spotify.Login()
	if err != nil {
		log.Fatal(err)
		return err
	}

	clients.Spotify.Auth = *res
	userConfig.Data.Spotify = *res
	userConfig.SaveUserConfig()

	return nil
}

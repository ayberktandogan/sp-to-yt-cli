package app

import (
	"encoding/json"
	"log"
	"os"
	"time"

	"github.com/alecthomas/kong"
	"github.com/ayberktandogan/melody/internal/spotify"
	"golang.org/x/oauth2"
)

type userConfigData struct {
	Spotify oauth2.Token `json:"spotify"`
}

type userConfig struct {
	Data userConfigData `json:"data"`
}

const userConfigFolder = "~/.config/melody"
const userConfigPath = userConfigFolder + "/config"

var defaultUserConfig = &userConfigData{
	Spotify: oauth2.Token{
		RefreshToken: "",
		AccessToken:  "",
		Expiry:       time.Now(),
		TokenType:    "",
	},
}

var UserConfig = &userConfig{}

func (u *userConfig) LoadUserConfig() error {
	createDirIfNotExists()
	createFileIfNotExists(userConfigPath)

	data := readFromFile(userConfigPath)
	u.Data = *data

	refreshTokenIfNecessary(u)

	return nil
}

func (u *userConfig) SaveUserConfig() error {
	err := writeToFile(userConfigPath, u.Data)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	return nil
}

func (u *userConfig) DeleteUserConfig() error {
	if err := deleteFile(userConfigPath); err != nil {
		return err
	}
	return nil
}

func createDirIfNotExists() {
	err := os.MkdirAll(kong.ExpandPath(userConfigFolder), os.ModePerm)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
}

func createFileIfNotExists(filePath string) {
	f, err := os.OpenFile(kong.ExpandPath(filePath), os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	if s, _ := f.Stat(); s.Size() == 0 {
		f.Close()
		writeToFile(filePath, defaultUserConfig)
	}
}

func readFromFile(filePath string) *userConfigData {
	f, err := os.ReadFile(kong.ExpandPath(filePath))
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	var d userConfigData

	ferr := json.Unmarshal(f, &d)
	if ferr != nil {
		panic(ferr)
	}

	return &d
}

func writeToFile(filePath string, d any) error {
	f, err := os.OpenFile(kong.ExpandPath(filePath), os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	c, err := json.Marshal(d)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	_, ferr := f.Write(c)
	if ferr != nil {
		log.Fatal(ferr)
		panic(ferr)
	}

	f.Close()
	return nil
}

func deleteFile(filePath string) error {
	if err := os.Remove(kong.ExpandPath(filePath)); err != nil {
		panic(err)
	}
	return nil
}

func refreshTokenIfNecessary(config *userConfig) {
	if config.Data.Spotify.RefreshToken != "" && config.Data.Spotify.Expiry.Before(time.Now()) {
		sc := &spotify.SpotifyClient{
			Auth: config.Data.Spotify,
		}

		if err := sc.RefreshToken(); err != nil {
			config.DeleteUserConfig()
			panic(err)
		}

		config.Data.Spotify = sc.Auth

		writeToFile(userConfigPath, config.Data)
	}
}

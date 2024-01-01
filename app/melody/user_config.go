package app

import (
	"log"
	"os"
	"time"

	"github.com/alecthomas/kong"
	"github.com/ayberktandogan/melody/internal/spotify"
	"golang.org/x/oauth2"
	"gopkg.in/yaml.v3"
)

type userConfig struct {
	Spotify oauth2.Token
}

const userConfigFolder = "~/.config/melody"
const userConfigPath = userConfigFolder + "/config"

var defaultUserConfig = &userConfig{
	Spotify: oauth2.Token{
		RefreshToken: "",
		AccessToken:  "",
		Expiry:       time.Now(),
		TokenType:    "",
	},
}

var UserConfig = &userConfig{}

func LoadUserConfig() (userConfig, error) {
	createDirIfNotExists()
	createFileIfNotExists(userConfigPath)

	config := readFromFile(userConfigPath)
	refreshTokenIfNecessary(&config)

	UserConfig = &config

	return *UserConfig, nil
}

func SaveUserConfig() error {
	err := writeToFile(userConfigPath, UserConfig)
	if err != nil {
		log.Fatal(err)
		panic(err)
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

func readFromFile(filePath string) (userConfig userConfig) {
	f, err := os.ReadFile(kong.ExpandPath(filePath))
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	ferr := yaml.Unmarshal(f, &userConfig)
	if ferr != nil {
		log.Fatal(ferr)
		panic(ferr)
	}

	return
}

func writeToFile(filePath string, d any) error {
	f, err := os.OpenFile(kong.ExpandPath(filePath), os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	c, err := yaml.Marshal(d)
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
	if config.Spotify.Expiry.Before(time.Now()) {
		sc := &spotify.SpotifyClient{
			Auth: config.Spotify,
		}

		if err := sc.RefreshToken(); err != nil {
			deleteFile(userConfigPath)
			panic(err)
		}

		config.Spotify = sc.Auth

		writeToFile(userConfigPath, config)
	}
}

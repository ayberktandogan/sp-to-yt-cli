package app

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/alecthomas/kong"
	"github.com/ayberktandogan/melody/internal/spotify"
	"gopkg.in/yaml.v3"
)

type UserConfig struct {
	Spotify spotify.AccessTokenResponse
}

const userConfigPath = "~/.config/melody"

var defaultUserConfig = UserConfig{
	Spotify: spotify.AccessTokenResponse{
		RefreshToken: "",
		AccessToken:  "",
		Scope:        "",
		ExpiresIn:    0,
		TokenType:    "",
	},
}

func LoadUserConfig() (UserConfig, error) {
	configFile := filepath.Join(userConfigPath, "config")

	createDirIfNotExists()
	createFileIfNotExists(configFile)

	config := readFromFile(configFile)

	return config, nil
}

func createDirIfNotExists() {
	err := os.MkdirAll(kong.ExpandPath(userConfigPath), os.ModePerm)
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
		c, err := yaml.Marshal(defaultUserConfig)
		if err != nil {
			log.Fatal(err)
			panic(err)
		}

		b, ferr := f.Write(c)
		if ferr != nil {
			log.Fatal(ferr)
			panic(ferr)
		}
		fmt.Println(b)
	}
}

func readFromFile(filePath string) (userConfig UserConfig) {
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

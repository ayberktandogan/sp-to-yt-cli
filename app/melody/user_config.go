package app

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/alecthomas/kong"
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

	c, err := yaml.Marshal(d)
	fmt.Println(c)
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

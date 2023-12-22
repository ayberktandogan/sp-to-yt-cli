package config

import (
	"log"
	"os"
	"strconv"
)

type Spotify struct {
	ClientId     string
	ClientSecret string
	AuthorizeUrl string
	Scopes       string
	RedirectUri  string
}

type System struct {
	Port        int
	Environment string
	Version     string
}

type EnvConfig struct {
	Spotify
	System
}

var (
	Config  *EnvConfig
	version *string
)

func InitEnvConfigs(v string) {
	version = &v
	Config = loadEnvVariables()
}

// func loadEnvVariables() (config *envConfig) {
// Can't do unmarshall from env variables as BindStruct is broken
// and unusable, and I have no way of binding keys without getting the err
// '' expected a map, got 'ptr'
// Will comeback to this if it ever gets fixed
// https://github.com/spf13/viper/issues/1706

// viper.SetEnvPrefix("MELODY")
// viper.AutomaticEnv()

// // Viper unmarshals the loaded env varialbes into the struct
// if err := viper.Unmarshal(&config); err != nil {
// 	log.Fatal(err)
// }

// 	return
// }

func loadEnvVariables() (config *EnvConfig) {
	p, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		log.Fatalf("%v", err)
	}

	config = &EnvConfig{
		Spotify: Spotify{
			ClientId:     os.Getenv("SPOTIFY_CLIENT_ID"),
			ClientSecret: os.Getenv("SPOTIFY_CLIENT_SECRET"),
			AuthorizeUrl: os.Getenv("SPOTIFY_AUTH_URL"),
			Scopes:       os.Getenv("SPOTIFY_SCOPES"),
			RedirectUri:  os.Getenv("SPOTIFY_REDIRECT_URI"),
		},
		System: System{
			Port:        p,
			Environment: os.Getenv("ENVIRONMENT"),
			Version:     *version,
		},
	}

	return
}

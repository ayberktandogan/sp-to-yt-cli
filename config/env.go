package config

type Spotify struct {
	ClientId     string
	AuthorizeUrl string
	TokenUrl     string
	Scopes       []string
	RedirectUri  string
	BaseAPIUri   string
}

type System struct {
	Port        string
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

func InitConfig(v string) {
	version = &v
	Config = loadEnvVariables()
}

func loadEnvVariables() (config *EnvConfig) {
	config = &EnvConfig{
		Spotify: Spotify{
			ClientId:     "b536e6ccfe114da181340c67e2ff4831",
			AuthorizeUrl: "https://accounts.spotify.com/authorize",
			TokenUrl:     "https://accounts.spotify.com/api/token",
			Scopes:       []string{"user-read-private", "user-read-email", "user-top-read", "playlist-modify-public", "playlist-modify-private", "user-follow-read", "user-follow-modify"},
			RedirectUri:  "http://127.0.0.1:8080/auth/callback",
			BaseAPIUri:   "https://api.spotify.com/v1",
		},
		System: System{
			Port:        "8080",
			Environment: "LOCAL",
			Version:     *version,
		},
	}

	return
}

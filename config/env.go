package config

type Spotify struct {
	ClientId     string
	AuthorizeUrl string
	TokenUrl     string
	Scopes       []string
	RedirectUri  string
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
			Scopes:       make([]string, 0),
			RedirectUri:  "http://localhost:8080/auth/callback",
		},
		System: System{
			Port:        "8080",
			Environment: "LOCAL",
			Version:     *version,
		},
	}

	return
}

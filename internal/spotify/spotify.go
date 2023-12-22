package spotify

type AccessTokenResponse struct {
	AccessToken  string `yaml:"spotify.accesstoken,omitempty"`
	TokenType    string `yaml:"spotify.tokentype,omitempty"`
	Scope        string `yaml:"spotify.scope,omitempty"`
	ExpiresIn    int    `yaml:"spotify.expiresin,omitempty"`
	RefreshToken string `yaml:"spotify.refreshtoken,omitempty"`
}

type SpotifyClient struct {
	AccessToken AccessTokenResponse
}

func (s *SpotifyClient) Login() {

}

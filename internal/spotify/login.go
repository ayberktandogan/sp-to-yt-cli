package spotify

import (
	"context"
	"fmt"
	"log"

	"github.com/ayberktandogan/melody/config"
	"github.com/ayberktandogan/melody/internal/utils"
	"golang.org/x/oauth2"
)

type SpotifyClient[T any] struct {
	Auth oauth2.Token
}

func (s *SpotifyClient[T]) Login() (*oauth2.Token, error) {
	ctx := context.Background()

	st := utils.StateGenerator()
	ver := oauth2.GenerateVerifier()

	res := make(chan string)
	err := make(chan error)

	conf := &oauth2.Config{
		ClientID: config.Config.ClientId,
		Scopes:   config.Config.Scopes,
		Endpoint: oauth2.Endpoint{
			AuthURL:  config.Config.AuthorizeUrl,
			TokenURL: config.Config.TokenUrl,
		},
		RedirectURL: config.Config.RedirectUri,
	}

	u := conf.AuthCodeURL(st, oauth2.AccessTypeOffline, oauth2.S256ChallengeOption(ver))
	fmt.Printf("Waiting for you to complete the login process, if a new tab didn't open, please use this link: %v \n", u)
	utils.OpenUrl(u)

	go openWebServerForCallback(st, res, err)

	select {
	case code := <-res:
		return requestAccessCode(ctx, conf, ver, code)
	case rerr := <-err:
		return nil, rerr
	}
}

func (s *SpotifyClient[T]) RefreshToken(token *oauth2.Token) (*oauth2.Token, error) {
	ctx := context.Background()
	conf := *&oauth2.Config{}
	src := conf.TokenSource(ctx, token)
	newToken, err := src.Token() // this actually goes and renews the tokens
	if err != nil {
		panic(err)
	}
	if newToken.AccessToken != token.AccessToken {
		return newToken, nil
	}
	return token, nil
}

func requestAccessCode(ctx context.Context, conf *oauth2.Config, verifier string, code string) (*oauth2.Token, error) {
	tok, err := conf.Exchange(ctx, code, oauth2.VerifierOption(verifier))
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return tok, nil
}

package app

import (
	"fmt"
	"log"

	"github.com/ayberktandogan/melody/internal/spotify"
)

type userTopItemsCmd struct {
	TopItems string `help:"Get the current user's top artists or tracks based on calculated affinity."`
}

type userMeCmd struct {
	Me string `help:"Get data about logged in user"`
}

type userCmd struct {
	Me       userMeCmd       `cmd:"" name:"me" help:"Get data about logged in user"`
	TopItems userTopItemsCmd `cmd:"" name:"top-items" short:"t" help:"Get the current user's top artists or tracks based on calculated affinity."`
}

func (i *userMeCmd) Run() error {
	sc := spotify.SpotifyClient{}
	res, err := sc.GetMe(UserConfig.Spotify)
	if err != nil {
		log.Fatal(err)
		return err
	}

	fmt.Printf("Username: %s \nEmail: %s\nUser's subscription level: %s\n", res.DisplayName, res.Email, res.Product)

	return nil
}

// func (i *userTopItemsCmd) Run() error {
// 	sc := spotify.SpotifyClient{}
// 	res, err := sc.GetMe()
// 	if err != nil {
// 		log.Fatal(err)
// 		return err
// 	}

// 	fmt.Println(res)
// 	return nil
// }

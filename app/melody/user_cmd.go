package app

import (
	"fmt"
	"log"

	"github.com/ayberktandogan/melody/internal/spotify"
)

type userTopArtistsCmd struct {
}

type userTopTracksCmd struct {
}

type userTopItemsCmd struct {
	Artists userTopArtistsCmd `cmd:"" group:"user" help:"Get the current user's top artists based on calculated affinity."`
	Tracks  userTopTracksCmd  `cmd:"" group:"user" help:"Get the current user's top tracks based on calculated affinity."`
}

type userMeCmd struct {
	Me string `help:"Get data about logged in user"`
}

type userCmd struct {
	Me       userMeCmd       `cmd:"" group:"user" name:"me" help:"Get data about logged in user"`
	TopItems userTopItemsCmd `cmd:"" group:"user" name:"top" short:"t" help:"Get the current user's top artists or tracks based on calculated affinity."`
}

func (i *userMeCmd) Run() error {
	sc := spotify.SpotifyClient[spotify.SpotifyUser]{
		Auth: UserConfig.Spotify,
	}
	res, err := sc.GetMe()
	if err != nil {
		log.Fatal(err)
		return err
	}

	fmt.Printf("Username: %s \nEmail: %s\nUser's subscription level: %s\n", res.DisplayName, res.Email, res.Product)

	return nil
}

func (i *userTopArtistsCmd) Run() error {
	sc := spotify.SpotifyClient[spotify.SpotifyUserTopItems[spotify.ArtistObject]]{
		Auth: UserConfig.Spotify,
	}
	res, err := sc.GetUserTopItems("artists")
	if err != nil {
		log.Fatal(err)
		return err
	}

	for idx, it := range res.Items {
		fmt.Printf("%d:\t%s\n", idx+1, it.Name)
	}

	return nil
}

func (i *userTopTracksCmd) Run() error {
	sc := spotify.SpotifyClient[spotify.SpotifyUserTopItems[spotify.TrackObject]]{
		Auth: UserConfig.Spotify,
	}
	res, err := sc.GetUserTopItems("tracks")
	if err != nil {
		log.Fatal(err)
		return err
	}

	for idx, it := range res.Items {
		fmt.Printf("%d:\t%s - %s\n", idx+1, it.Artists[0].Name, it.Name)
	}

	return nil
}

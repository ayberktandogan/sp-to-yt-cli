package app

import (
	"fmt"
	"log"

	"github.com/alecthomas/kong"
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

type userProfileCmd struct {
	Id string `arg:"" help:"Get public profile information about a Spotify user."`
}

type userPlaylistFollowCmd struct {
	Id string `arg:"" help:"Add the current user as a follower of a playlist."`
}

type userPlaylistUnfollowCmd struct {
	Id string `arg:"" help:"Remove the current user as a follower of a playlist."`
}

type userPlaylistCmd struct {
	Follow   userPlaylistFollowCmd   `cmd:"" help:"Add the current user as a follower of a playlist."`
	Unfollow userPlaylistUnfollowCmd `cmd:"" help:"Remove the current user as a follower of a playlist."`
}

type userCmd struct {
	Me       userMeCmd       `cmd:"" group:"user" name:"me" help:"Get data about logged in user"`
	TopItems userTopItemsCmd `cmd:"" group:"user" name:"top" short:"t" help:"Get the current user's top artists or tracks based on calculated affinity."`
	Playlist userPlaylistCmd `cmd:"" group:"user" name:"playlist" short:"p" help:"Playlist related actions"`
	Id       userProfileCmd  `arg:"" help:"Get public profile information about a Spotify user."`
}

func (i *userCmd) Run(ctx *kong.Context, clients *Clients) error {
	if len(ctx.Args) > 2 || ctx.Args[1] == "me" {
		return nil
	}

	data, _, err := clients.Spotify.GetProfile(ctx.Args[1])
	if err != nil {
		log.Fatal(err)
		return err
	}

	fmt.Printf("Username: %s\nId: %s\nImages: %s\n", data.DisplayName, data.Id, data.Images[0].Url)

	return nil
}

func (i *userMeCmd) Run(clients *Clients) error {
	data, _, err := clients.Spotify.GetMe()
	if err != nil {
		log.Fatal(err)
		return err
	}

	fmt.Printf("Username: %s \nEmail: %s\nUser's subscription level: %s\n", data.DisplayName, data.Email, data.Product)

	return nil
}

func (i *userTopArtistsCmd) Run(clients *Clients) error {
	data, _, err := clients.Spotify.GetUserTopArtists()
	if err != nil {
		log.Fatal(err)
		return err
	}

	for idx, it := range data.Items {
		fmt.Printf("%d:\t%s\n", idx+1, it.Name)
	}

	return nil
}

func (i *userTopTracksCmd) Run(clients *Clients) error {
	data, _, err := clients.Spotify.GetUserTopTracks()
	if err != nil {
		log.Fatal(err)
		return err
	}

	for idx, it := range data.Items {
		fmt.Printf("%d:\t%s - %s\n", idx+1, it.Artists[0].Name, it.Name)
	}

	return nil
}

func (i *userPlaylistFollowCmd) Run(ctx *kong.Context, clients *Clients) error {
	_, res, err := clients.Spotify.FollowPlaylist(ctx.Args[len(ctx.Args)-1])
	if err != nil {
		log.Fatal(err)
		return err
	}

	if res.StatusCode == 200 {
		fmt.Println("Followed successfully.")
	} else {
		fmt.Println("There was a problem.")
	}

	return nil
}

func (i *userPlaylistUnfollowCmd) Run(ctx *kong.Context, clients *Clients) error {
	_, res, err := clients.Spotify.UnfollowPlaylist(ctx.Args[len(ctx.Args)-1])
	if err != nil {
		log.Fatal(err)
		return err
	}

	if res.StatusCode == 200 {
		fmt.Println("Unfollowed successfully.")
	} else {
		fmt.Println("There was a problem.")
	}

	return nil
}

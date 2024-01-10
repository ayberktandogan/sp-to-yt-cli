package app

import (
	"fmt"
	"log"
	"strconv"

	"github.com/alecthomas/kong"
	"github.com/ayberktandogan/melody/internal/spotify"
)

type userTopArtistsCmd struct {
	TimeRange string `flag:"" help:"Over what time frame the affinities are computed. Valid values: long_term (calculated from several years of data and including all new data as it becomes available), medium_term (approximately last 6 months), short_term (approximately last 4 weeks). Default: medium_term"`
	Limit     int    `flag:"" help:"The maximum number of items to return. Default: 20. Minimum: 1. Maximum: 50."`
	Offset    int    `flag:"" help:"The index of the first item to return. Default: 0 (the first item). Use with limit to get the next set of items."`
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
	Id     string `arg:"" help:"Add the current user as a follower of a playlist."`
	Public bool   `flag:"" help:"Defaults to true. If true the playlist will be included in user's public playlists, if false it will remain private."`
}

type userPlaylistUnfollowCmd struct {
	Id string `arg:"" help:"Remove the current user as a follower of a playlist."`
}

type userPlaylistCmd struct {
	Follow   userPlaylistFollowCmd   `cmd:"" help:"Add the current user as a follower of a playlist."`
	Unfollow userPlaylistUnfollowCmd `cmd:"" help:"Remove the current user as a follower of a playlist."`
}

type userFollowedArtistsCmd struct {
	Type  string `flag:"" help:"The ID type: currently only artist is supported."`
	After string `flag:"" help:"The last artist ID retrieved from the previous request."`
	Limit int    `flag:"" help:"The maximum number of items to return. Default: 20. Minimum: 1. Maximum: 50."`
}

type userCmd struct {
	Me              userMeCmd              `cmd:"" group:"user" name:"me" help:"Get data about logged in user"`
	TopItems        userTopItemsCmd        `cmd:"" group:"user" name:"top" short:"t" help:"Get the current user's top artists or tracks based on calculated affinity."`
	Playlist        userPlaylistCmd        `cmd:"" group:"user" name:"playlist" short:"p" help:"Playlist related actions"`
	FollowedArtists userFollowedArtistsCmd `cmd:"" group:"user" name:"followed-artists" short:"f" help:"Get the current user's followed artists."`
	Get             userProfileCmd         `cmd:"" help:"Get public profile information about a Spotify user."`
}

func (i *userProfileCmd) Run(ctx *kong.Context, clients *Clients) error {
	data, _, err := clients.Spotify.GetProfile(i.Id)
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
	b := &spotify.SpotifyPlaylistFollowRequest{
		Public: i.Public,
	}

	_, res, err := clients.Spotify.FollowPlaylist(i.Id, b)
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
	_, res, err := clients.Spotify.UnfollowPlaylist(i.Id)
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

func (i *userFollowedArtistsCmd) Run(clients *Clients) error {
	clients.Spotify.QueryParams = map[string]string{
		"type":  i.Type,
		"after": i.After,
		"limit": strconv.Itoa(i.Limit),
	}

	data, _, err := clients.Spotify.GetFollowedArtists()
	if err != nil {
		log.Fatal(err)
		return err
	}

	if len(data.Artists.Items) == 0 {
		fmt.Printf("You're not following any artists.\n")
		return nil
	}

	for idx, it := range data.Artists.Items {
		fmt.Printf("%d:\t%s\n", idx+1, it.Name)
	}

	return nil
}

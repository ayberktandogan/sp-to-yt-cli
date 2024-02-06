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
	userTopArtistsCmd
}

type userMeCmd struct {
	Me string `help:"Get data about logged in user"`
}

type userProfileCmd struct {
	Id string `arg:"" help:"Get public profile information about a Spotify user."`
}

type userPlaylistFollowCmd struct {
	Id     string `arg:"" help:"The Spotify ID of the playlist."`
	Public bool   `flag:"" help:"Defaults to true. If true the playlist will be included in user's public playlists, if false it will remain private."`
}

type userPlaylistUnfollowCmd struct {
	Id string `arg:"" help:"The Spotify ID of the playlist."`
}

type userPlaylistCheckCmd struct {
	Id string `arg:"" help:"The Spotify ID of the playlist."`
}

type userPlaylistCmd struct {
	Check    userPlaylistCheckCmd    `cmd:"" help:"Check to see if one or more Spotify users are following a specified playlist."`
	Follow   userPlaylistFollowCmd   `cmd:"" help:"Add the current user as a follower of a playlist."`
	Unfollow userPlaylistUnfollowCmd `cmd:"" help:"Remove the current user as a follower of a playlist."`
}

type userArtistsFollowedCmd struct {
	Type  string `flag:"" help:"The ID type: currently only artist is supported."`
	After string `flag:"" help:"The last artist ID retrieved from the previous request."`
	Limit int    `flag:"" help:"The maximum number of items to return. Default: 20. Minimum: 1. Maximum: 50."`
}

type userArtistsFollowCmd struct {
	Ids string `flag:"" help:"A comma-separated list of the artist Spotify IDs. A maximum of 50 IDs can be sent in one request."`
}

type userFollowCmd struct {
	Check userFollowCheckCmd `cmd:"" help:"Check to see if the current user is following one or more other Spotify users."`
	Ids   string             `flag:"" help:"A comma-separated list of the user Spotify IDs. A maximum of 50 IDs can be sent in one request."`
}

type userFollowCheckCmd struct {
	Ids string `flag:"" help:"A comma-separated list of the user Spotify IDs. A maximum of 50 IDs can be sent in one request."`
}

type userUnfollowCmd struct {
	Ids string `flag:"" help:"A comma-separated list of the user Spotify IDs. A maximum of 50 IDs can be sent in one request."`
}

type userArtistsCmd struct {
	Top      userTopArtistsCmd      `cmd:"" help:"Get the current user's top artists based on calculated affinity."`
	Follow   userArtistsFollowCmd   `cmd:"" help:"Add the current user as a follower of one or more artists."`
	Followed userArtistsFollowedCmd `cmd:"" group:"user" name:"followed-artists" help:"Get the current user's followed artists."`
	Unfollow userArtistsUnfollowCmd `cmd:"" help:"Remove the current user as a follower of one or more artists."`
}

type userArtistsUnfollowCmd struct {
	Ids string `flag:"" help:"A comma-separated list of the artist Spotify IDs. A maximum of 50 IDs can be sent in one request."`
}

type userTracksCmd struct {
	Top userTopTracksCmd `cmd:"" group:"user" help:"Get the current user's top tracks based on calculated affinity."`
}

type userCmd struct {
	Artists  userArtistsCmd  `cmd:"" group:"user" name:"artists" short:"a" help:"Get data about artists"`
	Tracks   userTracksCmd   `cmd:"" group:"user" name:"tracks" short:"t" help:"Get data about tracks"`
	Me       userMeCmd       `cmd:"" group:"user" name:"me" help:"Get data about logged in user"`
	Playlist userPlaylistCmd `cmd:"" group:"user" name:"playlist" short:"p" help:"Playlist related actions"`
	Follow   userFollowCmd   `cmd:"" group:"user" name:"follow" short:"f" help:"Add the current user as a follower of one or more artists or other Spotify users."`
	Unfollow userUnfollowCmd `cmd:"" group:"user" name:"unfollow" short:"u" help:"Remove the current user as a follower of one or more other Spotify users."`
	Get      userProfileCmd  `cmd:"" help:"Get public profile information about a Spotify user."`
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
	clients.Spotify.QueryParams = map[string]string{
		"time_range": i.TimeRange,
		"limit":      strconv.Itoa(i.Limit),
		"offset":     strconv.Itoa(i.Offset),
	}

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
	clients.Spotify.QueryParams = map[string]string{
		"time_range": i.TimeRange,
		"limit":      strconv.Itoa(i.Limit),
		"offset":     strconv.Itoa(i.Offset),
	}

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

	_, _, err := clients.Spotify.FollowPlaylist(i.Id, b)
	if err != nil {
		log.Fatal(err)
		return err
	}

	fmt.Println("Followed successfully.")

	return nil
}

func (i *userPlaylistUnfollowCmd) Run(ctx *kong.Context, clients *Clients) error {
	_, _, err := clients.Spotify.UnfollowPlaylist(i.Id)
	if err != nil {
		log.Fatal(err)
		return err
	}

	fmt.Println("Unfollowed successfully.")

	return nil
}

func (i *userArtistsFollowedCmd) Run(clients *Clients) error {
	clients.Spotify.QueryParams = map[string]string{
		"type":  "artist",
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

func (i *userFollowCmd) Run(clients *Clients) error {
	clients.Spotify.QueryParams = map[string]string{
		"type": "user",
		"ids":  i.Ids,
	}

	_, _, err := clients.Spotify.FollowUsersOrArtists()
	if err != nil {
		log.Fatal(err)
		return err
	}

	fmt.Println("Followed successfully.")

	return nil
}

func (i *userUnfollowCmd) Run(clients *Clients) error {
	clients.Spotify.QueryParams = map[string]string{
		"type": "user",
		"ids":  i.Ids,
	}

	_, _, err := clients.Spotify.UnfollowUsersOrArtists()
	if err != nil {
		log.Fatal(err)
		return err
	}

	fmt.Println("Unfollowed successfully.")

	return nil
}

func (i *userArtistsFollowCmd) Run(ctx *kong.Context, clients *Clients) error {
	clients.Spotify.QueryParams = map[string]string{
		"type": "artist",
		"ids":  i.Ids,
	}

	_, _, err := clients.Spotify.FollowUsersOrArtists()
	if err != nil {
		log.Fatal(err)
		return err
	}

	fmt.Println("Followed successfully.")

	return nil
}

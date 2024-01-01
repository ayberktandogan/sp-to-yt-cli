package spotify

import (
	"encoding/json"
	"log"
	"net/http"
)

type explicitContent struct {
	FilterEnabled bool `json:"filter_enabled"`
	FilterLocked  bool `json:"filter_locked"`
}

type externalUrls struct {
	Spotify string `json:"spotify"`
}

type followers struct {
	Href  string `json:"href"`
	Total int    `json:"total"`
}

type image struct {
	Url    string `json:"url"`
	Height int    `json:"height"`
	Width  int    `json:"width"`
}

type restrictions struct {
	Reason string `json:"reason"`
}

type album struct {
	AlbumType            string                   `json:"album_type"`
	TotalTracks          int                      `json:"total_tracks"`
	AvailableMarkets     []string                 `json:"available_markets"`
	ExternalUrls         externalUrls             `json:"external_urls"`
	Href                 string                   `json:"href"`
	Id                   string                   `json:"id"`
	Images               []image                  `json:"images"`
	Name                 string                   `json:"name"`
	ReleaseDate          string                   `json:"release_date"`
	ReleaseDatePrecision string                   `json:"release_date_precision"`
	Restrictions         restrictions             `json:"restrictions"`
	Type                 string                   `json:"type"`
	Uri                  string                   `json:"uri"`
	Artists              []simplifiedArtistObject `json:"artists"`
}

type SpotifyUser struct {
	Country         string          `json:"country,omitempty"`
	DisplayName     string          `json:"display_name"`
	Email           string          `json:"email,omitempty"`
	ExplicitContent explicitContent `json:"explicit_content,omitempty"`
	ExternalUrls    externalUrls    `json:"external_urls"`
	Followers       followers       `json:"followers"`
	Href            string          `json:"href"`
	Id              string          `json:"id"`
	Images          []image         `json:"images"`
	Product         string          `json:"product,omitempty"`
	Type            string          `json:"type"`
	Uri             string          `json:"uri"`
}

type simplifiedArtistObject struct {
	ExternalUrls externalUrls `json:"external_urls"`
	Href         string       `json:"href"`
	Id           string       `json:"id"`
	Name         string       `json:"name"`
	Type         string       `json:"type"`
	Uri          string       `json:"uri"`
}

type ArtistObject struct {
	simplifiedArtistObject
	Followers  followers `json:"followers"`
	Genres     []string  `json:"genres"`
	Images     []image   `json:"images"`
	Popularity int       `json:"popularity"`
}

type externalIds struct {
	Isrc string `json:"isrc"`
	Ean  string `json:"ean"`
	Upc  string `json:"upc"`
}

type TrackObject struct {
	Album            album          `json:"album"`
	Artists          []ArtistObject `json:"artists"`
	AvailableMarkets []string       `json:"available_markets"`
	DiscNumber       int            `json:"disc_number"`
	DurationMs       int            `json:"duration_ms"`
	Explicit         bool           `json:"explicit"`
	ExternalIds      externalIds    `json:"external_ids"`
	ExternalUrls     externalUrls   `json:"external_urls"`
	Href             string         `json:"href"`
	Id               string         `json:"id"`
	IsPlayable       bool           `json:"is_playable"`
	LinkedFrom       *TrackObject   `json:"linked_from"`
	Restrictions     restrictions   `json:"restrictions"`
	Name             string         `json:"name"`
	Popularity       int            `json:"popularity"`
	PreviewUrl       string         `json:"preview_url"`
	TrackNumber      int            `json:"track_number"`
	Type             string         `json:"type"`
	Uri              string         `json:"uri"`
	IsLocal          bool           `json:"is_local"`
}

type SpotifyUserTopItems[K any] struct {
	Href     string `json:"href"`
	Limit    int    `json:"limit"`
	Next     string `json:"next"`
	Offset   int    `json:"offset"`
	Previous string `json:"previous"`
	Total    int    `json:"total"`
	Items    []K    `json:"items"`
}

type SpotifyPlaylistFollowRequest struct {
	Public bool `json:"public"`
}

type SpotifyEmptyResponse struct{}

func (s *SpotifyClient) GetMe() (*SpotifyUser, *http.Response, error) {
	return Request[SpotifyUser](s, http.MethodGet, "/me", nil)
}

func (s *SpotifyClient) GetUserTopArtists() (*SpotifyUserTopItems[ArtistObject], *http.Response, error) {
	return Request[SpotifyUserTopItems[ArtistObject]](s, http.MethodGet, "/me/top/artists", nil)
}

func (s *SpotifyClient) GetUserTopTracks() (*SpotifyUserTopItems[TrackObject], *http.Response, error) {
	return Request[SpotifyUserTopItems[TrackObject]](s, http.MethodGet, "/me/top/tracks", nil)
}

func (s *SpotifyClient) GetProfile(id string) (*SpotifyUser, *http.Response, error) {
	return Request[SpotifyUser](s, http.MethodGet, "/users/"+id, nil)
}

func (s *SpotifyClient) FollowPlaylist(id string) (*SpotifyEmptyResponse, *http.Response, error) {
	d := &SpotifyPlaylistFollowRequest{
		Public: false,
	}

	enc, err := json.Marshal(d)
	if err != nil {
		log.Fatal(err)
	}

	return Request[SpotifyEmptyResponse](s, http.MethodPut, "/playlists/"+id+"/followers", enc)
}

func (s *SpotifyClient) UnfollowPlaylist(id string) (*SpotifyEmptyResponse, *http.Response, error) {
	return Request[SpotifyEmptyResponse](s, http.MethodDelete, "/playlists/"+id+"/followers", nil)
}

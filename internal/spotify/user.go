package spotify

import (
	"net/http"
)

func (s *SpotifyClient) GetMe() (*SpotifyUser, *http.Response, error) {
	return Request[SpotifyUser](s, http.MethodGet, "/me")
}

func (s *SpotifyClient) GetUserTopArtists() (*SpotifyUserTopItems[ArtistObject], *http.Response, error) {
	return Request[SpotifyUserTopItems[ArtistObject]](s, http.MethodGet, "/me/top/artists")
}

func (s *SpotifyClient) GetUserTopTracks() (*SpotifyUserTopItems[TrackObject], *http.Response, error) {
	return Request[SpotifyUserTopItems[TrackObject]](s, http.MethodGet, "/me/top/tracks")
}

func (s *SpotifyClient) GetProfile(id string) (*SpotifyUser, *http.Response, error) {
	return Request[SpotifyUser](s, http.MethodGet, "/users/"+id)
}

func (s *SpotifyClient) FollowPlaylist(id string, b *SpotifyPlaylistFollowRequest) (*SpotifyEmptyResponse, *http.Response, error) {
	ParseBody[SpotifyPlaylistFollowRequest](s, b)

	return Request[SpotifyEmptyResponse](s, http.MethodPut, "/playlists/"+id+"/followers")
}

func (s *SpotifyClient) UnfollowPlaylist(id string) (*SpotifyEmptyResponse, *http.Response, error) {
	return Request[SpotifyEmptyResponse](s, http.MethodDelete, "/playlists/"+id+"/followers")
}

func (s *SpotifyClient) GetFollowedArtists() (*SpotifyFollowedArtistsResponse[ArtistObject], *http.Response, error) {
	return Request[SpotifyFollowedArtistsResponse[ArtistObject]](s, http.MethodGet, "/me/following")
}

func (s *SpotifyClient) FollowUsersOrArtists() (*SpotifyEmptyResponse, *http.Response, error) {
	return Request[SpotifyEmptyResponse](s, http.MethodPut, "/me/following")
}

func (s *SpotifyClient) UnfollowUsersOrArtists() (*SpotifyEmptyResponse, *http.Response, error) {
	return Request[SpotifyEmptyResponse](s, http.MethodDelete, "/me/following")
}

func (s *SpotifyClient) CheckUserFollowsUsersOrArtists() (*[]bool, *http.Response, error) {
	return Request[[]bool](s, http.MethodGet, "/me/following/contains")
}

func (s *SpotifyClient) CheckUserFollowsPlaylist(id string) (*[]bool, *http.Response, error) {
	return Request[[]bool](s, http.MethodGet, "/playlists/"+id+"/followers/contains")
}

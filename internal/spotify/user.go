package spotify

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
	Country         string          `json:"country"`
	DisplayName     string          `json:"display_name"`
	Email           string          `json:"email"`
	ExplicitContent explicitContent `json:"explicit_content"`
	ExternalUrls    externalUrls    `json:"external_urls"`
	Followers       followers       `json:"followers"`
	Href            string          `json:"href"`
	Id              string          `json:"id"`
	Images          []image         `json:"images"`
	Product         string          `json:"product"`
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

func (s *SpotifyClient[T]) GetMe() (*T, error) {
	return s.Request("/me")
}

func (s *SpotifyClient[T]) GetUserTopItems(t string) (*T, error) {
	return s.Request("/me/top/" + t)
}

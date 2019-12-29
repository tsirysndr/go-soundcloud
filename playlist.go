package soundcloud

type PlaylistService service

type Playlist struct {
	ArtworkURL    string  `json:"artwork_url,omitempty"`
	CreatedAt     string  `json:"created_at,omitempty"`
	Description   string  `json:"description,omitempty"`
	Downloadable  bool    `json:"downloadable,omitempty"`
	Duration      int     `json:"duration,omitempty"`
	EAN           string  `json:"ean,omitempty"`
	EmbeddableBy  string  `json:"embeddable_by,omitempty"`
	Genre         string  `json:"genre,omitempty"`
	ID            int     `json:"id,omitempty"`
	Kind          string  `json:"kind,omitempty"`
	LabelID       int     `json:"label_id,omitempty"`
	LabelName     string  `json:"label_name,omitempty"`
	LastModified  string  `json:"last_modified,omitempty"`
	License       string  `json:"license,omitempty"`
	LikesCount    int     `json:"likes_count,omitempty"`
	Permalink     string  `json:"permalink,omitempty"`
	PermalinkURL  string  `json:"permalink_url,omitempty"`
	PlaylistType  string  `json:"playlist_type,omitempty"`
	PurchaseTitle string  `json:"purchase_title,omitempty"`
	PurchaseURL   string  `json:"purchase_url,omitempty"`
	Release       int     `json:"release,omitempty"`
	ReleaseDay    int     `json:"release_day,omitempty"`
	ReleaseMonth  int     `json:"release_month,omitempty"`
	ReleaseYear   int     `json:"release_year,omitempty"`
	RepostsCount  int     `json:"reposts_count,omitempty"`
	Sharing       string  `json:"sharing,omitempty"`
	Streamable    bool    `json:"streamable,omitempty"`
	TagList       string  `json:"tag_list,omitempty"`
	Title         string  `json:"title,omitempty"`
	TrackCount    int     `json:"track_count,omitempty"`
	Tracks        []Track `json:"tracks,omitempty"`
	URI           string  `json:"uri,omitempty"`
	User          *User   `json:"user,omitempty"`
}

func (s *PlaylistService) List() ([]Playlist, error) {
	var err error
	playlists := new([]Playlist)
	s.client.base.Get("/playlists").Receive(playlists, err)
	return *playlists, err
}

func (s *PlaylistService) Get(ID string) (*Playlist, error) {
	var err error
	playlist := new(Playlist)
	s.client.base.Path("playlists/").Get(ID).Receive(playlist, err)
	return playlist, err
}

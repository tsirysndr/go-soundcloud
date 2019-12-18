package soundcloud

type PlaylistService service

type Playlist struct {
	ArtworkURL    string  `json:"artwork_url"`
	CreatedAt     string  `json:"created_at"`
	Description   string  `json:"description"`
	Downloadable  bool    `json:"downloadable"`
	Duration      int     `json:"duration"`
	EAN           string  `json:"ean"`
	EmbeddableBy  string  `json:"embeddable_by"`
	Genre         string  `json:"genre"`
	ID            int     `json:"id"`
	Kind          string  `json:"kind"`
	LabelID       int     `json:"label_id"`
	LabelName     string  `json:"label_name"`
	LastModified  string  `json:"last_modified"`
	License       string  `json:"license"`
	LikesCount    int     `json:"likes_count"`
	Permalink     string  `json:"permalink"`
	PermalinkURL  string  `json:"permalink_url"`
	PlaylistType  string  `json:"playlist_type"`
	PurchaseTitle string  `json:"purchase_title"`
	PurchaseURL   string  `json:"purchase_url"`
	Release       int     `json:"release"`
	ReleaseDay    int     `json:"release_day"`
	ReleaseMonth  int     `json:"release_month"`
	ReleaseYear   int     `json:"release_year"`
	RepostsCount  int     `json:"reposts_count"`
	Sharing       string  `json:"sharing"`
	Streamable    bool    `json:"streamable"`
	TagList       string  `json:"tag_list"`
	Title         string  `json:"title"`
	TrackCount    int     `json:"track_count"`
	Tracks        []Track `json:"tracks"`
	URI           string  `json:"uri"`
	User          *User   `json:"user"`
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

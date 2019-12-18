package soundcloud

type TrackService service

type Track struct {
	ArtworkURL          string `json:"artwork_url"`
	Bpm                 int    `json:"bpm"`
	CommentCount        int    `json:"comment_count"`
	Commentable         bool   `json:"commentable"`
	CreatedAt           string `json:"created_at"`
	Description         string `json:"description"`
	DownloadCount       int    `json:"download_count"`
	DownloadURL         string `json:"download_url"`
	Downloadable        bool   `json:"downloadable"`
	Duration            int    `json:"duration"`
	EmbeddableBy        string `json:"embeddable_by"`
	FavoritingsCount    int    `json:"favoritings_count"`
	Genre               string `json:"genre"`
	ID                  int    `json:"id"`
	ISRC                string `json:"isrc"`
	KeySignature        string `json:"key_signature"`
	Kind                string `json:"kind"`
	LabelID             int    `json:"label_id"`
	LabelName           string `json:"label_name"`
	LastModified        string `json:"last_modified"`
	License             string `json:"license"`
	MonetizationModel   string `json:"monetization_model"`
	OriginalContentSize int    `json:"original_content_size"`
	OriginalFormat      string `json:"original_format"`
	Permalink           string `json:"permalink"`
	PermalinkURL        string `json:"permalink_url"`
	PlaybackCount       int    `json:"playback_count"`
	Policy              string `json:"policy"`
	PurchaseTitle       string `json:"purchase_title"`
	PurchaseURL         string `json:"purchase_url"`
	Release             int    `json:"release"`
	ReleaseDay          int    `json:"release_day"`
	ReleaseMonth        int    `json:"release_month"`
	ReleaseYear         int    `json:"release_year"`
	RepostsCount        int    `json:"reposts_count"`
	Sharing             string `json:"sharing"`
	State               string `json:"state"`
	StreamURL           string `json:"stream_url"`
	Streamable          bool   `json:"streamable"`
	TagList             string `json:"tag_list"`
	Title               string `json:"title"`
	TrackType           string `json:"track_type"`
	URI                 string `json:"uri"`
	User                User   `json:"user"`
	UserID              int    `json:"user_id"`
	VideoURL            string `json:"video_url"`
	WaveformURL         string `json:"waveform_url"`
}

func (s *TrackService) List() ([]Track, error) {
	var err error
	tracks := new([]Track)
	s.client.base.Get("/tracks").Receive(tracks, err)
	return *tracks, err
}

func (s *TrackService) Get(ID string) (*Track, error) {
	var err error
	track := new(Track)
	s.client.base.Path("tracks/").Get(ID).Receive(track, err)
	return track, err
}

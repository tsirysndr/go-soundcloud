package soundcloud

type TrackService service

type Track struct {
	ArtworkURL          string `json:"artwork_url,omitempty"`
	Bpm                 int    `json:"bpm,omitempty"`
	CommentCount        int    `json:"comment_count,omitempty"`
	Commentable         bool   `json:"commentable,omitempty"`
	CreatedAt           string `json:"created_at,omitempty"`
	Description         string `json:"description,omitempty"`
	DownloadCount       int    `json:"download_count,omitempty"`
	DownloadURL         string `json:"download_url,omitempty"`
	Downloadable        bool   `json:"downloadable,omitempty"`
	Duration            int    `json:"duration,omitempty"`
	EmbeddableBy        string `json:"embeddable_by,omitempty"`
	FavoritingsCount    int    `json:"favoritings_count,omitempty"`
	Genre               string `json:"genre,omitempty"`
	ID                  int    `json:"id,omitempty"`
	ISRC                string `json:"isrc,omitempty"`
	KeySignature        string `json:"key_signature,omitempty"`
	Kind                string `json:"kind,omitempty"`
	LabelID             int    `json:"label_id,omitempty"`
	LabelName           string `json:"label_name,omitempty"`
	LastModified        string `json:"last_modified,omitempty"`
	License             string `json:"license,omitempty"`
	MonetizationModel   string `json:"monetization_model,omitempty"`
	OriginalContentSize int    `json:"original_content_size,omitempty"`
	OriginalFormat      string `json:"original_format,omitempty"`
	Permalink           string `json:"permalink,omitempty"`
	PermalinkURL        string `json:"permalink_url,omitempty"`
	PlaybackCount       int    `json:"playback_count,omitempty"`
	Policy              string `json:"policy,omitempty"`
	PurchaseTitle       string `json:"purchase_title,omitempty"`
	PurchaseURL         string `json:"purchase_url,omitempty"`
	Release             int    `json:"release,omitempty"`
	ReleaseDay          int    `json:"release_day,omitempty"`
	ReleaseMonth        int    `json:"release_month,omitempty"`
	ReleaseYear         int    `json:"release_year,omitempty"`
	RepostsCount        int    `json:"reposts_count,omitempty"`
	Sharing             string `json:"sharing,omitempty"`
	State               string `json:"state,omitempty"`
	StreamURL           string `json:"stream_url,omitempty"`
	Streamable          bool   `json:"streamable,omitempty"`
	TagList             string `json:"tag_list,omitempty"`
	Title               string `json:"title,omitempty"`
	TrackType           string `json:"track_type,omitempty"`
	URI                 string `json:"uri,omitempty"`
	User                *User  `json:"user,omitempty"`
	UserID              int    `json:"user_id,omitempty"`
	VideoURL            string `json:"video_url,omitempty"`
	WaveformURL         string `json:"waveform_url,omitempty"`
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

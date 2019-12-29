package soundcloud

type UserService service

type User struct {
	AvatarURL    string `json:"avatar_url,omitempty"`
	ID           int    `json:"id,omitempty"`
	Kind         string `json:"kind,omitempty"`
	LastModified string `json:"last_modified,omitempty"`
	Permalink    string `json:"permalink,omitempty"`
	PermalinkURL string `json:"permalink_url,omitempty"`
	URI          string `json:"uri,omitempty"`
	Username     string `json:"username,omitempty"`
}

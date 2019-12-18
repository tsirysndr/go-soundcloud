package soundcloud

type UserService service

type User struct {
	AvatarURL    string `json:"avatar_url"`
	ID           int    `json:"id"`
	Kind         string `json:"kind"`
	LastModified string `json:"last_modified"`
	Permalink    string `json:"permalink"`
	PermalinkURL string `json:"permalink_url"`
	URI          string `json:"uri"`
	Username     string `json:"username"`
}

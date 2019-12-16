package soundcloud

import (
	"github.com/dghubble/sling"
)

type Client struct {
	base       *sling.Sling
	common     service
	Activity   *ActivityService
	App        *AppService
	Comment    *CommentService
	Connection *ConnectionService
	Me         *MeService
	Playlist   *PlaylistService
	Track      *TrackService
	User       *UserService
}

type service struct {
	client *Client
}

func NewClient() *Client {
	base := sling.New().Base("https://api.soundcloud.com/")
	c := &Client{}
	c.base = base
	c.common.client = c
	c.Activity = (*ActivityService)(&c.common)
	c.App = (*AppService)(&c.common)
	c.Comment = (*CommentService)(&c.common)
	c.Connection = (*ConnectionService)(&c.common)
	c.Me = (*MeService)(&c.common)
	c.Playlist = (*PlaylistService)(&c.common)
	c.Track = (*TrackService)(&c.common)
	c.User = (*UserService)(&c.common)
	return c
}

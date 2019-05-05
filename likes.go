package dribbble

import (
	"encoding/json"
	"fmt"
	"time"
)

// Likes client
type Likes struct {
	*Client
}

// LikeOut response structure
type LikeOut struct {
	ID        int       `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Shot      struct {
		ID          int    `json:"id"`
		Title       string `json:"title"`
		Description string `json:"description"`
		Images      struct {
			Hidpi  interface{} `json:"hidpi"`
			Normal string      `json:"normal"`
			Teaser string      `json:"teaser"`
		} `json:"images"`
		PublishedAt time.Time `json:"published_at"`
		HTMLURL     string    `json:"html_url"`
		Height      int       `json:"height"`
		Width       int       `json:"width"`
	} `json:"shot"`
	User struct {
		ID      int    `json:"id"`
		Name    string `json:"name"`
		Login   string `json:"login"`
		HTMLURL string `json:"html_url"`
	} `json:"user"`
}

// LikedShotOut response structure
type LikedShotOut struct {
	ID        int       `json:"id"`
	CreatedAt time.Time `json:"created_at"`
}

// GetLikes returns list of authenticated user’s liked shots
// Note: This is available only to select applications with our approval
func (c *Likes) GetLikes() (out *[]LikeOut, err error) {
	body, err := c.call("GET", "/user/likes", nil)
	if err != nil {
		return nil, err
	}
	defer body.Close()

	err = json.NewDecoder(body).Decode(&out)
	return
}

// GetShotLike checks if you like a shot
// Note: This is available only to select applications with our approval
func (c *Likes) GetShotLike(id int) (out *LikedShotOut, err error) {
	body, err := c.call("GET", fmt.Sprintf("/shots/%d/like", id), nil)
	if err != nil {
		return nil, err
	}
	defer body.Close()

	err = json.NewDecoder(body).Decode(&out)
	return
}

// LikeShot with given id
// Note: This is available only to select applications with our approval
func (c *Likes) LikeShot(id int) (out *LikedShotOut, err error) {
	body, err := c.call("POST", fmt.Sprintf("/shots/%d/like", id), nil)
	if err != nil {
		return nil, err
	}
	defer body.Close()

	err = json.NewDecoder(body).Decode(&out)
	return
}

// UnlikeShot with given id
// Note: This is available only to select applications with our approval
// Unliking a shot requires the user to be authenticated with the write scope
func (c *Likes) UnlikeShot(id int) error {
	body, err := c.call("DELETE", fmt.Sprintf("/shots/%d/like", id), nil)
	if err != nil {
		return err
	}
	defer body.Close()

	return nil
}

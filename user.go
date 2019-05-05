package dribbble

import (
	"encoding/json"
	"time"
)

// User client
type User struct {
	*Client
}

// UserOut response structure
type UserOut struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Login     string `json:"login"`
	HTMLURL   string `json:"html_url"`
	AvatarURL string `json:"avatar_url"`
	Bio       string `json:"bio"`
	Location  string `json:"location"`
	Links     struct {
		Web     string `json:"web"`
		Twitter string `json:"twitter"`
	} `json:"links"`
	CanUploadShot  bool      `json:"can_upload_shot"`
	Pro            bool      `json:"pro"`
	FollowersCount int       `json:"followers_count"`
	CreatedAt      time.Time `json:"created_at"`
	Type           string    `json:"type"`
	Teams          []struct {
		ID        int    `json:"id"`
		Name      string `json:"name"`
		Login     string `json:"login"`
		HTMLURL   string `json:"html_url"`
		AvatarURL string `json:"avatar_url"`
		Bio       string `json:"bio"`
		Location  string `json:"location"`
		Links     struct {
			Web     string `json:"web"`
			Twitter string `json:"twitter"`
		} `json:"links"`
		Type      string    `json:"type"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	} `json:"teams"`
}

// GetUser which is currenlty logged in
func (c *User) GetUser() (out *UserOut, err error) {
	body, err := c.call("GET", "/user", nil)
	if err != nil {
		return nil, err
	}
	defer body.Close()

	err = json.NewDecoder(body).Decode(&out)
	return
}

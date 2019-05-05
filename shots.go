package dribbble

import (
	"encoding/json"
	"fmt"
	"time"
)

// Shots instance
type Shots struct {
	*Client
}

// ShotOut single schema
type ShotOut struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Width       int    `json:"width"`
	Height      int    `json:"height"`
	Images      struct {
		Hidpi  interface{} `json:"hidpi"`
		Normal string      `json:"normal"`
		Teaser string      `json:"teaser"`
	} `json:"images"`
	PublishedAt time.Time `json:"published_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	HTMLURL     string    `json:"html_url"`
	Animated    bool      `json:"animated"`
	Tags        []string  `json:"tags"`
	Attachments []struct {
		ID           int       `json:"id"`
		URL          string    `json:"url"`
		ThumbnailURL string    `json:"thumbnail_url"`
		Size         int       `json:"size"`
		ContentType  string    `json:"content_type"`
		CreatedAt    time.Time `json:"created_at"`
	} `json:"attachments"`
	Projects []struct {
		ID          int       `json:"id"`
		Name        string    `json:"name"`
		Description string    `json:"description"`
		ShotsCount  int       `json:"shots_count"`
		CreatedAt   time.Time `json:"created_at"`
		UpdatedAt   time.Time `json:"updated_at"`
	} `json:"projects"`
	Team struct {
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
	} `json:"team"`
	Video struct {
		ID               int       `json:"id"`
		Duration         int       `json:"duration"`
		VideoFileName    string    `json:"video_file_name"`
		VideoFileSize    int       `json:"video_file_size"`
		Width            int       `json:"width"`
		Height           int       `json:"height"`
		Silent           bool      `json:"silent"`
		CreatedAt        time.Time `json:"created_at"`
		UpdatedAt        time.Time `json:"updated_at"`
		URL              string    `json:"url"`
		SmallPreviewURL  string    `json:"small_preview_url"`
		LargePreviewURL  string    `json:"large_preview_url"`
		XlargePreviewURL string    `json:"xlarge_preview_url"`
	} `json:"video"`
	LowProfile bool `json:"low_profile"`
}

// PopularShotOut schema
type PopularShotOut struct {
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
}

// UpdateShotIn for updating shot
type UpdateShotIn struct {
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Tags        []string `json:"tags"`
}

// GetShots of authenticated user
func (c *Shots) GetShots() (out *[]ShotOut, err error) {
	body, err := c.call("GET", "/user/shots", nil)
	if err != nil {
		return nil, err
	}
	defer body.Close()

	err = json.NewDecoder(body).Decode(&out)
	return
}

// GetPopularShots overall
// Note: This is available only to select applications with our approval
func (c *Shots) GetPopularShots() (out *[]PopularShotOut, err error) {
	body, err := c.call("GET", "/popular_shots", nil)
	if err != nil {
		return nil, err
	}
	defer body.Close()

	err = json.NewDecoder(body).Decode(&out)
	return
}

// GetShot with given id
// This method returns only shots owned by the currently authenticated user
func (c *Shots) GetShot(id int) (out *ShotOut, err error) {
	body, err := c.call("GET", fmt.Sprintf("/shots/%d", id), nil)
	if err != nil {
		return nil, err
	}
	defer body.Close()

	err = json.NewDecoder(body).Decode(&out)
	return
}

// UpdateShot with given id and payload
// Updating a shot requires the user to be authenticated with the upload scope
// The authenticated user must also own the shot
func (c *Shots) UpdateShot(id int, in *UpdateShotIn) (out *ShotOut, err error) {
	body, err := c.call("PUT", fmt.Sprintf("/shots/%d", id), in)
	if err != nil {
		return nil, err
	}
	defer body.Close()

	err = json.NewDecoder(body).Decode(&out)
	return
}

// DeleteShot with given id
// Deleting a shot requires the user to be authenticated with the upload scope
// The authenticated user must also own the shot
func (c *Shots) DeleteShot(id int) error {
	body, err := c.call("DELETE", fmt.Sprintf("/shots/%d", id), nil)
	if err != nil {
		return err
	}
	defer body.Close()

	return nil
}

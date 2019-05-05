package dribbble

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// Client struct
type Client struct {
	*Config
	User        *User
	Projects    *Projects
	Shots       *Shots
	Jobs        *Jobs
	Likes       *Likes
	Attachments *Attachments
}

// New returns new instance of Dribbble client
func New(config *Config) *Client {
	c := &Client{Config: config}
	c.User = &User{c}
	c.Projects = &Projects{c}
	c.Shots = &Shots{c}
	c.Jobs = &Jobs{c}
	c.Likes = &Likes{c}
	c.Attachments = &Attachments{c}
	return c
}

func (c *Client) call(method string, path string, body interface{}) (io.ReadCloser, error) {
	ep := "https://api.dribbble.com/v2" + path
	u, err := url.Parse(ep)
	if err != nil {
		return nil, err
	}

	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+c.AccessToken)
	req.Header.Set("Content-Type", "application/json")

	r, _, err := c.do(req)
	return r, err
}

func (c *Client) do(req *http.Request) (io.ReadCloser, int64, error) {
	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, 0, err
	}

	if res.StatusCode < 400 {
		return res.Body, res.ContentLength, err
	}

	defer res.Body.Close()

	e := &Error{
		StatusCode: res.StatusCode,
		Message:    res.Status,
	}

	ct := res.Header.Get("Content-Type")
	if strings.Contains(ct, "text/html") {
		return nil, 0, e
	}

	if err := json.NewDecoder(res.Body).Decode(e); err != nil {
		return nil, 0, err
	}

	return nil, 0, e
}

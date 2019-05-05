package dribbble

import "net/http"

// Config for Dribbble Auth
type Config struct {
	AccessToken string
	HTTPClient  *http.Client
}

// NewConfig for auth
func NewConfig(accessToken string) *Config {
	return &Config{
		AccessToken: accessToken,
		HTTPClient:  http.DefaultClient,
	}
}

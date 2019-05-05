package dribbble

// Error response.
type Error struct {
	StatusCode int
	Message    string `json:"message"`
}

// Error string.
func (e *Error) Error() string {
	return e.Message
}

package dribbble

import "fmt"

// Attachments client
type Attachments struct {
	*Client
}

// DeleteAttachment requires the user to be authenticated with the upload scope
// The authenticated user must also own the attachment
func (c *Attachments) DeleteAttachment(shotID int, attachmentID int) error {
	body, err := c.call("DELETE", fmt.Sprintf("/shots/%d/attachments/%d", shotID, attachmentID), nil)
	if err != nil {
		return err
	}
	defer body.Close()

	return nil
}

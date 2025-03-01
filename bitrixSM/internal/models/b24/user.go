package b24models

type User struct {
	ID        int    `json:"ID"`
	FirstName string `json:"NAME"`
	LastName  string `json:"LAST_NAME"`
}

func (c *Client) GetUser(userID int) (*User, error) {
	params := map[string]interface{}{"ID": userID}
	var user User
	err := c.CallMethod("user.get", params, &user)
	return &user, err
}

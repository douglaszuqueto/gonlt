package credentials

import "errors"

type Credentials struct {
	Email     string `json:"email"`
	Passwd    string `json:"passwd"`
	Token     string `json:"token"`
	AutoLogin bool   `json:"auto_login"`
}

// Validate credentials
func (c *Credentials) Validate() error {
	if c.Email == "" {
		return errors.New("Email is required")
	}

	if c.Passwd == "" {
		return errors.New("Password is required")
	}

	return nil
}

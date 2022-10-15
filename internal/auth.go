package internal

import "strings"

// auth
// manages to keep authentication data.
type auth struct {
	username string
	password string
}

// authenticate
// checks the user authentication.
func (a *auth) authenticate(token string) bool {
	parts := strings.Split(token, ":")

	return parts[0] == a.username && parts[1] == a.password
}

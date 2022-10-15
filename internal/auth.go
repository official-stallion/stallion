package internal

import "strings"

type auth struct {
	username string
	password string
}

func (a *auth) authenticate(token string) bool {
	parts := strings.Split(token, ":")

	return parts[0] == a.username && parts[1] == a.password
}
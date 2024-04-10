package shared

import (
	"time"
)

const (
	SessionUserKey        = "user"
	SessionAccessTokenKey = "accessToken"
)

type AuthUser struct {
	ID         string
	AuthID     string
	Email      string
	CreatedAt  time.Time
	Active     bool
	IsLoggedIn bool
}

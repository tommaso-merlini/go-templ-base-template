package shared

import (
	"time"
)

type AuthUser struct {
	ID         string
	AuthID     string
	Email      string
	CreatedAt  time.Time
	Active     bool
	IsLoggedIn bool
}

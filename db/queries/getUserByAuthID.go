package queries

import (
	"context"
	"fmt"

	"github.com/tommaso-merlini/go-templ-base-template/db"
)

func GetUserByAuthID(c context.Context, authID string) (db.User, error) {
	var u db.User
	err := db.DB.NewSelect().
		Model(&u).
		Where("auth_id = ?", authID).
		Limit(1).
		Scan(c)
	if err != nil {
		return u, fmt.Errorf("could not find user with auth_id %s: %w", authID, err)
	}
	return u, nil
}

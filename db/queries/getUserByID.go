package queries

import (
	"context"
	"fmt"

	"github.com/tommaso-merlini/go-templ-base-template/db"
)

func GetUserByID(c context.Context, id string) (db.User, error) {
	println(id)
	var u db.User
	err := db.DB.NewSelect().
		Model(&u).
		Where("id = ?", id).
		Limit(1).
		Scan(c)
	if err != nil {
		return u, fmt.Errorf("could not find user with id %s: %w", id, err)
	}
	return u, nil
}

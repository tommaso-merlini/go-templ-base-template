package db

import (
	"context"
	"fmt"

	"github.com/tommaso-merlini/go-templ-base-template/shared"
)

func GetUserByID(c context.Context, id string) (shared.User, error) {
	println(id)
	var u shared.User
	err := DB.NewSelect().
		Model(&u).
		Where("id = ?", id).
		Limit(1).
		Scan(c)
	if err != nil {
		return u, fmt.Errorf("could not find user with id %s: %w", id, err)
	}
	return u, nil
}

func GetUserByAuthID(c context.Context, authID string) (shared.User, error) {
	var u shared.User
	err := DB.NewSelect().
		Model(&u).
		Where("auth_id = ?", authID).
		Limit(1).
		Scan(c)
	if err != nil {
		return u, fmt.Errorf("could not find user with auth_id %s: %w", authID, err)
	}
	return u, nil
}

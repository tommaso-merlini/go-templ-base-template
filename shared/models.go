package shared

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type User struct {
	bun.BaseModel `bun:"table:users,alias:users"`

	ID        uuid.UUID `bun:",pk,type:uuid,default:uuid_generate_v4()"`
	CreatedAt time.Time `bun:",nullzero,notnull,default:now()"`
	Active    bool      `bun:",notnull,default:true"`
	Email     string    `bun:",notnull,unique"`
	AuthID    string    `bun:",nullzero,notnull"`
}

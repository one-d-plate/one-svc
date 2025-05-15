package entity

import (
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/uptrace/bun"
)

var validate = validator.New()

type User struct {
	bun.BaseModel `bun:"table:users"`

	ID        int64      `bun:"id,pk,autoincrement"`
	Username  string     `bun:"username,notnull"`
	Nama      string     `bun:"nama,notnull"`
	Email     string     `bun:"email,notnull,unique"`
	HP        *string    `bun:"hp"`
	Status    *string    `bun:"status"`
	CreatedAt time.Time  `bun:"created_at,notnull,default:current_timestamp"`
	UpdatedAt time.Time  `bun:"updated_at,notnull,default:current_timestamp"`
	DeletedAt *time.Time `bun:"deleted_at,soft_delete"`
}

func (u *User) BeforeInsert() error {
	return validate.Struct(u)
}

func (u *User) BeforeUpdate() error {
	return validate.Struct(u)
}

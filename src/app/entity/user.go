package entity

import (
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/uptrace/bun"
)

var validate = validator.New()

type User struct {
	bun.BaseModel `bun:"users"`

	ID        int64      `bun:"id,pk,autoincrement" json:"id"`
	Username  string     `bun:"username,notnull" json:"username"`
	Nama      string     `bun:"nama,notnull" json:"nama"`
	Email     string     `bun:"email,notnull,unique" json:"email"`
	HP        *string    `bun:"hp" json:"hp"`
	Status    *string    `bun:"status" json:"status"`
	CreatedAt time.Time  `bun:"created_at,notnull,default:current_timestamp" json:"created_at"`
	UpdatedAt time.Time  `bun:"updated_at,notnull,default:current_timestamp" json:"updated_at"`
	DeletedAt *time.Time `bun:"deleted_at,soft_delete" json:"deleted_at,omitempty"`
}

func (u *User) BeforeInsert() error {
	return validate.Struct(u)
}

func (u *User) BeforeUpdate() error {
	return validate.Struct(u)
}

package bootstrap

import "github.com/uptrace/bun"

type Database interface {
	Connect() (*bun.DB, error)
}

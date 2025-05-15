package bootstrap

import (
	"context"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/one-d-plate/one-svc.git/src/configs"
	"github.com/one-d-plate/one-svc.git/src/pkg"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/mysqldialect"
)

type database struct {
	conf configs.DbConfig
	ctx  context.Context
}

func NewDatabase(ctx context.Context) Database {
	dsn := configs.DbConfig{}
	return &database{
		conf: dsn,
		ctx:  ctx,
	}
}

func (d *database) Connect() (*bun.DB, error) {
	sqldb, err := sql.Open("mysql", d.conf.GetDSN())
	if err != nil {
		pkg.Logger.Error("Connection failed ", err)
		return nil, err
	}

	if err := sqldb.Ping(); err != nil {
		pkg.Logger.Error("Database ping failed ", err)
		return nil, err
	}

	db := bun.NewDB(sqldb, mysqldialect.New())
	return db, nil
}

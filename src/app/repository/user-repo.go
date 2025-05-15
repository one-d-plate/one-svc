package repository

import (
	"context"
	"database/sql"
	"errors"

	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/one-d-plate/one-svc.git/src/app/entity"
	"github.com/one-d-plate/one-svc.git/src/app/presentase"
	"github.com/uptrace/bun"
)

type userRepo struct {
	db   *bun.DB
	user *entity.User
}

func NewUserRepo(db *bun.DB, user *entity.User) UserRepository {
	return &userRepo{
		db:   db,
		user: user,
	}
}

func (u *userRepo) GetAll(ctx context.Context, req presentase.GetAllHeader) (*presentase.GetUserResponse, error) {
	users := make([]entity.User, 0)
	query := u.db.NewSelect().
		Model(&users).
		OrderExpr("id DESC").
		Limit(req.Limit).
		Offset(req.Page)

	if req.Search != "" {
		query = query.WhereGroup(" AND ", func(q *bun.SelectQuery) *bun.SelectQuery {
			return q.
				Where("username ILIKE ?", "%"+req.Search+"%").
				WhereOr("nama ILIKE ?", "%"+req.Search+"%").
				WhereOr("email ILIKE ?", "%"+req.Search+"%")
		})
	}

	total, err := query.ColumnExpr("count(*)").ScanAndCount(ctx)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fiber.ErrBadRequest
		}

		return nil, fiber.ErrInternalServerError
	}

	err = query.Scan(ctx)
	if err != nil {
		return nil, err
	}
	meta := presentase.Meta{
		Total: strconv.Itoa(total),
		Page:  strconv.Itoa(req.Page),
		Limit: strconv.Itoa(req.Limit),
	}

	return &presentase.GetUserResponse{
		List: users,
		Meta: meta,
	}, nil
}

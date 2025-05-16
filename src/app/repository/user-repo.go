package repository

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/one-d-plate/one-svc.git/src/app/entity"
	"github.com/one-d-plate/one-svc.git/src/app/presentase"
	"github.com/one-d-plate/one-svc.git/src/pkg"
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

func (u *userRepo) Insert(ctx context.Context, req presentase.CreateUserReq) error {
	user := map[string]interface{}{
		"username":   req.Username,
		"email":      req.Email,
		"nama":       req.Nama,
		"hp":         req.Hp,
		"status":     entity.UserInActive,
		"created_at": time.Now(),
		"updated_at": time.Now(),
	}

	_, err := u.db.NewInsert().
		Model(&user).
		Table("users").
		Exec(ctx)

	return err
}

func (u *userRepo) GetAll(ctx context.Context, req presentase.GetAllHeader) (*presentase.GetUsersResponse, error) {
	users := make([]entity.User, 0)

	baseQuery := u.db.NewSelect().
		Model(&users).
		OrderExpr("id DESC")

	if req.Search != "" {
		baseQuery = baseQuery.WhereGroup(" AND ", func(q *bun.SelectQuery) *bun.SelectQuery {
			return q.
				Where("username ILIKE ?", "%"+req.Search+"%").
				WhereOr("nama ILIKE ?", "%"+req.Search+"%").
				WhereOr("email ILIKE ?", "%"+req.Search+"%")
		})
	}

	// 1. Get total count
	var total int64
	countQuery := baseQuery.Clone()
	err := countQuery.ColumnExpr("count(*)").Scan(ctx, &total)
	if err != nil {
		pkg.LogError("error getting user count", err)
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fiber.ErrBadRequest
		}
		return nil, fiber.ErrInternalServerError
	}

	// 2. Get paged data
	err = baseQuery.Limit(req.Limit).Offset(pkg.GetOffset(req.Page)).Scan(ctx)
	if err != nil {
		return nil, err
	}

	meta := presentase.Meta{
		Total: strconv.FormatInt(total, 10),
		Page:  strconv.Itoa(req.Page),
		Limit: strconv.Itoa(req.Limit),
	}

	return &presentase.GetUsersResponse{
		List: users,
		Meta: meta,
	}, nil
}

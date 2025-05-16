package repository

import (
	"context"
	"database/sql"
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

	pkg.LogError("failed to insert user ", err)
	return fiber.NewError(fiber.StatusBadRequest, "gagal menginput data user")
}

func (u *userRepo) Get(ctx context.Context, req int) (*entity.User, error) {
	user := entity.User{}

	baseQuery := u.db.NewSelect().
		Model(&user).
		Where("id = ?", req).
		Limit(1)

	err := baseQuery.Scan(ctx)
	if err != nil {
		pkg.LogError("failed to get user ", err)
		if err == sql.ErrNoRows {
			err = fiber.NewError(fiber.StatusBadRequest, "user tidak ditemukan")
		}
		return nil, err
	}

	return &user, nil
}

func (u *userRepo) GetAll(ctx context.Context, req presentase.GetAllHeader) (*presentase.GetUsersResponse, error) {
	users := make([]entity.User, 0)

	baseQuery := u.db.NewSelect().
		Model(&users).
		Where("deleted_at IS NULL").
		OrderExpr("id DESC").
		Limit(req.Limit + 1)

	if req.Search != "" {
		baseQuery = baseQuery.WhereGroup(" AND ", func(q *bun.SelectQuery) *bun.SelectQuery {
			return q.
				Where("username LIKE ?", "%"+req.Search).
				WhereOr("nama LIKE ?", "%"+req.Search).
				WhereOr("email LIKE ?", "%"+req.Search)
		})
	}

	if req.Cursor != "" {
		lastID, err := pkg.DecryptCursor(req.Cursor)
		if err != nil {
			pkg.LogError("failed to get all users ", err)
			return nil, fiber.NewError(fiber.StatusBadRequest, "invalid cursor")
		}
		baseQuery = baseQuery.Where("id < ?", lastID)
	}

	err := baseQuery.Scan(ctx)
	if err != nil {
		pkg.LogError("failed to get all users ", err)
		return nil, err
	}

	var nextCursor string
	if len(users) > req.Limit {
		last := users[req.Limit-1]
		nextCursor = pkg.EncryptCursor(last.ID)
		users = users[:req.Limit]
	}

	meta := presentase.Meta{
		Limit:  strconv.Itoa(req.Limit),
		Cursor: nextCursor,
	}

	return &presentase.GetUsersResponse{
		List: users,
		Meta: meta,
	}, nil
}

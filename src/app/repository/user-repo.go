package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
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
		return nil, fiber.NewError(fiber.StatusBadRequest, "data tidak dapat diproses")
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

func (u *userRepo) Update(ctx context.Context, req int, payload presentase.CreateUserReq) error {
	_, err := u.Get(ctx, req)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return fiber.NewError(fiber.StatusBadRequest, "user tidak ditemukan")
		}
		return err // error lain (misalnya koneksi DB)
	}

	user := map[string]interface{}{
		"username":   payload.Username,
		"email":      payload.Email,
		"nama":       payload.Nama,
		"hp":         payload.Hp,
		"status":     payload.Status,
		"created_at": time.Now(),
		"updated_at": time.Now(),
	}

	_, err = u.db.NewUpdate().
		Model(&user).
		Table("users").
		Where("id = ?", req).
		Exec(ctx)

	if err != nil {
		pkg.LogError(fmt.Sprintf("failed to update users id %d", req), err)
		err = fiber.NewError(fiber.StatusBadRequest, "update user gagal")
		return err
	}

	return nil
}

func (u *userRepo) Delete(ctx context.Context, userIDs []int, include bool) error {
	now := time.Now()

	baseQuery := u.db.NewUpdate().
		Model(u.user).
		Set("deleted_at = ?", now)

	if !include {
		go func(ctx context.Context) {
			_, err := baseQuery.
				Where("id NOT IN (?)", bun.In(userIDs)).
				Where("deleted_at IS NULL").
				Exec(ctx)

			if err != nil {
				pkg.LogError(fmt.Sprintf("failed to soft delete users not in %v", userIDs), err)
			} else {
				pkg.LogInfo("goroutine successfully soft-deleted users not in the list")
			}
		}(ctx)

		return nil
	}

	_, err := baseQuery.
		Where("id IN (?)", bun.In(userIDs)).
		Exec(ctx)

	if err != nil {
		pkg.LogError(fmt.Sprintf("failed to soft delete users in %v", userIDs), err)
		return fiber.NewError(fiber.StatusBadRequest, "hapus user gagal")
	}

	return nil
}

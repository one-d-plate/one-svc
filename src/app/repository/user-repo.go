package repository

import (
	"context"
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

	// Base query
	baseQuery := u.db.NewSelect().
		Model(&users).
		OrderExpr("id DESC").
		Limit(req.Limit + 1) // Ambil 1 lebih banyak untuk deteksi nextCursor

	// Filter: Search
	if req.Search != "" {
		baseQuery = baseQuery.WhereGroup(" AND ", func(q *bun.SelectQuery) *bun.SelectQuery {
			return q.
				Where("username LIKE ?", "%"+req.Search).
				WhereOr("nama LIKE ?", "%"+req.Search).
				WhereOr("email LIKE ?", "%"+req.Search)
		})
	}

	// Filter: Cursor
	if req.Cursor != "" {
		lastID, err := pkg.DecryptCursor(req.Cursor)
		if err != nil {
			return nil, fiber.NewError(fiber.StatusBadRequest, "invalid cursor")
		}
		baseQuery = baseQuery.Where("id < ?", lastID)
	}

	// Ambil data
	err := baseQuery.Scan(ctx)
	if err != nil {
		return nil, err
	}

	// Siapkan nextCursor (jika ada data lebih)
	var nextCursor string
	if len(users) > req.Limit {
		last := users[req.Limit-1]
		nextCursor = pkg.EncryptCursor(last.ID)
		users = users[:req.Limit] // Pangkas 1 data ekstra
	}

	// Hitung total (opsional)
	var total int64
	countQuery := u.db.NewSelect().Model((*entity.User)(nil))
	if req.Search != "" {
		countQuery = countQuery.WhereGroup(" AND ", func(q *bun.SelectQuery) *bun.SelectQuery {
			return q.
				Where("username LIKE ?", "%"+req.Search).
				WhereOr("nama LIKE ?", "%"+req.Search).
				WhereOr("email LIKE ?", "%"+req.Search)
		})
	}
	err = countQuery.ColumnExpr("count(*)").Scan(ctx, &total)
	if err != nil {
		pkg.LogError("error getting user count", err)
		return nil, fiber.ErrInternalServerError
	}

	// Meta response
	meta := presentase.Meta{
		Total:  strconv.FormatInt(total, 10),
		Limit:  strconv.Itoa(req.Limit),
		Cursor: nextCursor,
	}

	return &presentase.GetUsersResponse{
		List: users,
		Meta: meta,
	}, nil
}

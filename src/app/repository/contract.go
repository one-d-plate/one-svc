package repository

import (
	"context"

	"github.com/one-d-plate/one-svc.git/src/app/entity"
	"github.com/one-d-plate/one-svc.git/src/app/presentase"
)

type UserRepository interface {
	Insert(ctx context.Context, req presentase.CreateUserReq) error
	Get(ctx context.Context, req int) (*entity.User, error)
	GetAll(ctx context.Context, req presentase.GetAllHeader) (*presentase.GetUsersResponse, error)
	Update(ctx context.Context, req int, payload presentase.CreateUserReq) error
	Delete(ctx context.Context, req []int, include bool) error
}

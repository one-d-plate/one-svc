package repository

import (
	"context"

	"github.com/one-d-plate/one-svc.git/src/app/presentase"
)

type UserRepository interface {
	Insert(ctx context.Context, req presentase.CreateUserReq) error
	GetAll(ctx context.Context, req presentase.GetAllHeader) (*presentase.GetUsersResponse, error)
}

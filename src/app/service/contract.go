package service

import (
	"context"

	"github.com/one-d-plate/one-svc.git/src/app/presentase"
)

type UserService interface {
	Create(ctx context.Context, req presentase.CreateUserReq) error
	Get(ctx context.Context, req int) (*presentase.GetAllResponse, error)
	GetAll(ctx context.Context, req presentase.GetAllHeader) (*presentase.GetAllResponse, error)
}

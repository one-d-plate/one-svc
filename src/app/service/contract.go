package service

import (
	"context"

	"github.com/one-d-plate/one-svc.git/src/app/presentase"
)

type UserService interface {
	GetAll(ctx context.Context, req presentase.GetAllHeader) (*presentase.GetAllResponse, error)
}

package repository

import (
	"context"

	"github.com/one-d-plate/one-svc.git/src/app/presentase"
)

type UserRepository interface {
	GetAll(ctx context.Context, req presentase.GetAllHeader) (*presentase.GetUserResponse, error)
}

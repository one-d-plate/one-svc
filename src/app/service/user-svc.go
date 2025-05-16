package service

import (
	"context"
	"fmt"

	"github.com/one-d-plate/one-svc.git/src/app/presentase"
	"github.com/one-d-plate/one-svc.git/src/app/repository"
	"github.com/one-d-plate/one-svc.git/src/pkg"
)

type userService struct {
	user repository.UserRepository
}

func NewUserService(user repository.UserRepository) UserService {
	return &userService{
		user: user,
	}
}

func (u *userService) Create(ctx context.Context, req presentase.CreateUserReq) error {
	err := u.user.Insert(ctx, req)
	if err != nil {
		pkg.LogError("failed to insert data", err)
		return fmt.Errorf("failed to insert data: %w", err)
	}
	return nil
}

func (u *userService) GetAll(ctx context.Context, req presentase.GetAllHeader) (*presentase.GetAllResponse, error) {
	res, err := u.user.GetAll(ctx, req)
	if err != nil {
		pkg.LogError("failed to fetch user", err)
		return nil, err
	}

	return &presentase.GetAllResponse{
		Message: "success",
		Data:    res,
	}, nil
}

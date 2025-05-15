package service

import (
	"context"
	"errors"

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

func (u *userService) GetAll(ctx context.Context, req presentase.GetAllHeader) (*presentase.GetAllResponse, error) {
	res, err := u.user.GetAll(ctx, req)
	if err != nil {
		pkg.LogError("Failed to fetch user", err)
		return nil, errors.New("user not found")
	}

	return &presentase.GetAllResponse{
		Message: "success",
		Data:    res,
	}, nil
}

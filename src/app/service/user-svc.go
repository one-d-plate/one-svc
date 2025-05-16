package service

import (
	"context"
	"fmt"

	"github.com/one-d-plate/one-svc.git/src/app/entity"
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

func (u *userService) Get(ctx context.Context, req int) (*presentase.GetAllResponse, error) {
	user, err := u.user.Get(ctx, req)
	if err != nil {
		pkg.LogError("failed to get data", err)
		return nil, err
	}
	return &presentase.GetAllResponse{
		Message: "success",
		Data:    user,
	}, nil
}

func (u *userService) GetAll(ctx context.Context, req presentase.GetAllHeader) (*presentase.GetAllResponse, error) {
	res, err := u.user.GetAll(ctx, req)
	if err != nil {
		pkg.LogError("failed to fetch user", err)
		return nil, err
	}

	users := []entity.User{}
	for _, v := range res.List {
		status := ""
		if v.Status != nil {
			status = string(entity.UserStatus[*v.Status])
		}

		newUser := v
		newUser.Status = &status
		users = append(users, newUser)
	}

	res.List = users

	return &presentase.GetAllResponse{
		Message: "success",
		Data:    res,
	}, nil
}

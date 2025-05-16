package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/one-d-plate/one-svc.git/src/app/presentase"
	"github.com/one-d-plate/one-svc.git/src/app/service"
	"github.com/one-d-plate/one-svc.git/src/pkg"
)

type userHandler struct {
	user service.UserService
}

func NewUserHandler(user service.UserService) *userHandler {
	return &userHandler{user: user}
}

func (h *userHandler) CreateUser(c *fiber.Ctx) error {
	var body presentase.CreateUserReq

	if err := c.BodyParser(&body); err != nil {
		return HandleFiberError(c, err)
	}

	err := h.user.Create(c.Context(), body)
	if err != nil {
		return HandleFiberError(c, err)
	}

	pkg.LogInfo("Success to create user")
	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"message": "success",
	})
}

func (h *userHandler) GetUser(c *fiber.Ctx) error {
	param, _ := c.ParamsInt("id")

	user, err := h.user.Get(c.Context(), param)
	if err != nil {
		return HandleFiberError(c, err)
	}
	return c.JSON(user)
}

func (h *userHandler) GetUsers(c *fiber.Ctx) error {
	headers := presentase.GetAllHeader{
		Limit:  c.QueryInt("limit", 10),
		Search: c.Query("search", ""),
		Cursor: c.Query("cursor"),
	}

	users, err := h.user.GetAll(c.Context(), headers)
	if err != nil {
		return HandleFiberError(c, err)
	}
	pkg.LogInfo("Success to get user data")
	return c.JSON(users)
}

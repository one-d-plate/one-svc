package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/one-d-plate/one-svc.git/src/app/presentase"
	"github.com/one-d-plate/one-svc.git/src/app/service"
	"github.com/one-d-plate/one-svc.git/src/pkg"
)

type UserHandler struct {
	user service.UserService
}

func NewUserHandler(user service.UserService) *UserHandler {
	return &UserHandler{user: user}
}

func (h *UserHandler) GetUsers(c *fiber.Ctx) error {
	headers := presentase.GetAllHeader{
		Page:   c.QueryInt("page", 1),
		Limit:  c.QueryInt("limit", 10),
		Search: c.Query("search", ""),
	}

	users, err := h.user.GetAll(c.Context(), headers)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	pkg.LogInfo("Success to get user data")
	return c.JSON(users)
}

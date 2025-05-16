package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/one-d-plate/one-svc.git/src/app/entity"
	"github.com/one-d-plate/one-svc.git/src/app/repository"
	"github.com/one-d-plate/one-svc.git/src/app/service"
	"github.com/one-d-plate/one-svc.git/src/handlers"
	"github.com/uptrace/bun"
)

func RouteRegistry(app *fiber.App, db *bun.DB) {
	userModel := entity.User{}
	userRepo := repository.NewUserRepo(db, &userModel)
	userService := service.NewUserService(userRepo)
	handler := handlers.NewUserHandler(userService)

	app.Get("users", handler.GetUsers)
	app.Post("user", handler.CreateUser)
}

package server

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/one-d-plate/one-svc.git/src/bootstrap"
	"github.com/one-d-plate/one-svc.git/src/pkg"
	"github.com/one-d-plate/one-svc.git/src/route"
)

type server struct{}

func NewServer() Server {
	return &server{}
}

func (s *server) Run() {
	pkg.InitLogger()

	app := fiber.New(fiber.Config{
		Prefork:       false,
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "Fiber",
		AppName:       "auth App v1.0.1",
	})

	// Hanya proses master (parent) yang menangani sinyal
	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, syscall.SIGINT, syscall.SIGTERM)

	if !fiber.IsChild() {
		pkg.WaitForSignalAndShutdown(app)
	}

	// Database dan route hanya dijalankan oleh semua proses (master dan child)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	db := bootstrap.NewDatabase(ctx)
	dbConnect, err := db.Connect()
	if err != nil {
		pkg.Logger.Error("Database connection failed ", err)
		return
	}

	route.RouteRegistry(app, dbConnect)

	if err := app.Listen(":8080"); err != nil {
		pkg.Logger.Fatalf("Failed to start server: %v", err)
	}
}

func (s *server) Done() {
	panic("unimplemented")
}

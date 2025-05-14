package pkg

import (
	"context"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v2"
)

var (
	signalChannel chan os.Signal
	once          sync.Once
)

func SignalInit() {
	once.Do(func() {
		signalChannel = make(chan os.Signal, 1)
		signal.Notify(signalChannel, os.Interrupt, syscall.SIGTERM)
	})
}

func TriggerShutdown() {
	go func() {
		signalChannel <- syscall.SIGTERM
	}()
}

func WaitForSignalAndShutdown(app *fiber.App) {
	go func() {
		sigReceived := <-signalChannel
		Logger.Infof("Received signal: %s. Shutting down gracefully...", sigReceived)

		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()

		app.ShutdownWithContext(ctx)
		Logger.Info("Shutdown complete.")
	}()
}

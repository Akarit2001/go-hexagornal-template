package main

import (
	"context"
	"go-hex-temp/internal/app"
	"go-hex-temp/internal/infrastructure/logx"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	server := app.NewServer()

	go func() {
		if err := server.Start(); err != nil && err.Error() != "http: Server closed" {
			logx.Fatalf("server error: %v", err)
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Stop(ctx); err != nil {
		logx.Fatalf("server shutdown failed: %v", err)
	}

	logx.Info("✅ Server exited properly")
}

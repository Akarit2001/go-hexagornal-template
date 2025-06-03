package app

import (
	"context"
	"fmt"
	userHttp "go-hex-temp/internal/adapters/in/httpx/user"
	"go-hex-temp/internal/core/service"
	"go-hex-temp/internal/infrastructure/config"
	"go-hex-temp/internal/infrastructure/logx"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

type Server struct {
	httpServer *http.Server
	Cfg        *config.Config
	wg         *sync.WaitGroup
}

func NewServer() *Server {

	server := new(Server)
	server.setUp()

	return server
}

func (s *Server) Start() error {
	logx.Infof("🚀 Starting server on http://%s:%s", s.Cfg.Host, s.Cfg.Port)
	return s.httpServer.ListenAndServe()
}

func (s *Server) Stop(ctx context.Context) error {
	logx.Info("🛑 Shutting down server...")
	err := s.httpServer.Shutdown(ctx)
	s.wg.Wait()
	return err
}

func (s *Server) setUp() {
	// === Load Config ===
	cfg := config.Load()

	// === Setup router ===
	r := gin.Default()
	api := r.Group("/api")

	// Example: profile route
	// profileRepo := memory.NewProfileRepository() or mongo.New...
	// profileService := service.NewProfileService(profileRepo)
	// profileHandler := profileHttp.NewHandler(profileService)
	// profileHttp.RegisterRoutes(api.Group("/profile"), profileHandler)
	userService := service.NewUserService(nil, service.NewQCompiler())
	userHandler := userHttp.NewUserHandler(userService)
	userHttp.RegisterRoutes(api.Group("users"), userHandler)

	httpServer := &http.Server{
		Addr:         fmt.Sprintf("%s:%s", cfg.Host, cfg.Port),
		Handler:      r,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  30 * time.Second,
	}

	s.httpServer = httpServer
	s.Cfg = cfg
	s.wg = new(sync.WaitGroup)
}

func (s *Server) RunBackground(f func()) {
	s.wg.Add(1)
	defer s.wg.Done()
	f()
}

package httpServer

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"nats_example/config"
)

type Server struct {
	cfg   *config.Config
	fiber *fiber.App
}

func NewServer(cfg *config.Config) *Server {
	return &Server{
		fiber: fiber.New(fiber.Config{DisableStartupMessage: true}),
		cfg:   cfg,
	}
}

func (s *Server) Run(ctx context.Context) error {
	if err := s.MapHandlers(ctx, s.fiber); err != nil {
		fmt.Sprintf("Cannot map handlers. Error: {%s}", err)
	}

	fmt.Sprintf("Start server on {host:port - %s:%s}", s.cfg.Server.Host, s.cfg.Server.Port)

	if err := s.fiber.Listen(fmt.Sprintf("%s:%s", s.cfg.Server.Host, s.cfg.Server.Port)); err != nil {
		fmt.Sprintf("Cannot listen. Error: {%s}", err)
	}

	return nil
}

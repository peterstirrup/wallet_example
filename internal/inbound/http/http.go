package http

import (
	"context"
	"fmt"
	"net/http"

	"github.com/peterstirrup/wallet_example/internal/domain/entities"
)

type Config struct {
	Port     int
	UseCases UseCases
}

type Server struct {
	addr     string
	handler  http.Handler
	useCases UseCases
}

type UseCases interface {
	GetWallet(ctx context.Context, userID string) (entities.Wallet, error)
}

func NewServer(cfg Config) *Server {
	s := &Server{
		addr:     fmt.Sprintf(":%d", cfg.Port),
		useCases: cfg.UseCases,
	}

	// TODO: Set HTTP handler
	return s
}

func (s *Server) Run(ctx context.Context) error {
	// Implement me!
	return nil
}

func (s *Server) GetWallet(w http.ResponseWriter, r *http.Request) {
	// Implement me!
	// w, err := s.useCases.GetWallet(userID)
}

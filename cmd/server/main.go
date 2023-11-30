package main

import (
	"context"
	"log"

	"github.com/peterstirrup/wallet_example/internal/domain/usecases"
	"github.com/peterstirrup/wallet_example/internal/inbound/http"
	"github.com/peterstirrup/wallet_example/internal/outbound/firestore"
)

var ctx context.Context

func main() {
	s := http.NewServer(http.Config{
		Port: 8000,
		UseCases: usecases.NewWallet(usecases.WalletConfig{
			Store: firestore.New(ctx, "projectID"),
		}),
	})

	if err := s.Run(ctx); err != nil {
		log.Fatal(err)
	}
}

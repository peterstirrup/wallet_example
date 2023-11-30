package main

import (
	"context"
	"log"

	"cloud.google.com/go/pubsub"
	"github.com/peterstirrup/wallet_example/internal/domain/usecases"
	"github.com/peterstirrup/wallet_example/internal/inbound/sub"
	"github.com/peterstirrup/wallet_example/internal/outbound/firestore"
)

var ctx context.Context

func main() {
	s := sub.NewOrderCompleted(sub.Config{
		Sub: &pubsub.Subscription{},
		UseCases: usecases.NewWallet(usecases.WalletConfig{
			Store: firestore.New(ctx, "projectID"),
		}),
	})

	if err := s.Run(ctx); err != nil {
		log.Fatal(err)
	}
}

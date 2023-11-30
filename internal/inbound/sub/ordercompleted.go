package sub

import (
	"context"

	"cloud.google.com/go/pubsub"
)

type Config struct {
	Sub      *pubsub.Subscription
	UseCases UseCases
}

type OrderCompleted struct {
	sub      *pubsub.Subscription
	useCases UseCases
}

type UseCases interface {
	UpdateWalletBalance(ctx context.Context, userID string, amount float64) error
}

func NewOrderCompleted(cfg Config) *OrderCompleted {
	return &OrderCompleted{
		sub:      cfg.Sub,
		useCases: cfg.UseCases,
	}
}

func (s *OrderCompleted) Run(ctx context.Context) error {
	return s.sub.Receive(ctx, func(ctx context.Context, msg *pubsub.Message) {
		defer msg.Ack()

		// Implement me!
		// s.useCases.UpdateWalletBalance(userID, amount)

		return
	})
}

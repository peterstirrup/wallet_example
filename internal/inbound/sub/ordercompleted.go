package sub

import (
	"context"

	"cloud.google.com/go/pubsub"
	"github.com/peterstirrup/wallet_example/internal/domain/entities"
	"github.com/rs/zerolog/log"
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
	HandleOrder(ctx context.Context, order entities.Order) error
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
		var order entities.Order
		if err := s.useCases.HandleOrder(ctx, order); err != nil {
			log.Err(err).Msgf("failed to handle order with ID: %s", order.OrderID)
		}

		return
	})
}

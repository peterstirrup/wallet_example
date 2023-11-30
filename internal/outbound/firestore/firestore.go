package firestore

import (
	"context"

	"cloud.google.com/go/firestore"
	"github.com/peterstirrup/wallet_example/internal/domain/entities"
)

type Client struct {
	fs *firestore.Client
}

func New(ctx context.Context, projectID string) *Client {
	/*
		Messy Firestore implementation logic
	*/
	return &Client{}
}

func (c *Client) GetWallet(ctx context.Context, userID string) (entities.Wallet, error) {
	// Implement me!
	return entities.Wallet{}, nil
}

func (c *Client) UpdateWalletBalance(ctx context.Context, userID string, newBalance float64) error {
	// Implement me!
	return nil
}

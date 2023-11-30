package sql

import (
	"context"
	"database/sql"

	"github.com/peterstirrup/wallet_example/internal/domain/entities"
)

type Client struct {
	mysql *sql.DB
}

func New(ctx context.Context, projectID string) *Client {
	/*
		Messy SQL implementation logic
	*/
	return &Client{}
}

func (c *Client) GetWallet(ctx context.Context, userID string) (entities.Wallet, error) {
	// Implement me!
	return entities.Wallet{}, nil
}

func (c *Client) UpdateWallet(userID string, newBalance float64) error {
	// Implement me!
	return nil
}

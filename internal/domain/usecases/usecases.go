package usecases

import (
	"context"

	"github.com/peterstirrup/wallet_example/internal/domain/entities"
)

type WalletConfig struct {
	Store Store
}

type Wallet struct {
	store Store
}

type Store interface {
	GetWallet(ctx context.Context, userID string) (entities.Wallet, error)
	UpdateWalletBalance(ctx context.Context, userID string, newBalance float64) error
}

func NewWallet(cfg WalletConfig) *Wallet {
	return &Wallet{
		store: cfg.Store,
	}
}

func (w *Wallet) GetWallet(ctx context.Context, userID string) (entities.Wallet, error) {
	return w.store.GetWallet(ctx, userID)
}

func (w *Wallet) UpdateWalletBalance(ctx context.Context, userID string, amount float64) error {
	wallet, err := w.store.GetWallet(ctx, userID)
	if err != nil {
		return err
	}

	newBalance := wallet.Balance + amount

	// Implement checks: negative balance, etc.

	return w.store.UpdateWalletBalance(ctx, userID, newBalance)
}

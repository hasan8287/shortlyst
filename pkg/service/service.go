package service

import (
	"context"
	"shortlyst/pkg/model"
)

// ItemService controller for item pkg
type ItemService interface {
	Find(ctx context.Context) ([]model.Items, error)
	Upsert(ctx context.Context, content string) ([]model.Items, error)
	Update(ctx context.Context, data model.Items) (model.Items, error)
}

type SaldoService interface {
	Find(ctx context.Context) ([]model.Saldo, error)
	Get(ctx context.Context, val int) (model.Saldo, error)
	Upsert(ctx context.Context, content string) ([]model.Saldo, error)
	CheckCoin(ctx context.Context, coin int, price int64) ([]model.Saldo, int, error)
	Update(ctx context.Context, data model.Saldo) (model.Saldo, error)
}

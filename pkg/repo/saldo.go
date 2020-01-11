package repo

import (
	"context"
	"shortlyst/pkg/model"
)

// SaldoRepo : initial saldo mongorepo
type SaldoRepo interface {
	Get(ctx context.Context, value int) (model.Saldo, error)
	Find(ctx context.Context, params map[string]interface{}, page, size int) ([]model.Saldo, int, error)
	UpSert(ctx context.Context, id string, data model.Saldo) (model.Saldo, error)
}

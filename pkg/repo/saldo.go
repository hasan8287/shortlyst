package repo

import (
	"context"
	"shortlyst/pkg/model"
)

// SaldoRepo : initial saldo mongorepo
type SaldoRepo interface {
	Get(ctx context.Context, value int) (model.Saldo, error)
	Find(ctx context.Context, params map[string]interface{}) ([]model.Saldo, int, error)
	Update(ctx context.Context, data model.Saldo) (model.Saldo, error)
	Add(ctx context.Context, data model.Saldo) (model.Saldo, error)
}

package repo

import (
	"context"
	"shortlyst/pkg/model"
)

// ItemRepo : initial item mongorepo
type ItemRepo interface {
	Get(ctx context.Context, id string) (model.Items, error)
	Find(ctx context.Context, params map[string]interface{}, page, size int) ([]model.Items, int, error)
	Add(ctx context.Context, data model.Items) (model.Items, error)
	Update(ctx context.Context, id string, data model.Items) (model.Items, error)
}

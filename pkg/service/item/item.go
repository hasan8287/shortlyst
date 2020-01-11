package item

import (
	"shortlyst/pkg/repo"
	"shortlyst/pkg/service"
)

type itemServiceImpl struct {
	repo repo.ItemRepo
}

// NewItemService : create new instans item service
func NewItemService(repo repo.ItemRepo) service.ItemService {
	return &itemServiceImpl{
		repo: repo,
	}
}

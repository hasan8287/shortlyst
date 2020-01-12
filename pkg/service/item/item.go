package item

import (
	"context"
	"errors"
	"shortlyst/pkg/model"
	"shortlyst/pkg/repo"
	"shortlyst/pkg/service"
	"shortlyst/pkg/service/saldo"
	"strconv"
	"strings"
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

func (svc *itemServiceImpl) Update(ctx context.Context, data model.Items) (model.Items, error) {
	res, err := svc.repo.Update(ctx, data)

	return res, err
}

// Upsert : update or insert coin from file
func (svc *itemServiceImpl) Upsert(ctx context.Context, content string) ([]model.Items, error) {
	data := strings.Split(content, "\n")

	dataItemModel := []model.Items{}
	for _, entrys := range data {
		entry, err := dataToModel(strings.Split(entrys, "|"))
		if err == nil {
			_, err := svc.repo.Add(ctx, entry)

			if err == nil && entry.Price%int64(saldo.RuleCoin[0]) == 0 {
				dataItemModel = append(dataItemModel, entry)
			}
		}
	}

	return dataItemModel, nil
}

func (svc *itemServiceImpl) Find(ctx context.Context) ([]model.Items, error) {
	data, _, err := svc.repo.Find(ctx, nil, 0, 0)

	return data, err
}

func dataToModel(data []string) (model.Items, error) {
	itemData := model.Items{}
	if len(data) < 3 {
		return itemData, errors.New("failed parse data")
	}

	price, err := strconv.Atoi(data[1])
	if err != nil {
		return model.Items{}, err
	}

	stock, err := strconv.Atoi(data[2])
	if err != nil {
		return model.Items{}, err
	}

	return model.Items{
		Name:   data[0],
		Price:  int64(price),
		Stock:  stock,
		Status: 1,
	}, nil
}

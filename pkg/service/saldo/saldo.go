package saldo

import (
	"context"
	"errors"
	"shortlyst/pkg/model"
	"shortlyst/pkg/repo"
	"shortlyst/pkg/service"
	"strconv"
	"strings"
)

var (
	// RuleCoin rule
	RuleCoin = [4]int{10, 50, 100, 500}
)

type saldoServiceImpl struct {
	repo repo.SaldoRepo
}

// NewSaldoService ds
func NewSaldoService(repo repo.SaldoRepo) service.SaldoService {
	return &saldoServiceImpl{
		repo: repo,
	}
}

func (svc *saldoServiceImpl) Find(ctx context.Context) ([]model.Saldo, error) {
	data, _, err := svc.repo.Find(ctx, nil)
	return data, err
}

func (svc *saldoServiceImpl) Get(ctx context.Context, val int) (model.Saldo, error) {
	data, err := svc.repo.Get(ctx, val)

	return data, err
}

// CheckCoin check and rest coin for user
func (svc *saldoServiceImpl) CheckCoin(ctx context.Context, coin int, price int64) ([]model.Saldo, int, error) {
	rest := coin - int(price)
	var index int
	for i, ruleCoinEntry := range RuleCoin {
		if ruleCoinEntry <= rest {
			index = i
		}
	}

	var restCoin []model.Saldo
	for indexCoin := index; indexCoin >= 0; indexCoin-- {
		if rest > 0 {
			dataValue, err := svc.repo.Get(ctx, RuleCoin[indexCoin])
			if err == nil {
				restCoinEntry := model.Saldo{Value: dataValue.Value, Count: 0}
				for indexDataValue := 1; indexDataValue <= dataValue.Count; indexDataValue++ {
					if rest != 0 && rest >= dataValue.Value {
						restCoinEntry.Count++
						rest = rest - dataValue.Value

					}
				}

				restCoin = append(restCoin, restCoinEntry)
			}

		}

	}

	if rest != 0 {
		return []model.Saldo{}, 0, errors.New("invalid coin")
	}

	return restCoin, coin - int(price), nil

}

// Upsert : update or insert coin from file
func (svc *saldoServiceImpl) Upsert(ctx context.Context, content string) ([]model.Saldo, error) {
	data := strings.Split(content, "\n")

	dataSaldoModel := []model.Saldo{}
	for _, entrys := range data {
		entry, err := dataToModel(strings.Split(entrys, "|"))

		_, err = svc.repo.Add(ctx, entry)

		if err != nil {
			return []model.Saldo{}, err
		}

		if err == nil && entry.Value%RuleCoin[0] == 0 {
			dataSaldoModel = append(dataSaldoModel, entry)
		}

	}

	return dataSaldoModel, nil
}

// Update : update or insert coin from file
func (svc *saldoServiceImpl) Update(ctx context.Context, data model.Saldo) (model.Saldo, error) {
	_, err := svc.repo.Update(ctx, data)
	return data, err
}

func dataToModel(data []string) (model.Saldo, error) {
	saldoData := model.Saldo{}
	if len(data) < 2 {
		return saldoData, errors.New("failed parse data")
	}

	value, err := strconv.Atoi(data[0])
	if err != nil {
		return model.Saldo{}, err
	}

	count, err := strconv.Atoi(data[1])
	if err != nil {
		return model.Saldo{}, err
	}

	return model.Saldo{
		Value: value,
		Count: count,
	}, nil
}

package mongorepo

import (
	"context"
	"shortlyst/pkg/model"
	"shortlyst/pkg/repo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"

	"go.mongodb.org/mongo-driver/mongo"
)

const (
	sadloCollection = "saldo"
)

type repoSaldoImpl struct {
	db    *mongo.Database
	saldo *mongo.Collection
}

//NewSaldoRepo new instance of SaldoRepo using mongodb
func NewSaldoRepo(db *mongo.Database) repo.SaldoRepo {
	return repoSaldoImpl{db: db, item: db.Collection(sadloCollection)}
}

func (repo *repoSaldoImpl) Get(ctx context.Context, value int) (model.Items, error) {

	item := model.Items{}
	err = repo.item.FindOne(ctx, bson.M{"value": value}).Decode(&item)

	return item, err
}

// UpSert : update or insert saldo by value
func (repo *repoSaldoImpl) UpSert(ctx context.Context, data model.Saldo) (model.Saldo, error) {
	update := bson.D{
		{
			Key:   "$set",
			Value: data,
		},
	}

	upsert := options.Update().SetUpsert(true)
	_, err := repo.saldo.UpdateOne(ctx,
		bson.M{
			"value": data.Value,
		},
		update,
		upsert)

	if err != nil {
		return data, err
	}
	return data, nil

}

// Find data saldo
func (repo *repoSaldoImpl) Find(ctx context.Context, params map[string]interface{}, page, size int) ([]model.Saldo, int, error) {
	skip := int64((page - 1) * size)
	limit := int64(size)

	total, err := repo.saldo.CountDocuments(ctx, params)

	if err != nil {
		return nil, 0, err
	}

	opts := options.FindOptions{Sort: bson.M{
		"create_at": -1,
	}}

	if skip > 0 && limit > 0 {
		opts.Skip = &skip
		opts.Limit = &limit
	}
	res, err := repo.saldo.Find(
		ctx,
		params,
		&opts)

	if err != nil {
		return nil, 0, err
	}

	data := []model.Saldo{}
	defer res.Close(ctx)
	for res.Next(ctx) {
		saldo := model.Saldo{}
		err := res.Decode(&saldo)
		if err != nil {
			return nil, 0, err
		}

		data = append(data, saldo)
	}

	return data, int(total), nil
}

package mongorepo

import (
	"context"
	"shortlyst/pkg/model"
	"shortlyst/pkg/repo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	return &repoSaldoImpl{db: db, saldo: db.Collection(sadloCollection)}
}

func (repo *repoSaldoImpl) Get(ctx context.Context, value int) (model.Saldo, error) {

	item := model.Saldo{}
	err := repo.saldo.FindOne(ctx, bson.M{"value": value}).Decode(&item)

	return item, err
}

// Add create data Item
func (repo *repoSaldoImpl) Add(ctx context.Context, data model.Saldo) (model.Saldo, error) {
	res, err := repo.saldo.InsertOne(ctx, data)
	if err != nil {
		return model.Saldo{}, err
	}

	data.ObjectID = res.InsertedID.(primitive.ObjectID)

	return data, nil
}

// Update : update  saldo by value
func (repo *repoSaldoImpl) Update(ctx context.Context, data model.Saldo) (model.Saldo, error) {
	update := bson.D{
		{
			Key:   "$set",
			Value: data,
		},
	}

	_, err := repo.saldo.UpdateOne(ctx,
		bson.M{
			"value": data.Value,
		},
		update)

	if err != nil {
		return data, err
	}
	return data, nil

}

// Find data saldo
func (repo *repoSaldoImpl) Find(ctx context.Context, params map[string]interface{}) ([]model.Saldo, int, error) {

	total, err := repo.saldo.CountDocuments(ctx, params)

	if err != nil {
		return nil, 0, err
	}

	opts := options.FindOptions{Sort: bson.M{
		"value": -1,
	}}

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

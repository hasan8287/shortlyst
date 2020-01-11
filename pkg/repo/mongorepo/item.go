package mongorepo

import (
	"context"
	"shortlyst/pkg/model"
	"shortlyst/pkg/repo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	itemCollection = "item"
)

type repoItemImpl struct {
	db   *mongo.Database
	item *mongo.Collection
}

//NewItemRepo new instance of ItemRepo using mongodb
func NewItemRepo(db *mongo.Database) repo.ItemRepo {
	return repoItemImpl{db: db, item: db.Collection(itemCollection)}
}

func (repo *repoItemImpl) Get(ctx context.Context, id string) (model.Items, error) {
	ObjectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return model.Items{}, err
	}

	item := model.Items{}
	err = repo.item.FindOne(ctx, bson.M{"_id": ObjectID}).Decode(&item)

	return item, err
}

// Add create data Item
func (repo *repoItemImpl) Add(ctx context.Context, data model.Items) (model.Items, error) {
	res, err := repo.item.InsertOne(ctx, data)
	if err != nil {
		return model.Items{}, err
	}

	data.ObjectID = res.InsertedID.(primitive.ObjectID)

	return data, nil
}

// Update : update data Item
func (repo *repoItemImpl) Update(ctx context.Context, data model.Items) (model.Items, error) {
	update := bson.D{
		{
			Key:   "$set",
			Value: data,
		},
	}

	_, err := repo.item.UpdateOne(ctx,
		bson.M{
			"_id": data.ObjectID,
		},
		update)

	if err != nil {
		return data, err
	}
	return data, nil

}

// Find data item
func (repo *repoItemImpl) Find(ctx context.Context, params map[string]interface{}, page, size int) ([]model.Items, int, error) {
	skip := int64((page - 1) * size)
	limit := int64(size)

	total, err := repo.item.CountDocuments(ctx, params)

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
	res, err := repo.item.Find(
		ctx,
		params,
		&opts)

	if err != nil {
		return nil, 0, err
	}

	data := []model.Items{}
	defer res.Close(ctx)
	for res.Next(ctx) {
		item := model.Items{}
		err := res.Decode(&item)
		if err != nil {
			return nil, 0, err
		}

		data = append(data, item)
	}

	return data, int(total), nil
}

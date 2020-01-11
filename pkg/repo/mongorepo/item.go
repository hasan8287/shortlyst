package mongorepo

import "go.mongodb.org/mongo-driver/mongo"

type repoImpl struct {
	db   *mongo.Database
	item *mongo.Collection
}

package model

import (
	"context"
	"errors"
	"go-crud-mongo/database"
	"go-crud-mongo/helper"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	MongoCollectionUsers = `users`
)

type User struct {
	Id    primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name  string             `json:"name" bson:"name"`
	Email string             `json:"email" bson:"email"`
	Age   uint8              `json:"age" bson:"age"`
}

func NewUser() *User {
	return &User{}
}

var (
	Err500UserInsertFailed = errors.New(`failed to add user`)
)

func (u *User) Insert() error {
	collection := database.MongoClient.Database(MongoDatabaseCustomer).Collection(MongoCollectionUsers)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := collection.InsertOne(ctx, &u)
	if err != nil {
		helper.Log.Error(err, Err500UserInsertFailed.Error())

		return Err500UserInsertFailed
	}

	u.Id = res.InsertedID.(primitive.ObjectID)

	return nil
}

package model

import (
	"context"
	"errors"
	"go-crud-mongo/database"
	"go-crud-mongo/helper"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

const (
	MongoCollectionUsers = `users`
)

type User struct {
	Id    bson.ObjectID `json:"id" bson:"id"`
	Name  string        `json:"name" bson:"name"`
	Email string        `json:"email" bson:"email"`
	Age   uint8         `json:"age" bson:"age"`
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

	u.Id = bson.NewObjectID()

	res, err := collection.InsertOne(ctx, u)
	if err != nil {
		helper.Log.Error(err, Err500UserInsertFailed.Error())

		return Err500UserInsertFailed
	}

	u.Id = res.InsertedID.(bson.ObjectID)

	return nil
}

var (
	Err500UserGetAll = errors.New(`failed to get users`)
)

func (u *User) GetAll() ([]User, error) {
	var users = []User{}
	collection := database.MongoClient.Database(MongoDatabaseCustomer).Collection(MongoCollectionUsers)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{
		"id":    bson.M{"$ne": bson.NilObjectID},
		"name":  bson.M{"$ne": ""},
		"email": bson.M{"$ne": ""},
		"age":   bson.M{"$ne": 0},
	})
	if err != nil {
		helper.Log.Error(err, Err500UserGetAll.Error())
		return users, Err500UserGetAll
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var user User
		cursor.Decode(&user)
		users = append(users, user)
	}

	return users, nil
}

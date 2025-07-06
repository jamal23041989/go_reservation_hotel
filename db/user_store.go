package db

import (
	"context"
	"fmt"
	"github.com/jamal23041989/go_reservation_hotel/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	userColl = "users"
)

type UserStore interface {
	GetUserByID(context.Context, string) (*types.User, error)
	GetUsers(context.Context) ([]*types.User, error)
	InsertUser(context.Context, *types.User) (*types.User, error)
	UpdateUser(context.Context, string, bson.M) error
	DeleteUser(context.Context, string) error
}

type MongoUserStore struct {
	client *mongo.Client
	coll   *mongo.Collection
}

func NewMongoUserStore(client *mongo.Client) *MongoUserStore {
	return &MongoUserStore{
		client: client,
		coll:   client.Database(dbName).Collection(userColl),
	}
}

func (s *MongoUserStore) GetUserByID(ctx context.Context, id string) (*types.User, error) {
	objectID, err := convertToObjectID(id)
	if err != nil {
		return nil, err
	}

	var user types.User
	if err := s.coll.FindOne(ctx, bson.M{"_id": objectID}).Decode(&user); err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *MongoUserStore) GetUsers(ctx context.Context) ([]*types.User, error) {
	cursor, err := s.coll.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var users []*types.User
	if err := cursor.All(ctx, &users); err != nil {
		return nil, err
	}

	return users, nil
}

func (s *MongoUserStore) InsertUser(ctx context.Context, user *types.User) (*types.User, error) {
	res, err := s.coll.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}
	user.ID = res.InsertedID.(primitive.ObjectID)
	return user, nil
}

func (s *MongoUserStore) UpdateUser(ctx context.Context, id string, update bson.M) error {
	objectID, err := convertToObjectID(id)
	if err != nil {
		return err
	}

	updateOne, err := s.coll.UpdateOne(ctx, bson.M{"_id": objectID}, bson.M{"$set": update})
	if err != nil {
		return err
	}
	if updateOne.MatchedCount == 0 {
		return mongo.ErrNoDocuments
	}

	return nil
}

func (s *MongoUserStore) DeleteUser(ctx context.Context, id string) error {
	objectID, err := convertToObjectID(id)
	if err != nil {
		return err
	}

	deleteOne, err := s.coll.DeleteOne(ctx, bson.M{"_id": objectID})
	if err != nil {
		return err
	}
	if deleteOne.DeletedCount == 0 {
		return fmt.Errorf("not a single user was deleted")
	}

	return nil
}

package mongodb

import (
	"context"
	"fmt"
	"github.com/jamal23041989/go_reservation_hotel/internal/domain"
	"github.com/jamal23041989/go_reservation_hotel/internal/domain/entity"
	"github.com/jamal23041989/go_reservation_hotel/pkg"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const userColl = "users"

type MongoUserRepository struct {
	client *mongo.Client
	coll   *mongo.Collection
}

func NewMongoUserRepository(client *mongo.Client) *MongoUserRepository {
	return &MongoUserRepository{
		client: client,
		coll:   client.Database(DbName).Collection(userColl),
	}
}

func (s *MongoUserRepository) GetUserByID(ctx context.Context, id string) (*entity.User, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, pkg.ErrInvalidID()
	}

	var user entity.User
	if err := s.coll.FindOne(ctx, bson.M{"_id": objectID}).Decode(&user); err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *MongoUserRepository) GetUsers(ctx context.Context) ([]*entity.User, error) {
	cursor, err := s.coll.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var users []*entity.User
	if err := cursor.All(ctx, &users); err != nil {
		return nil, err
	}

	return users, nil
}

func (s *MongoUserRepository) GetByEmailUser(ctx context.Context, email string) (*entity.User, error) {
	var user entity.User
	if err := s.coll.FindOne(ctx, bson.M{"email": email}).Decode(&user); err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *MongoUserRepository) CreateUser(ctx context.Context, user *entity.User) (*entity.User, error) {
	resp, err := s.coll.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}
	user.ID = resp.InsertedID.(primitive.ObjectID)
	return user, nil
}

func (s *MongoUserRepository) UpdateUser(ctx context.Context, id string, update domain.Map) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return pkg.ErrInvalidID()
	}

	res, err := s.coll.UpdateOne(ctx, bson.M{"_id": objectID}, bson.M{"$set": update})
	if err != nil {
		return err
	}
	if res.MatchedCount == 0 {
		return mongo.ErrNoDocuments
	}

	return nil
}

func (s *MongoUserRepository) DeleteUser(ctx context.Context, id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return pkg.ErrInvalidID()
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

func (s *MongoUserRepository) DropUser(ctx context.Context) error {
	fmt.Println("--- dropping user collection")
	return s.coll.Drop(ctx)
}

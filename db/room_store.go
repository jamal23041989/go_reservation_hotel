package db

import (
	"context"
	"github.com/jamal23041989/go_reservation_hotel/types"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"os"
)

const (
	roomColl = "rooms"
)

type RoomStore interface {
	InsertRoom(context.Context, *types.Room) (*types.Room, error)
	GetRooms(context.Context, Map) ([]*types.Room, error)
}

type MongoRoomStore struct {
	client     *mongo.Client
	coll       *mongo.Collection
	hotelStore HotelStore
}

func NewMongoRoomStore(client *mongo.Client, hotelStore HotelStore) *MongoRoomStore {
	dbName := os.Getenv(MongoDbNameEnvName)
	return &MongoRoomStore{
		client:     client,
		coll:       client.Database(dbName).Collection(roomColl),
		hotelStore: hotelStore,
	}
}

func (s *MongoRoomStore) InsertRoom(ctx context.Context, room *types.Room) (*types.Room, error) {
	resp, err := s.coll.InsertOne(ctx, room)
	if err != nil {
		return nil, err
	}

	room.ID = resp.InsertedID.(primitive.ObjectID)

	filter := Map{"_id": room.HotelID}
	update := Map{"$push": Map{"rooms": room.ID}}

	if err := s.hotelStore.Update(ctx, filter, update); err != nil {
		return nil, err
	}

	return room, nil
}

func (s *MongoRoomStore) GetRooms(ctx context.Context, filter Map) ([]*types.Room, error) {
	cursor, err := s.coll.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var rooms []*types.Room
	if err := cursor.All(ctx, &rooms); err != nil {
		return nil, err
	}

	return rooms, nil
}

package mongodb

import (
	"context"
	"github.com/jamal23041989/go_reservation_hotel/internal/domain"
	"github.com/jamal23041989/go_reservation_hotel/internal/domain/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const roomColl = "rooms"

type MongoRoomRepository struct {
	client               *mongo.Client
	coll                 *mongo.Collection
	MongoHotelRepository MongoHotelRepository
}

func NewMongoRoomRepository(client *mongo.Client, MongoHotelRepository MongoHotelRepository) *MongoRoomRepository {
	return &MongoRoomRepository{
		client:               client,
		coll:                 client.Database(DbName).Collection(roomColl),
		MongoHotelRepository: MongoHotelRepository,
	}
}

func (s *MongoRoomRepository) CreateRoom(ctx context.Context, room *entity.Room) (*entity.Room, error) {
	resp, err := s.coll.InsertOne(ctx, room)
	if err != nil {
		return nil, err
	}

	room.ID = resp.InsertedID.(primitive.ObjectID)

	filter := domain.Map{"_id": room.HotelID}
	update := domain.Map{"$push": domain.Map{"rooms": room.ID}}

	if err := s.MongoHotelRepository.UpdateHotel(ctx, filter, update); err != nil {
		return nil, err
	}

	return room, nil
}

func (s *MongoRoomRepository) GetRooms(ctx context.Context, filter domain.Map) ([]*entity.Room, error) {
	cursor, err := s.coll.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var rooms []*entity.Room
	if err := cursor.All(ctx, &rooms); err != nil {
		return nil, err
	}

	return rooms, nil
}

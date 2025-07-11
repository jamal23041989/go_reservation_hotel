package handler

import (
	"context"
	"github.com/jamal23041989/go_reservation_hotel/internal/repository/mongodb"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"testing"
)

type testDB struct {
	client *mongo.Client
	Store  *Store
}

type Store struct {
	Booking *mongodb.MongoBookingRepository
	User    *mongodb.MongoUserRepository
	Hotel   *mongodb.MongoHotelRepository
	Room    *mongodb.MongoRoomRepository
}

func (tdb *testDB) teardown(t *testing.T) {
	if err := godotenv.Load("../.env"); err != nil {
		log.Fatal("Ошибка при загрузке .env файла")
	}

	if err := tdb.client.Database("hotel-reservation-test").Drop(context.Background()); err != nil {
		t.Fatal(err)
	}
}

func setup(t *testing.T) *testDB {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}

	hotelStore := mongodb.NewMongoHotelRepository(client)
	store := &Store{
		Booking: mongodb.NewMongoBookingRepository(client),
		User:    mongodb.NewMongoUserRepository(client),
		Room:    mongodb.NewMongoRoomRepository(client, *hotelStore),
		Hotel:   hotelStore,
	}

	return &testDB{
		client: client,
		Store:  store,
	}
}

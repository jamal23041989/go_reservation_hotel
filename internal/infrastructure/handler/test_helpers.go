package handler

import (
	"context"
	mongodb2 "github.com/jamal23041989/go_reservation_hotel/internal/infrastructure/db/mongodb"
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
	Booking *mongodb2.MongoBookingRepository
	User    *mongodb2.MongoUserRepository
	Hotel   *mongodb2.MongoHotelRepository
	Room    *mongodb2.MongoRoomRepository
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

	hotelStore := mongodb2.NewMongoHotelRepository(client)
	store := &Store{
		Booking: mongodb2.NewMongoBookingRepository(client),
		User:    mongodb2.NewMongoUserRepository(client),
		Room:    mongodb2.NewMongoRoomRepository(client, *hotelStore),
		Hotel:   hotelStore,
	}

	return &testDB{
		client: client,
		Store:  store,
	}
}

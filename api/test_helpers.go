package api

import (
	"context"
	"github.com/jamal23041989/go_reservation_hotel/db"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"testing"
)

type testDB struct {
	client *mongo.Client
	*db.Store
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

	hotelStore := db.NewMongoHotelStore(client)
	store := &db.Store{
		User:    db.NewMongoUserStore(client),
		Hotel:   hotelStore,
		Room:    db.NewMongoRoomStore(client, hotelStore),
		Booking: db.NewMongoBookingStore(client),
	}

	return &testDB{
		client: client,
		Store:  store,
	}
}

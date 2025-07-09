package api

import (
	"context"
	"github.com/jamal23041989/go_reservation_hotel/db"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"testing"
)

const (
	TestDbName = "hotel-reservation-test"
	TestDbUri  = "mongodb://localhost:27017"
)

type testDB struct {
	client *mongo.Client
	*db.Store
}

func (tdb *testDB) teardown(t *testing.T) {
	if err := tdb.client.Database(TestDbName).Drop(context.Background()); err != nil {
		t.Fatal(err)
	}
}

func setup(t *testing.T) *testDB {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(TestDbUri))
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

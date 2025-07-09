package main

import (
	"context"
	"fmt"
	"github.com/jamal23041989/go_reservation_hotel/api"
	"github.com/jamal23041989/go_reservation_hotel/db"
	"github.com/jamal23041989/go_reservation_hotel/db/fixtures"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

func main() {
	ctx := context.Background()

	var err error
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(db.UriDb))
	if err != nil {
		log.Fatal(err)
	}

	if err := client.Database(db.NameDb).Drop(ctx); err != nil {
		log.Fatal(err)
	}

	hotelStore := db.NewMongoHotelStore(client)
	store := &db.Store{
		User:    db.NewMongoUserStore(client),
		Room:    db.NewMongoRoomStore(client, hotelStore),
		Booking: db.NewMongoBookingStore(client),
		Hotel:   hotelStore,
	}

	user := fixtures.AddUser(store, "foo", "foo", false)
	fmt.Println("User -> ", api.CreateTokenFromUser(user))
	admin := fixtures.AddUser(store, "admin", "admin", true)
	fmt.Println("Admin -> ", api.CreateTokenFromUser(admin))
	hotel := fixtures.AddHotel(store, "some hotel", "bermude", 5, nil)
	room := fixtures.AddRoom(store, "medium", false, 129.29, hotel.ID)
	booking := fixtures.AddBooking(store, user.ID, room.ID, time.Now(), time.Now().AddDate(0, 0, 2))
	fmt.Println(booking)
}

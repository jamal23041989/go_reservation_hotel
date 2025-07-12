package fixtures

import (
	"context"
	"fmt"
	types2 "github.com/jamal23041989/go_reservation_hotel/aold/types"
	"github.com/jamal23041989/go_reservation_hotel/internal/infrastructure/db/mongodb"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"time"
)

func AddUser(store *mongodb.Store, fn, ln string, admin bool) *types2.User {
	user, err := types2.NewUserFromParams(types2.CreateUserParams{
		Email:     fmt.Sprintf("%s@%s.com", fn, ln),
		FirstName: fn,
		LastName:  ln,
		Password:  fmt.Sprintf("%s_%s", fn, ln),
	})
	if err != nil {
		log.Fatal(err)
	}

	user.IsAdmin = admin

	insertedUser, err := store.User.InsertUser(context.Background(), user)
	if err != nil {
		log.Fatal(err)
	}
	return insertedUser
}

func AddHotel(store *mongodb.Store, name, loc string, rating int, rooms []primitive.ObjectID) *types2.Hotel {
	var roomIDString = rooms
	if rooms == nil {
		roomIDString = []primitive.ObjectID{}
	}
	hotel := types2.Hotel{
		Name:     name,
		Location: loc,
		Rating:   rating,
		Rooms:    roomIDString,
	}

	insertedHotel, err := store.Hotel.Insert(context.Background(), &hotel)
	if err != nil {
		log.Fatal(err)
	}
	return insertedHotel
}

func AddRoom(store *mongodb.Store, size string, ss bool, price float64, hotelID primitive.ObjectID) *types2.Room {
	room := types2.Room{
		Size:    size,
		Price:   price,
		Seaside: ss,
		HotelID: hotelID,
	}

	insertedRoom, err := store.Room.InsertRoom(context.Background(), &room)
	if err != nil {
		log.Fatal(err)
	}
	return insertedRoom
}

func AddBooking(store *mongodb.Store, uid, rid primitive.ObjectID, from, till time.Time) *types2.Booking {
	booking := types2.Booking{
		UserID:   uid,
		RoomID:   rid,
		FromDate: from,
		TillDate: till,
	}

	insertBooking, err := store.Booking.InsertBooking(context.Background(), &booking)
	if err != nil {
		log.Fatal(err)
	}

	return insertBooking
}

package fixtures

import (
	"context"
	"fmt"
	"github.com/jamal23041989/go_reservation_hotel/db"
	"github.com/jamal23041989/go_reservation_hotel/types"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"time"
)

func AddUser(store *db.Store, fn, ln string, admin bool) *types.User {
	user, err := types.NewUserFromParams(types.CreateUserParams{
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

func AddHotel(store *db.Store, name, loc string, rating int, rooms []primitive.ObjectID) *types.Hotel {
	var roomIDString = rooms
	if rooms == nil {
		roomIDString = []primitive.ObjectID{}
	}
	hotel := types.Hotel{
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

func AddRoom(store *db.Store, size string, ss bool, price float64, hotelID primitive.ObjectID) *types.Room {
	room := types.Room{
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

func AddBooking(store *db.Store, uid, rid primitive.ObjectID, from, till time.Time) *types.Booking {
	booking := types.Booking{
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

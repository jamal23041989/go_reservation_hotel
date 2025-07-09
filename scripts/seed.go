package main

import (
	"context"
	"github.com/jamal23041989/go_reservation_hotel/db"
	"github.com/jamal23041989/go_reservation_hotel/types"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

var (
	client       *mongo.Client
	roomStore    db.RoomStore
	hotelStore   db.HotelStore
	userStore    db.UserStore
	bookingStore db.BookingStore
	ctx          = context.Background()
)

func seedUser(isAdmin bool, firstName, lastName, email, password string) *types.User {
	user, err := types.NewUserFromParams(types.CreateUserParams{
		Email:     email,
		FirstName: firstName,
		LastName:  lastName,
		Password:  password,
	})
	if err != nil {
		log.Fatal(err)
	}

	user.IsAdmin = isAdmin

	insertedUser, err := userStore.InsertUser(ctx, user)
	if err != nil {
		log.Fatal(err)
	}

	return insertedUser
}

func seedRoom(size string, side bool, price float64, hotelID primitive.ObjectID) *types.Room {
	room := &types.Room{
		Size:    size,
		Price:   price,
		Seaside: side,
		HotelID: hotelID,
	}

	insertedRoom, err := roomStore.InsertRoom(ctx, room)
	if err != nil {
		return nil
	}

	return insertedRoom
}

func seedBooking(userID, roomID primitive.ObjectID, from, till time.Time) *types.Booking {
	booking := &types.Booking{
		UserID:   userID,
		RoomID:   roomID,
		FromDate: from,
		TillDate: till,
	}

	insertBooking, err := bookingStore.InsertBooking(ctx, booking)
	if err != nil {
		log.Fatal(err)
	}

	return insertBooking
}

func seedHotel(name, location string, rating int) *types.Hotel {
	hotel := types.Hotel{
		Name:     name,
		Location: location,
		Rooms:    []primitive.ObjectID{},
		Rating:   rating,
	}

	insertHotel, err := hotelStore.Insert(ctx, &hotel)
	if err != nil {
		log.Fatal(err)
	}

	return insertHotel
}

func main() {
	james := seedUser(false, "james", "foo", "james@foo.com", "supersecurepassword")
	seedUser(true, "admin", "admin", "admin@gmail.com", "admin")

	seedHotel("Bellucia", "France", 4)
	seedHotel("The cozy", "Spain", 5)
	hotel := seedHotel("Porto", "Portugal", 3)

	seedRoom("small", false, 89.99, hotel.ID)
	seedRoom("medium", true, 129.99, hotel.ID)
	room := seedRoom("large", true, 199.99, hotel.ID)

	seedBooking(james.ID, room.ID, time.Now(), time.Now().AddDate(0, 0, 2))
}

func init() {
	var err error
	client, err = mongo.Connect(ctx, options.Client().ApplyURI(db.UriDb))
	if err != nil {
		log.Fatal(err)
	}

	if err := client.Database(db.NameDb).Drop(ctx); err != nil {
		log.Fatal(err)
	}

	hotelStore = db.NewMongoHotelStore(client)
	roomStore = db.NewMongoRoomStore(client, hotelStore)
	userStore = db.NewMongoUserStore(client)
	bookingStore = db.NewMongoBookingStore(client)
}

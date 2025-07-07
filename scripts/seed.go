package main

import (
	"context"
	"fmt"
	"github.com/jamal23041989/go_reservation_hotel/db"
	"github.com/jamal23041989/go_reservation_hotel/types"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

var (
	client     *mongo.Client
	roomStore  db.RoomStore
	hotelStore db.HotelStore
	userStore  db.UserStore
	ctx        = context.Background()
)

func seedUser(firstName, lastName, email string) {
	user, err := types.NewUserFromParams(types.CreateUserParams{
		Email:     email,
		FirstName: firstName,
		LastName:  lastName,
		Password:  "supersecurepassword",
	})

	if err != nil {
		log.Fatal(err)
	}

	if _, err := userStore.InsertUser(context.Background(), user); err != nil {
		log.Fatal(err)
	}
}

func seedHotel(name, location string, rating int) {
	hotel := types.Hotel{
		Name:     name,
		Location: location,
		Rooms:    []primitive.ObjectID{},
		Rating:   rating,
	}

	rooms := []types.Room{
		{
			Size:      "small",
			BasePrice: 99.9,
		},
		{
			Size:      "king",
			BasePrice: 299.9,
		},
		{
			Size:      "normal",
			BasePrice: 159.9,
		},
	}

	insertHotel, err := hotelStore.Insert(ctx, &hotel)
	if err != nil {
		log.Fatal(err)
	}

	for _, room := range rooms {
		room.HotelID = insertHotel.ID

		insertRoom, err := roomStore.InsertRoom(ctx, &room)
		if err != nil {
			log.Fatal(err)
		}

		hotel.Rooms = append(hotel.Rooms, insertRoom.ID)
		fmt.Println(insertRoom)
	}

	fmt.Println(insertHotel)
}

func main() {
	seedHotel("Bellucia", "France", 4)
	seedHotel("The cozy", "Spain", 5)
	seedHotel("Porto", "Portugal", 3)
	seedUser("james", "foo", "james@foo.com")
}

func init() {
	var err error
	client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(db.UriDb))
	if err != nil {
		log.Fatal(err)
	}

	if err := client.Database(db.NameDb).Drop(ctx); err != nil {
		log.Fatal(err)
	}

	hotelStore = db.NewMongoHotelStore(client)
	roomStore = db.NewMongoRoomStore(client, hotelStore)
	userStore = db.NewMongoUserStore(client)
}

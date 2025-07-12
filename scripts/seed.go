package main

import (
	"context"
	"fmt"
	"github.com/jamal23041989/go_reservation_hotel/db"
	mongo2 "github.com/jamal23041989/go_reservation_hotel/internal/db/mongodb"
	"github.com/jamal23041989/go_reservation_hotel/internal/infrastructure/db/mongodb"
	"github.com/jamal23041989/go_reservation_hotel/internal/infrastructure/db/mongodb/fixtures"
	"github.com/jamal23041989/go_reservation_hotel/internal/infrastructure/handler"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/exp/rand"
	"log"
	"os"
	"time"
)

func init() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Ошибка при загрузке .env файла")
	}
}

func main() {
	dbUri := os.Getenv("MONGO_URI")
	dbName := os.Getenv("MONGO_DB_NAME")
	ctx := context.Background()

	var err error
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dbUri))
	if err != nil {
		log.Fatal(err)
	}

	if err := client.Database(dbName).Drop(ctx); err != nil {
		log.Fatal(err)
	}

	hotelStore := db.NewMongoHotelStore(client)
	store := &mongodb.Store{
		User:    mongo2.NewMongoUserStore(client),
		Room:    mongo2.NewMongoRoomStore(client, hotelStore),
		Booking: mongo2.NewMongoBookingStore(client),
		Hotel:   hotelStore,
	}

	user := fixtures.AddUser(store, "foo", "foo", false)
	fmt.Println("User -> ", handler.CreateTokenFromUser(user))
	admin := fixtures.AddUser(store, "admin", "admin", true)
	fmt.Println("Admin -> ", handler.CreateTokenFromUser(admin))
	hotel := fixtures.AddHotel(store, "some hotel", "bermude", 5, nil)
	room := fixtures.AddRoom(store, "medium", false, 129.29, hotel.ID)
	booking := fixtures.AddBooking(store, user.ID, room.ID, time.Now(), time.Now().AddDate(0, 0, 2))
	fmt.Println(booking)

	for i := 1; i <= 100; i++ {
		name := fmt.Sprintf("fake hotel number %d", i)
		location := fmt.Sprintf("location %d", i)
		fixtures.AddHotel(store, name, location, rand.Intn(5)+1, nil)
	}
}

package main

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/jamal23041989/go_reservation_hotel/api"
	"github.com/jamal23041989/go_reservation_hotel/db"
	"github.com/jamal23041989/go_reservation_hotel/middleware"
	"github.com/jamal23041989/go_reservation_hotel/pkg"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
)

func init() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Ошибка при загрузке .env файла")
	}
}

var config = fiber.Config{
	ErrorHandler: pkg.ErrorHandler,
}

func main() {
	dbUri := os.Getenv("MONGO_URI")
	httpListenAddress := os.Getenv("MONGO_DB_NAME")

	// mongodb client
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(dbUri))
	if err != nil {
		log.Fatal(err)
	}

	// store
	var (
		userStore    = db.NewMongoUserStore(client)
		hotelStore   = db.NewMongoHotelStore(client)
		roomStore    = db.NewMongoRoomStore(client, hotelStore)
		bookingStore = db.NewMongoBookingStore(client)
		store        = &db.Store{
			User:    userStore,
			Hotel:   hotelStore,
			Room:    roomStore,
			Booking: bookingStore,
		}
	)

	// handlers init
	var (
		userHandler    = api.NewUserHandler(userStore)
		hotelHandler   = api.NewHotelHandler(store)
		authHandler    = api.NewAuthHandler(userStore)
		roomHandler    = api.NewRoomHandler(store)
		bookingHandler = api.NewBookingHandler(store)
	)

	// fiber config and group handler
	app := fiber.New(config)
	apiV1 := app.Group("/api/v1", middleware.JwtAuthentication(userStore))
	auth := app.Group("/api")
	admin := apiV1.Group("/admin", middleware.AdminAuth)

	// auth
	auth.Post("/auth", authHandler.HandleAuthenticate)

	// user handlers
	apiV1.Get("/user", userHandler.HandleGetUsers)
	apiV1.Get("/user/:id", userHandler.HandleGetUser)
	apiV1.Post("/user", userHandler.HandleInsertUser)
	apiV1.Put("/user/:id", userHandler.HandleUpdateUser)
	apiV1.Delete("/user/:id", userHandler.HandleDeleteUser)

	// hotel handlers
	apiV1.Get("/hotel", hotelHandler.HandleGetHotels)
	apiV1.Get("/hotel/:id", hotelHandler.HandleGetHotel)
	apiV1.Get("/hotel/:id/rooms", hotelHandler.HandleGetHotelByIDRooms)

	// room handlers
	apiV1.Get("/room", roomHandler.HandleGetRooms)
	apiV1.Post("/room/:id/book", roomHandler.HandleBookRoom)

	// booking handlers
	apiV1.Get("/booking/:id", bookingHandler.HandleGetBooking)
	apiV1.Get("/booking/:id/cancel", bookingHandler.HandleCancelBooking)

	// admin handlers
	admin.Get("/booking", bookingHandler.HandleGetBookings)

	// http server
	app.Listen(httpListenAddress)
}

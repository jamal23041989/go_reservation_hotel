package main

import (
	"context"
	"flag"
	"github.com/gofiber/fiber/v2"
	"github.com/jamal23041989/go_reservation_hotel/api"
	"github.com/jamal23041989/go_reservation_hotel/db"
	"github.com/jamal23041989/go_reservation_hotel/middleware"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

var config = fiber.Config{
	ErrorHandler: func(ctx *fiber.Ctx, err error) error {
		return ctx.JSON(map[string]string{"error": err.Error()})
	},
}

func main() {
	// listenAddr
	listenAddr := flag.String("listenAddr", ":5000", "The listen address of the API server")
	flag.Parse()

	// mongodb client
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(db.UriDb))
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
		userHandler  = api.NewUserHandler(userStore)
		hotelHandler = api.NewHotelHandler(store)
		authHandler  = api.NewAuthHandler(userStore)
		roomHandler  = api.NewRoomHandler(store)
	)

	// fiber config and group handler
	app := fiber.New(config)
	apiV1 := app.Group("/api/v1", middleware.JwtAuthentication(userStore))
	auth := app.Group("/api")

	// auth
	auth.Post("/auth", authHandler.HandleAuthenticate)

	// user handler
	apiV1.Get("/user", userHandler.HandleGetUsers)
	apiV1.Get("/user/:id", userHandler.HandleGetUser)
	apiV1.Post("/user", userHandler.HandleInsertUser)
	apiV1.Put("/user/:id", userHandler.HandleUpdateUser)
	apiV1.Delete("/user/:id", userHandler.HandleDeleteUser)

	// hotel handler
	apiV1.Get("/hotel", hotelHandler.HandleGetHotels)
	apiV1.Get("/hotel/:id", hotelHandler.HandleGetHotelByID)
	apiV1.Get("/hotel/:id/rooms", hotelHandler.HandleGetHotelByIDRooms)

	// booking handler
	apiV1.Get("/room", roomHandler.HandleGetRooms)
	apiV1.Post("/room/:id/book", roomHandler.HandleBookRoom)

	// http server
	app.Listen(*listenAddr)
}

package main

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/jamal23041989/go_reservation_hotel/internal/handler"
	"github.com/jamal23041989/go_reservation_hotel/internal/middleware"
	"github.com/jamal23041989/go_reservation_hotel/internal/repository/mongodb"
	"github.com/jamal23041989/go_reservation_hotel/internal/usecase"
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

	var (
		// repository init
		userRepository    = mongodb.NewMongoUserRepository(client)
		bookingRepository = mongodb.NewMongoBookingRepository(client)
		hotelRepository   = mongodb.NewMongoHotelRepository(client)
		roomRepository    = mongodb.NewMongoRoomRepository(client, *hotelRepository)

		// usecase init
		userCases    = usecase.NewUserUsecase(userRepository)
		bookingCases = usecase.NewBookingUsecase(bookingRepository)
		hotelCases   = usecase.NewHotelUsecase(hotelRepository)
		roomCases    = usecase.NewRoomUsecase(roomRepository)

		// handlers init
		userHandler    = handler.NewUserHandler(userCases)
		authHandler    = handler.NewAuthHandler(userCases)
		bookingHandler = handler.NewBookingHandler(bookingCases)
		roomHandler    = handler.NewRoomHandler(roomCases, bookingCases)
		hotelHandler   = handler.NewHotelHandler(hotelCases, roomCases)
	)

	// fiber config and group handler
	app := fiber.New(config)
	apiV1 := app.Group("/api/v1", middleware.AuthMiddleware())
	auth := app.Group("/api")
	//admin := apiV1.Group("/admin",)

	// auth
	auth.Post("/auth", authHandler.HandleAuthenticate)

	// user handlers
	apiV1.Get("/user", userHandler.HandleGetUsers)
	apiV1.Get("/user/:id", userHandler.HandleGetUserByID)
	apiV1.Post("/user", userHandler.HandleCreateUser)
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

package main

import (
	"context"
	"flag"
	"github.com/gofiber/fiber/v2"
	"github.com/jamal23041989/go_reservation_hotel/api"
	"github.com/jamal23041989/go_reservation_hotel/db"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

const (
	dbName = "hotel-reservation"
	dbUri  = "mongodb://localhost:27017"
)

var config = fiber.Config{
	ErrorHandler: func(ctx *fiber.Ctx, err error) error {
		return ctx.JSON(map[string]string{"error": err.Error()})
	},
}

func main() {

	listenAddr := flag.String("listenAddr", ":5000", "The listen address of the API server")
	flag.Parse()

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(dbUri))
	if err != nil {
		log.Fatal(err)
	}

	// handlers init
	userHandler := api.NewUserHandler(db.NewMongoUserStore(client, dbName))

	app := fiber.New(config)
	apiV1 := app.Group("/api/v1")

	apiV1.Get("/user", userHandler.HandleGetUsers)
	apiV1.Get("/user/:id", userHandler.HandleGetUser)
	apiV1.Post("/user", userHandler.HandleInsertUser)
	apiV1.Put("/user/:id", userHandler.HandleUpdateUser)
	apiV1.Delete("/user/:id", userHandler.HandleDeleteUser)

	app.Listen(*listenAddr)
}

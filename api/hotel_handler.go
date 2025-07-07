package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jamal23041989/go_reservation_hotel/db"
	"go.mongodb.org/mongo-driver/bson"
)

type HotelHandler struct {
	store *db.Store
}

func NewHotelHandler(store *db.Store) *HotelHandler {
	return &HotelHandler{
		store: store,
	}
}

func (h *HotelHandler) HandleGetHotels(c *fiber.Ctx) error {
	hotels, err := h.store.Hotel.GetHotels(c.Context(), bson.M{})
	if err != nil {
		return err
	}
	return c.JSON(hotels)
}

func (h *HotelHandler) HandleGetHotelByIDRooms(c *fiber.Ctx) error {
	id := c.Params("id")
	objectID, err := db.ConvertToObjectID(id)
	if err != nil {
		return err
	}

	filter := bson.M{"hotel_id": objectID}
	rooms, err := h.store.Room.GetRooms(c.Context(), filter)
	if err != nil {
		return err
	}

	return c.JSON(rooms)
}

func (h *HotelHandler) HandleGetHotelByID(c *fiber.Ctx) error {
	id := c.Params("id")

	hotel, err := h.store.Hotel.GetHotelByID(c.Context(), id)
	if err != nil {
		return err
	}

	return c.JSON(hotel)
}

package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jamal23041989/go_reservation_hotel/db"
	"github.com/jamal23041989/go_reservation_hotel/pkg"
)

type HotelHandler struct {
	store *db.Store
}

func NewHotelHandler(store *db.Store) *HotelHandler {
	return &HotelHandler{
		store: store,
	}
}

type ResourceResp struct {
	Results int `json:"results"`
	Data    any `json:"data"`
	Page    int `json:"page"`
}

type HotelFilter struct {
	db.Pagination
	Rating int
}

func (h *HotelHandler) HandleGetHotels(c *fiber.Ctx) error {
	var hotelQueryParams HotelFilter
	if err := c.QueryParser(&hotelQueryParams); err != nil {
		return pkg.ErrBadRequest()
	}

	if hotelQueryParams.Limit < 0 {
		return fiber.NewError(fiber.StatusBadRequest, "Limit cannot be negative")
	}
	if hotelQueryParams.Page < 0 {
		return fiber.NewError(fiber.StatusBadRequest, "Page cannot be negative")
	}
	if hotelQueryParams.Limit > 100 {
		hotelQueryParams.Limit = 100
	}

	if hotelQueryParams.Rating < 0 {
		return fiber.NewError(fiber.StatusBadRequest, "Rating cannot be negative")
	}

	if hotelQueryParams.Rating <= 0 {
		hotelQueryParams.Rating = 1
	}
	if hotelQueryParams.Rating > 5 {
		hotelQueryParams.Rating = 5
	}

	filter := db.Map{
		"rating": hotelQueryParams.Rating,
	}

	hotels, err := h.store.Hotel.GetHotels(c.Context(), filter, &hotelQueryParams.Pagination)
	if err != nil {
		return err
	}

	resp := ResourceResp{
		Data:    hotels,
		Results: len(hotels),
		Page:    int(hotelQueryParams.Page),
	}

	return c.JSON(resp)
}

func (h *HotelHandler) HandleGetHotelByIDRooms(c *fiber.Ctx) error {
	id := c.Params("id")

	rooms, err := h.store.Room.GetRooms(c.Context(), db.Map{"hotel_id": id})
	if err != nil {
		return err
	}

	return c.JSON(rooms)
}

func (h *HotelHandler) HandleGetHotel(c *fiber.Ctx) error {
	id := c.Params("id")

	hotel, err := h.store.Hotel.GetHotelByID(c.Context(), id)
	if err != nil {
		return err
	}

	return c.JSON(hotel)
}

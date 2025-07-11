package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jamal23041989/go_reservation_hotel/internal/domain"
	"github.com/jamal23041989/go_reservation_hotel/internal/usecase"
	"github.com/jamal23041989/go_reservation_hotel/pkg"
)

type HotelHandler struct {
	hotelUsecase usecase.HotelUsecase
	roomUsecase  usecase.RoomUsecase
}

func NewHotelHandler(hotelUsecase *usecase.HotelUsecase, roomUsecase *usecase.RoomUsecase) *HotelHandler {
	return &HotelHandler{
		hotelUsecase: *hotelUsecase,
		roomUsecase:  *roomUsecase,
	}
}

func (h *HotelHandler) HandleGetHotels(c *fiber.Ctx) error {
	var hotelQueryParams domain.HotelFilter
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

	filter := domain.Map{
		"rating": hotelQueryParams.Rating,
	}

	hotels, err := h.hotelUsecase.GetAllHotels(c.Context(), filter, &hotelQueryParams.Pagination)
	if err != nil {
		return err
	}

	resp := domain.ResourceResp{
		Data:    hotels,
		Results: len(hotels),
		Page:    int(hotelQueryParams.Page),
	}

	return c.JSON(resp)
}

func (h *HotelHandler) HandleGetHotelByIDRooms(c *fiber.Ctx) error {
	id := c.Params("id")

	rooms, err := h.roomUsecase.GetRooms(c.Context(), domain.Map{"hotel_id": id})
	if err != nil {
		return err
	}

	return c.JSON(rooms)
}

func (h *HotelHandler) HandleGetHotel(c *fiber.Ctx) error {
	id := c.Params("id")

	hotel, err := h.hotelUsecase.GetByIDHotel(c.Context(), id)
	if err != nil {
		return err
	}

	return c.JSON(hotel)
}

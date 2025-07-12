package handler

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/jamal23041989/go_reservation_hotel/internal/domain"
	"github.com/jamal23041989/go_reservation_hotel/internal/domain/models"
	"github.com/jamal23041989/go_reservation_hotel/internal/domain/usecases"
	"github.com/jamal23041989/go_reservation_hotel/pkg"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"time"
)

type RoomHandler struct {
	roomUsecase    usecases.RoomUsecase
	bookingUsecase usecases.BookingUsecase
}

func NewRoomHandler(roomUsecase usecases.RoomUsecase, bookingUsecase usecases.BookingUsecase) *RoomHandler {
	return &RoomHandler{
		roomUsecase:    roomUsecase,
		bookingUsecase: bookingUsecase,
	}
}

func (h *RoomHandler) HandleBookRoom(c *fiber.Ctx) error {
	var params domain.BookRoomParams
	if err := c.BodyParser(&params); err != nil {
		return err
	}

	now := time.Now()
	if now.After(params.FromDate) || now.After(params.TillDate) {
		return fmt.Errorf("cannot book a room in the past")
	}

	roomID, err := primitive.ObjectIDFromHex(c.Params("id"))
	if err != nil {
		return pkg.ErrInvalidID()
	}

	user, ok := c.Context().Value("user").(models.User)
	if !ok {
		return c.Status(http.StatusInternalServerError).JSON(domain.GenericResp{
			Type: "error",
			Msg:  "internal server error",
		})
	}

	ok, err = h.isRoomAvailableForBooking(c.Context(), roomID, params)
	if err != nil {
		return err
	}
	if !ok {
		return c.Status(http.StatusBadRequest).JSON(domain.GenericResp{
			Type: "error",
			Msg:  fmt.Sprintf("room %s alredy booked", roomID),
		})
	}

	booking := models.Booking{
		UserID:     user.ID,
		RoomID:     roomID,
		NumPersons: params.NumPersons,
		FromDate:   params.FromDate,
		TillDate:   params.TillDate,
	}

	insertBooking, err := h.bookingUsecase.CreateBooking(c.Context(), &booking)
	if err != nil {
		return err
	}

	return c.JSON(insertBooking)
}

func (h *RoomHandler) isRoomAvailableForBooking(ctx context.Context, roomID primitive.ObjectID, params domain.BookRoomParams) (bool, error) {
	filter := domain.Map{
		"room_id": roomID,
		"$or": []bson.M{
			{
				"from_date": bson.M{"$lt": params.TillDate},
				"till_date": bson.M{"$gt": params.FromDate},
			},
		},
	}

	bookings, err := h.bookingUsecase.GetBookings(ctx, filter)
	if err != nil {
		return false, err
	}

	return len(bookings) == 0, nil
}

func (h *RoomHandler) HandleGetRooms(c *fiber.Ctx) error {
	rooms, err := h.roomUsecase.GetRooms(c.Context(), domain.Map{})
	if err != nil {
		return err
	}
	return c.JSON(rooms)
}

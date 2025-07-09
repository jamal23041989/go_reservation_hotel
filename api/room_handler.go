package api

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/jamal23041989/go_reservation_hotel/db"
	"github.com/jamal23041989/go_reservation_hotel/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"time"
)

type GenericResp struct {
	Type string `json:"type"`
	Msg  string `json:"msg"`
}

type BookRoomParams struct {
	FromDate   time.Time `json:"from_date"`
	TillDate   time.Time `json:"till_date"`
	NumPersons int       `json:"num_persons"`
}

func (p BookRoomParams) validate() error {
	now := time.Now()
	if now.After(p.FromDate) || now.After(p.TillDate) {
		return fmt.Errorf("cannot book a room in the past")
	}
	return nil
}

type RoomHandler struct {
	store *db.Store
}

func NewRoomHandler(store *db.Store) *RoomHandler {
	return &RoomHandler{
		store: store,
	}
}

func (h *RoomHandler) HandleBookRoom(c *fiber.Ctx) error {
	var params BookRoomParams
	if err := c.BodyParser(&params); err != nil {
		return err
	}
	if err := params.validate(); err != nil {
		return err
	}

	roomID, err := db.ConvertToObjectID(c.Params("id"))
	if err != nil {
		return err
	}

	user, ok := c.Context().Value("user").(*types.User)
	if !ok {
		return c.Status(http.StatusInternalServerError).JSON(GenericResp{
			Type: "error",
			Msg:  "internal server error",
		})
	}

	ok, err = h.isRoomAvailableForBooking(c.Context(), roomID, params)
	if err != nil {
		return err
	}
	if !ok {
		return c.Status(http.StatusBadRequest).JSON(GenericResp{
			Type: "error",
			Msg:  fmt.Sprintf("room %s alredy booked", roomID),
		})
	}

	booking := types.Booking{
		UserID:     user.ID,
		RoomID:     roomID,
		NumPersons: params.NumPersons,
		FromDate:   params.FromDate,
		TillDate:   params.TillDate,
	}

	insertBooking, err := h.store.Booking.InsertBooking(c.Context(), &booking)
	if err != nil {
		return err
	}

	return c.JSON(insertBooking)
}

func (h *RoomHandler) isRoomAvailableForBooking(ctx context.Context, roomID primitive.ObjectID, params BookRoomParams) (bool, error) {
	filter := bson.M{
		"room_id": roomID,
		"from_date": bson.M{
			"$gte": params.FromDate,
		},
		"till_date": bson.M{
			"$lte": params.TillDate,
		},
	}

	bookings, err := h.store.Booking.GetBookings(ctx, filter)
	if err != nil {
		return false, err
	}

	return len(bookings) == 0, nil
}

func (h *RoomHandler) HandleGetRooms(c *fiber.Ctx) error {
	rooms, err := h.store.Room.GetRooms(c.Context(), bson.M{})
	if err != nil {
		return err
	}
	return c.JSON(rooms)
}

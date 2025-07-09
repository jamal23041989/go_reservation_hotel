package api

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/jamal23041989/go_reservation_hotel/db/fixtures"
	"github.com/jamal23041989/go_reservation_hotel/middleware"
	"github.com/jamal23041989/go_reservation_hotel/types"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestGetBookings(t *testing.T) {
	db := setup(t)
	defer db.teardown(t)

	var (
		adminUser      = fixtures.AddUser(db.Store, "admin", "admin", true)
		user           = fixtures.AddUser(db.Store, "james", "foo", false)
		hotel          = fixtures.AddHotel(db.Store, "bar hotel", "a", 4, nil)
		room           = fixtures.AddRoom(db.Store, "small", true, 4.4, hotel.ID)
		from           = time.Now()
		till           = from.AddDate(0, 0, 5)
		booking        = fixtures.AddBooking(db.Store, user.ID, room.ID, from, till)
		app            = fiber.New()
		admin          = app.Group("/", middleware.JwtAuthentication(db.User), middleware.AdminAuth)
		bookingHandler = NewBookingHandler(db.Store)
	)

	admin.Get("/", bookingHandler.HandleGetBookings)
	req := httptest.NewRequest("GET", "/", nil)
	req.Header.Add("X-Api-Token", CreateTokenFromUser(adminUser))

	resp, err := app.Test(req)
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected http status of 200 but got %d", resp.StatusCode)
	}

	var bookings []*types.Booking
	if err := json.NewDecoder(resp.Body).Decode(&bookings); err != nil {
		t.Fatal(err)
	}

	fmt.Println("*************", len(bookings))
	fmt.Printf("%+v\n", bookings[0])
	fmt.Printf("%+v\n", booking)

	if len(bookings) != 1 {
		t.Fatalf("expected 1 booking but got %d", len(bookings))
	}
	if bookings[0].ID != booking.ID {
		t.Fatalf("expected %s but got %s", bookings[0].ID, booking.ID)
	}
	if bookings[0].RoomID != booking.RoomID {
		t.Fatalf("expected %s but got %s", bookings[0].RoomID, booking.RoomID)
	}
	if bookings[0].UserID != booking.UserID {
		t.Fatalf("expected %s but got %s", bookings[0].UserID, booking.UserID)
	}
}

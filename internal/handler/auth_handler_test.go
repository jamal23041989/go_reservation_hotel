package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/jamal23041989/go_reservation_hotel/internal/domain"
	"github.com/jamal23041989/go_reservation_hotel/internal/domain/models"
	"github.com/jamal23041989/go_reservation_hotel/internal/repository/mongodb/fixtures"
	"github.com/jamal23041989/go_reservation_hotel/internal/usecase"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func insertTestUser(t *testing.T, userCase usecase.UserUsecase) *models.User {
	user, err := models.NewUserFromParams(models.CreateUserParams{
		Email:     "james2@gmail.com",
		FirstName: "james",
		LastName:  "scot",
		Password:  "supersecurepassword",
	})
	if err != nil {
		t.Fatal(err)
	}

	if _, err := userCase.CreateUser(context.Background(), user); err != nil {
		t.Fatal(err)
	}

	return user
}

func TestAuthenticateSuccess(t *testing.T) {
	tdb := setup(t)
	defer tdb.teardown(t)
	insertedUser := fixtures.AddUser(tdb.Store, "james", "foo", false)

	app := fiber.New()
	authHandler := NewAuthHandler(tdb.User)
	app.Post("/", authHandler.HandleAuthenticate)

	authParams := domain.AuthParams{
		Email:    "james@foo.com",
		Password: "james_foo",
	}

	b, _ := json.Marshal(authParams)
	req := httptest.NewRequest("POST", "/", bytes.NewReader(b))
	req.Header.Add("Content-Type", "application/json")

	resp, err := app.Test(req)
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected http status of 200 but got %d", err)
	}

	var authResp domain.AuthResponse
	if err := json.NewDecoder(resp.Body).Decode(&authResp); err != nil {
		t.Fatal(err)
	}

	if authResp.Token == "" {
		t.Fatalf("expected the JWT token to be present in the auth response")
	}

	insertedUser.EncryptedPassword = ""
	if !reflect.DeepEqual(insertedUser, authResp.User) {
		t.Fatalf("expected the user to be the insertedUser")
	}
}

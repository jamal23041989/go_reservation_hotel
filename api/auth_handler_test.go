package api

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/jamal23041989/go_reservation_hotel/db"
	"github.com/jamal23041989/go_reservation_hotel/db/fixtures"
	"github.com/jamal23041989/go_reservation_hotel/types"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func insertTestUser(t *testing.T, userStore db.UserStore) *types.User {
	user, err := types.NewUserFromParams(types.CreateUserParams{
		Email:     "james2@gmail.com",
		FirstName: "james",
		LastName:  "scot",
		Password:  "supersecurepassword",
	})
	if err != nil {
		t.Fatal(err)
	}

	if _, err := userStore.InsertUser(context.Background(), user); err != nil {
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

	authParams := AuthParams{
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

	var authResp AuthResponse
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

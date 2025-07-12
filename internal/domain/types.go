package domain

import (
	"github.com/jamal23041989/go_reservation_hotel/internal/domain/entity"
	"time"
)

type Map map[string]any

type ResourceResp struct {
	Results int `json:"results"`
	Data    any `json:"data"`
	Page    int `json:"page"`
}

type Pagination struct {
	Limit int64
	Page  int64
}

type HotelFilter struct {
	Pagination
	Rating int
}

type GenericResp struct {
	Type string `json:"type"`
	Msg  string `json:"msg"`
}

type BookRoomParams struct {
	FromDate   time.Time `json:"from_date"`
	TillDate   time.Time `json:"till_date"`
	NumPersons int       `json:"num_persons"`
}

type AuthParams struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthResponse struct {
	User  *entity.User `json:"user"`
	Token string       `json:"token"`
}

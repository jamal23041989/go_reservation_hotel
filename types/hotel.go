package types

import "go.mongodb.org/mongo-driver/bson/primitive"

type RoomType int

const (
	_ RoomType = iota
	SingleRoomType
	DoubleRoomType
	SeaSideRoomType
	DeluxeRoomType
)

type Room struct {
	ID primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	//Type      RoomType           `bson:"type" json:"type"`
	Size      string             `bson:"size" json:"size"`
	Seaside   bool               `bson:"seaside" json:"seaside"`
	BasePrice float64            `bson:"base_price" json:"base_price"`
	Price     float64            `bson:"price" json:"price"`
	HotelID   primitive.ObjectID `bson:"hotel_id" json:"hotel_id"`
}

type Hotel struct {
	ID       primitive.ObjectID   `bson:"_id,omitempty" json:"id,omitempty"`
	Name     string               `bson:"name" json:"name"`
	Location string               `bson:"location" json:"location"`
	Rooms    []primitive.ObjectID `bson:"rooms" json:"rooms"`
	Rating   int                  `bson:"rating" json:"rating"`
}

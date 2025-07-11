package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Room struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Size      string             `bson:"size" json:"size"`
	Seaside   bool               `bson:"seaside" json:"seaside"`
	BasePrice float64            `bson:"base_price" json:"base_price"`
	Price     float64            `bson:"price" json:"price"`
	HotelID   primitive.ObjectID `bson:"hotel_id" json:"hotel_id"`
	Available bool               `bson:"-" json:"available"`
}

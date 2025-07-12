package mongodb

const DbName = "hotel-reservation"

type Store struct {
	User    *MongoUserRepository
	Hotel   *MongoHotelRepository
	Room    *MongoRoomRepository
	Booking *MongoBookingRepository
}

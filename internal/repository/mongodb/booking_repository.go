package mongodb

import (
	"context"
	"github.com/jamal23041989/go_reservation_hotel/internal/domain"
	"github.com/jamal23041989/go_reservation_hotel/internal/domain/models"
	"github.com/jamal23041989/go_reservation_hotel/pkg"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const bookingColl = "bookings"

type MongoBookingRepository struct {
	client *mongo.Client
	coll   *mongo.Collection
}

func NewMongoBookingRepository(client *mongo.Client) *MongoBookingRepository {
	return &MongoBookingRepository{
		client: client,
		coll:   client.Database(DbName).Collection(bookingColl),
	}
}

func (s *MongoBookingRepository) CreateBooking(ctx context.Context, booking *models.Booking) (*models.Booking, error) {
	resp, err := s.coll.InsertOne(ctx, booking)
	if err != nil {
		return nil, err
	}
	booking.ID = resp.InsertedID.(primitive.ObjectID)
	return booking, nil
}

func (s *MongoBookingRepository) GetBookings(ctx context.Context, filter domain.Map) ([]*models.Booking, error) {
	cursor, err := s.coll.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var bookings []*models.Booking
	if err := cursor.All(ctx, &bookings); err != nil {
		return nil, err
	}

	return bookings, nil
}

func (s *MongoBookingRepository) GetBookingByID(ctx context.Context, id string) (*models.Booking, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, pkg.ErrInvalidID()
	}

	var booking models.Booking
	if err := s.coll.FindOne(ctx, domain.Map{"_id": objectID}).Decode(&booking); err != nil {
		return nil, err
	}

	return &booking, nil
}

func (s *MongoBookingRepository) UpdateBooking(ctx context.Context, id string, update domain.Map) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return pkg.ErrInvalidID()
	}

	res, err := s.coll.UpdateByID(ctx, objectID, domain.Map{"$set": update})
	if err != nil {
		return err
	}
	if res.MatchedCount == 0 {
		return mongo.ErrNoDocuments
	}

	return nil
}

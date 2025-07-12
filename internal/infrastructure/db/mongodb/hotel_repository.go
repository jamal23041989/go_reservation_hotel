package mongodb

import (
	"context"
	"github.com/jamal23041989/go_reservation_hotel/internal/domain"
	"github.com/jamal23041989/go_reservation_hotel/internal/domain/entity"
	"github.com/jamal23041989/go_reservation_hotel/pkg"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const hotelColl = "hotels"

type MongoHotelRepository struct {
	client *mongo.Client
	coll   *mongo.Collection
}

func NewMongoHotelRepository(client *mongo.Client) *MongoHotelRepository {
	return &MongoHotelRepository{
		client: client,
		coll:   client.Database(DbName).Collection(hotelColl),
	}
}

func (s *MongoHotelRepository) CreateHotel(ctx context.Context, hotel *entity.Hotel) (*entity.Hotel, error) {
	resp, err := s.coll.InsertOne(ctx, hotel)
	if err != nil {
		return nil, err
	}

	hotel.ID = resp.InsertedID.(primitive.ObjectID)
	return hotel, nil
}

func (s *MongoHotelRepository) UpdateHotel(ctx context.Context, filter domain.Map, update domain.Map) error {
	res, err := s.coll.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	if res.MatchedCount == 0 {
		return mongo.ErrNoDocuments
	}

	return nil
}

func (s *MongoHotelRepository) GetAllHotels(
	ctx context.Context,
	filter domain.Map,
	pag *domain.Pagination,
) ([]*entity.Hotel, error) {
	if pag.Limit <= 0 {
		pag.Limit = 10
	}
	if pag.Page <= 0 {
		pag.Page = 1
	}

	opts := options.FindOptions{}
	opts.SetSkip((pag.Page - 1) * pag.Limit)
	opts.SetLimit(pag.Limit)

	cursor, err := s.coll.Find(ctx, &filter, &opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var hotels []*entity.Hotel
	if err := cursor.All(ctx, &hotels); err != nil {
		return nil, err
	}

	return hotels, nil
}

func (s *MongoHotelRepository) GetByIDHotel(ctx context.Context, id string) (*entity.Hotel, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, pkg.ErrInvalidID()
	}

	var hotel entity.Hotel
	if err := s.coll.FindOne(ctx, domain.Map{"_id": objectID}).Decode(&hotel); err != nil {
		return nil, err
	}

	return &hotel, nil
}

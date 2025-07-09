package db

import (
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	NameDb = "hotel-reservation"
	UriDb  = "mongodb://localhost:27017"
)

type Store struct {
	User    UserStore
	Hotel   HotelStore
	Room    RoomStore
	Booking BookingStore
}

// ConvertToObjectID преобразует строковый идентификатор документа в формат MongoDB ObjectID.
// Данная функция необходима для корректной работы с идентификаторами в MongoDB, где каждый документ
// использует 12-байтовый бинарный формат вместо строкового представления.
//
// Параметры:
//   - id - строка, содержащая 24-символьный hex-идентификатор (например, "507f1f77bcf86cd799439011")
//
// Возвращает:
//   - primitive.ObjectID: бинарное представление идентификатора для MongoDB
//   - error: ошибку, если строка имеет неверный формат (длина != 24, не hex-символы)
//
// Особенности:
//   - Генерирует ошибку при передаче пустой строки или некорректных данных
//   - Для валидации использует стандартный метод ObjectIDFromHex
//   - Возвращает специальное значение primitive.NilObjectID при ошибках
//
// Пример использования:
//
//	objectID, err := ConvertToObjectID("5f7d9c8e6a3b1d0e4c8b9a7d")
//	if err != nil {
//	    // обработка ошибки
//	}
func ConvertToObjectID(id string) (primitive.ObjectID, error) {
	// Проверка на пустой идентификатор перед преобразованием
	if id == "" {
		return primitive.NilObjectID, errors.New("empty id provided")
	}

	// Стандартное преобразование hex-строки в ObjectID
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		// Возвращаем типизированную ошибку с контекстом
		return primitive.NilObjectID, fmt.Errorf("invalid id format: %w", err)
	}

	return objectID, nil
}

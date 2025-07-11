package handler

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/jamal23041989/go_reservation_hotel/internal/domain/models"
	"github.com/jamal23041989/go_reservation_hotel/internal/usecase"
	"github.com/jamal23041989/go_reservation_hotel/pkg"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserHandler struct {
	userUsecase usecase.UserUsecase
}

func NewUserHandler(userUsecase *usecase.UserUsecase) *UserHandler {
	return &UserHandler{
		userUsecase: *userUsecase,
	}
}

func (h *UserHandler) HandleGetUserByID(c *fiber.Ctx) error {
	id := c.Params("id")
	user, err := h.userUsecase.GetUserByID(c.Context(), id)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return c.Status(pkg.ErrNotFound().Code).JSON(pkg.ErrNotFound())
		}
		return err
	}

	return c.JSON(user)
}

func (h *UserHandler) HandleGetUsers(c *fiber.Ctx) error {
	users, err := h.userUsecase.GetUsers(c.Context())
	if err != nil {
		return err
	}
	return c.JSON(users)
}

func (h *UserHandler) HandleCreateUser(c *fiber.Ctx) error {
	var params models.CreateUserParams
	if err := c.BodyParser(&params); err != nil {
		return pkg.ErrBadRequest()
	}
	if errorsValid := params.Validate(); len(errorsValid) > 0 {
		return c.JSON(errorsValid)
	}

	user, err := models.NewUserFromParams(params)
	if err != nil {
		return err
	}

	createdUser, err := h.userUsecase.CreateUser(c.Context(), user)
	if err != nil {
		return err
	}

	return c.JSON(createdUser)
}

func (h *UserHandler) HandleUpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")

	var updateData models.UpdateUserParams
	if err := c.BodyParser(&updateData); err != nil {
		return pkg.ErrBadRequest()
	}

	if err := h.userUsecase.UpdateUser(c.Context(), id, updateData.ToBSON()); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return pkg.ErrNotFound()
		}
		return err
	}

	return c.JSON(map[string]string{"updated": id})
}

func (h *UserHandler) HandleDeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")

	if err := h.userUsecase.DeleteUser(c.Context(), id); err != nil {
		return err
	}

	return c.JSON(map[string]string{"deleted": id})
}

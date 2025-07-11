package middleware

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/jamal23041989/go_reservation_hotel/db"
	"github.com/jamal23041989/go_reservation_hotel/pkg"
	"net/http"
	"os"
	"strconv"
	"time"
)

func JwtAuthentication(userStore db.UserStore) fiber.Handler {
	return func(c *fiber.Ctx) error {
		token, ok := c.GetReqHeaders()["X-Api-Token"]
		if !ok {
			jwtError := pkg.NewError(http.StatusUnauthorized, "missing token")
			return c.Status(jwtError.Code).JSON(jwtError)
		}

		claims, err := validateToken(token[0])
		if err != nil {
			jwtError := pkg.NewError(http.StatusUnauthorized, err.Error())
			return c.Status(jwtError.Code).JSON(jwtError)
		}

		expiresValue := claims["expires"]

		var expires int64
		switch v := expiresValue.(type) {
		case float64:
			expires = int64(v)
		case string:
			expiresInt, err := strconv.ParseInt(v, 10, 64)
			if err != nil {
				fmt.Println("invalid expires format:", v)
				jwtError := pkg.NewError(http.StatusUnauthorized, "invalid token format")
				return c.Status(jwtError.Code).JSON(jwtError)
			}
			expires = expiresInt
		default:
			fmt.Printf("unexpected expires type: %T\n", v)
			jwtError := pkg.NewError(http.StatusUnauthorized, "invalid token format")
			return c.Status(jwtError.Code).JSON(jwtError)
		}

		if time.Now().Unix() > expires {
			jwtError := pkg.NewError(http.StatusUnauthorized, "token expired")
			return c.Status(jwtError.Code).JSON(jwtError)
		}

		userID := claims["user_id"].(string)
		user, err := userStore.GetUserByID(c.Context(), userID)
		if err != nil {
			return err
		}
		c.Context().SetUserValue("user", user)

		return c.Next()
	}

}

func validateToken(tokenStr string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			fmt.Println("invalid signing method", token.Header["alg"])
			return nil, pkg.ErrUnauthorized()
		}
		secret := os.Getenv("JWT_SECRET")
		return []byte(secret), nil
	})

	if err != nil {
		fmt.Println("failed to parse JWT token:", err)
		return nil, pkg.ErrUnauthorized()
	}

	if !token.Valid {
		fmt.Println("invalid token:")
		return nil, pkg.ErrUnauthorized()
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, pkg.ErrUnauthorized()
	}

	return claims, nil
}

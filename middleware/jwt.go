package middleware

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"strconv"
	"time"
)

func JwtAuthentication(c *fiber.Ctx) error {
	token, ok := c.GetReqHeaders()["X-Api-Token"]
	if !ok {
		return c.Status(fiber.StatusUnauthorized).SendString("missing token")
	}

	claims, err := validateToken(token[0])
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).SendString(err.Error())
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
			return c.Status(fiber.StatusUnauthorized).SendString("invalid token format")
		}
		expires = expiresInt
	default:
		fmt.Printf("unexpected expires type: %T\n", v)
		return c.Status(fiber.StatusUnauthorized).SendString("invalid token format")
	}

	if time.Now().Unix() > expires {
		return c.Status(fiber.StatusUnauthorized).SendString("token expired")
	}

	return c.Next()
}

func validateToken(tokenStr string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			fmt.Println("invalid signing method", token.Header["alg"])
			return nil, fmt.Errorf("unauthorized")
		}
		secret := os.Getenv("JWT_SECRET")
		return []byte(secret), nil
	})

	if err != nil {
		fmt.Println("failed to parse JWT token:", err)
		return nil, fmt.Errorf("unauthorized")
	}

	if !token.Valid {
		fmt.Println("invalid token:")
		return nil, fmt.Errorf("unauthorized")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("unauthorized")
	}

	return claims, nil
}

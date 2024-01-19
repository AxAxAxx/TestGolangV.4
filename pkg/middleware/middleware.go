package middleware

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

var secretKey = []byte("secret-key")
var refreshSecretKey = []byte("refresh-secret-key")
var TokenBlacklist = make(map[string]bool)

/*-----------GenToken----------------*/
func GenerateTokens(username string, id, role_id int) (string, string) {
	accessToken := generateAccessToken(username, id, role_id)
	refreshToken := generateRefreshToken(username, id, role_id)

	return accessToken, refreshToken
}

func generateAccessToken(username string, id, role_id int) string {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = username
	claims["user_id"] = id
	claims["role_id"] = role_id
	claims["exp"] = time.Now().Add(time.Minute * 10).Unix()

	tokenString, _ := token.SignedString([]byte(secretKey))
	return tokenString
}

func generateRefreshToken(username string, id, role_id int) string {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = username
	claims["user_id"] = id
	claims["role_id"] = role_id
	claims["exp"] = time.Now().Add(time.Hour * 24 * 7).Unix()

	tokenString, _ := token.SignedString([]byte(refreshSecretKey))
	return tokenString
}

/*-----------VerifyToken----------------*/
func VerifyToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}

/*-----------Middleware----------------*/
func AuthMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		s := c.Get("Authorization")
		token := strings.TrimPrefix(s, "Bearer ")
		claims, err := VerifyToken(token)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
		}
		if TokenBlacklist[token] {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
		}
		c.Locals("username", claims)
		return c.Next()
	}
}

func AdminOnly() fiber.Handler {
	return func(c *fiber.Ctx) error {
		claims, ok := c.Locals("username").(jwt.MapClaims)
		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
		}
		roleValue, exists := claims["role_id"]
		if !exists || roleValue == nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Role not found or nil"})
		}
		role_id := roleValue.(float64)
		fmt.Println(role_id)
		if role_id != 2 {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "Access forbidden for non-admin users"})
		}
		return c.Next()
	}
}

/*-----------HashPassword----------------*/
func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func VerifyPassword(hashedPassword, inputPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(inputPassword))
	return err
}

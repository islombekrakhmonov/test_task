package security

import (
	"crypto/rand"
	"encoding/base64"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GenerateSecretKey(length int) (string, error) {
	key := make([]byte, length)
	_, err := rand.Read(key)
	if err != nil {
		return "", err
	}
	secretKey := base64.URLEncoding.EncodeToString(key)
	return secretKey, nil
}

func GenerateToken(userID, username string) (string, error) {
	
	// Создание нового токена с указанием алгоритма шифрования и ключа
	claims := jwt.MapClaims{
		"user_id": userID,
		"login":   username,
		"exp":     time.Now().Add(time.Hour).Unix(), // Токен истекает через час
	}

	// Создание токена с указанием метода шифрования HS256 и claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte("kawasaki"))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func GenerateRefreshToken(userID string) (string, error) {
	refreshTokenLength := 32
	tokenBytes := make([]byte, refreshTokenLength)

	_, err := rand.Read(tokenBytes)
	if err != nil {
		return "", err
	}

	token := base64.URLEncoding.EncodeToString(tokenBytes)

	return token, nil
}
package jwt

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

type JWTManager[T any] struct {
	secretKey string
}

func NewJWTManager[T any](secretKey string) *JWTManager[T] {
	return &JWTManager[T]{secretKey: secretKey}
}

func (m *JWTManager[T]) Generate(content T) (string, error) {
	claims := jwt.MapClaims{
		"content": content,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(m.secretKey))
}

func (m *JWTManager[T]) Validate(tokenStr string) (*T, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(m.secretKey), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("invalid token claims")
	}

	content, ok := claims["content"].(T)
	if !ok {
		return nil, fmt.Errorf("invalid user_id claim")
	}

	return &content, nil
}

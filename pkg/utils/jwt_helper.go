package utils

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"time"
)

var jwtSecret = []byte("secret_key")

type Claims struct {
	UserID uuid.UUID `json:"user_id"`
	Role   string    `json:"role"`
	IP     string    `json:"ip"`
	jwt.RegisteredClaims
}

func GenerateJWT(userID uuid.UUID, role string, ip string) (string, string, error) {
	jti, err := generateJTI()
	if err != nil {
		return "", "", err
	}

	claims := Claims{
		UserID: userID,
		Role:   role,
		IP:     ip,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(15 * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ID:        jti,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	signedToken, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", "", err
	}

	return signedToken, jti, nil
}

func ValidateJWT(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}

// refresh: raw + bcrypt
func GenerateSecureRefreshToken() (rawToken string, hashedToken string, err error) {
	randomBytes := make([]byte, 32)
	_, err = rand.Read(randomBytes)
	if err != nil {
		return "", "", err
	}

	rawToken = base64.URLEncoding.EncodeToString(randomBytes)

	hashed, err := bcrypt.GenerateFromPassword([]byte(rawToken), bcrypt.DefaultCost)
	if err != nil {
		return "", "", err
	}

	return rawToken, string(hashed), nil
}

func generateJTI() (string, error) {
	randomBytes := make([]byte, 16)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(randomBytes), nil
}

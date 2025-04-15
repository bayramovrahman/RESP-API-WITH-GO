package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// Sabit bir gizli anahtar tanımlanıyor; JWT imzalama işlemlerinde kullanılacak
const secretKey = "supersecret"

func GenerateToken(email string, userId int64) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,                                 // Token içerisine e-posta bilgisi ekleniyor
		"userId": userId,                                // Token içerisine kullanıcı ID'si ekleniyor
		"exp":    time.Now().Add(time.Hour * 12).Unix(), // Token geçerlilik süresi şu andan itibaren 2 saat olarak ayarlanıyor
	})

	// Token, belirlenen gizli anahtar kullanılarak imzalanıyor ve string olarak döndürülüyor
	return token.SignedString([]byte(secretKey))
}

func VerifyToken(token string) (int64, error) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("unexpected signing method")
		}

		return []byte(secretKey), nil
	})

	if err != nil {
		return 0, errors.New("could not parse token")
	}

	tokenIsValid := parsedToken.Valid
	if !tokenIsValid {
		return 0, errors.New("invalid token")
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("invalid token claims")
	}

	// email := claims["email"].(string)
	userId := int64(claims["userId"].(float64))

	return userId, nil
}

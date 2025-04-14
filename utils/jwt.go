package utils

import (
	"time"
	"github.com/golang-jwt/jwt/v5"
)

// Sabit bir gizli anahtar tanımlanıyor; JWT imzalama işlemlerinde kullanılacak
const secretKey = "supersecret"

func GenerateToken(email string, userId int64) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,                                     // Token içerisine e-posta bilgisi ekleniyor
		"userId": userId,                                   // Token içerisine kullanıcı ID'si ekleniyor
		"exp": time.Now().Add(time.Hour * 12).Unix(),       // Token geçerlilik süresi şu andan itibaren 2 saat olarak ayarlanıyor
	})

	// Token, belirlenen gizli anahtar kullanılarak imzalanıyor ve string olarak döndürülüyor
	return token.SignedString([]byte(secretKey))
}

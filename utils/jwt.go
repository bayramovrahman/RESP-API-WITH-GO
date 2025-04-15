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
	// jwt.Parse fonksiyonu ile token çözülmeye çalışılır.
	// İkinci parametre olan anonim fonksiyon, imzalama yönteminin doğru olup olmadığını kontrol eder
	// ve doğrulama için kullanılacak gizli anahtarı (secretKey) döner.
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		// Token'ın imzalama algoritması HMAC türünde mi kontrol edilir.
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			// Eğer beklenmeyen bir imzalama yöntemi varsa, hata döndürülür.
			return nil, errors.New("unexpected signing method")
		}

		// Token doğrulaması için gizli anahtar döndürülür.
		return []byte(secretKey), nil
	})

	// Token çözümleme sırasında hata oluşmuşsa, hata döndürülür.
	if err != nil {
		return 0, errors.New("could not parse token")
	}

	// Token'ın geçerli olup olmadığı kontrol edilir.
	tokenIsValid := parsedToken.Valid
	if !tokenIsValid {
		return 0, errors.New("invalid token")
	}

	// Token'dan claims (taşıdığı veriler) alınır ve doğru türde olup olmadığı kontrol edilir.
	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("invalid token claims")
	}

	// Claims içinden userId alınır.
	// JWT içerisinde sayılar float64 formatında tutulduğu için önce float64'e cast edilir,
	// ardından int64'e dönüştürülerek geri döndürülür.
	userId := int64(claims["userId"].(float64))

	// Geçerli bir userId elde edildiyse, hata olmadan geriye döndürülür.
	return userId, nil
}

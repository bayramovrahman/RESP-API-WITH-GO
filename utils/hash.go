package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))

	/*
		Hata umurumda değil, bunun yerine bu işlemi burada döndürerek,
		şifre geçersizse false döndürüyorum çünkü şifre geçerliyse
		hata nil olacak ve şifre geçersizse nil olmayacak.
	*/

	return err == nil
}

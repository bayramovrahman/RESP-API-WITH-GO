package middlewares

import (
	"net/http"
	"example.com/rest-api/utils"
	"github.com/gin-gonic/gin"
)

func Authenticate(context *gin.Context) {
	// İstek başlığından "Authorization" header'ını alır (genellikle "Bearer <token>" formatında olur).
	token := context.Request.Header.Get("Authorization")

	// Eğer token header'da yoksa, kullanıcı yetkisiz sayılır. İstek iptal edilir ve 401 hatası döner.
	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not authorized"})
		return
	}

	// Token doğrulaması yapılır. Bu işlem utils paketindeki VerifyToken fonksiyonu ile gerçekleştirilir.
	// Bu fonksiyon token'ı çözümleyip geçerli olup olmadığını kontrol eder ve geçerliyse userId döner.
	userId, err := utils.VerifyToken(token)
	if err != nil {
		// Eğer token geçersizse ya da doğrulama sırasında bir hata oluşursa, yine 401 Unauthorized hatası döner.
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not authorized"})
		return
	}

	// Eğer token geçerliyse, elde edilen userId context'e "userId" anahtarıyla eklenir.
	// Böylece sonraki handler'lar bu kullanıcı ID'sine erişebilir.
	context.Set("userId", userId)

	// İstek işlenmeye devam eder, bir sonraki middleware ya da handler'a geçilir.
	context.Next()
}
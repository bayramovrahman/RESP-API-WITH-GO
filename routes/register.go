package routes

import (
	"net/http"
	"strconv"
	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"
)

func registerForEvent(context *gin.Context) {
	userId := context.GetInt64("userId")

	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		// Eğer dönüşümde hata oluşursa (geçersiz id formatı gibi), 400 Bad Request hatası döner
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id: " + err.Error()})
		return
	}

	event, err := models.GetEventById(eventId) // Belirtilen ID'ye sahip etkinliğin veritabanında olup olmadığını kontrol eder
	if err != nil {
		// Eğer veri çekilirken hata olursa (örneğin: event yoksa), 500 Internal Server Error döner
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event: " + err.Error()})
		return
	}

	err = event.Register(userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not register user for event: " + err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Successfully registered!"})
}

func cancelRegistration(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		// Eğer dönüşümde hata oluşursa (geçersiz id formatı gibi), 400 Bad Request hatası döner
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id: " + err.Error()})
		return
	}

	var event models.Event
	event.ID = eventId

	err = event.CancelRegistration(userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not cancel registration: " + err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "The register successfully cancelled!"})
}

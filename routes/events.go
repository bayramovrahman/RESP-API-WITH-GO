package routes

import (
	"net/http"
	"strconv"
	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"
)

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		// context.JSON(http.StatusBadRequest, gin.H{"message": "Could not fetch events. Please try again"})
		return
	}

	context.JSON(http.StatusOK, events)
}

func getEvent(context *gin.Context) {
	// URL'den gelen "id" parametresini alır ve int64 türüne çevirir (örneğin: /events/5 → id = 5)
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		// context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id"})
		return
	}

	event, err := models.GetEventById(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		// context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event"})
		return
	}

	context.JSON(http.StatusOK, event)
}

func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		// context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}

	event.ID = 1
	event.UserID = 1

	err = event.Save()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		// context.JSON(http.StatusBadRequest, gin.H{"message": "Could not create event. Please try again"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "The event successfully created!", "event": event})
}

func updateEvent(context *gin.Context) {
	// URL'den gelen "id" parametresini alır ve int64 türüne çevirir (örneğin: /events/5 → id = 5)
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		// Eğer dönüşümde hata oluşursa (geçersiz id formatı gibi), 400 Bad Request hatası döner
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		// context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id"})
		return
	}

	// Belirtilen ID'ye sahip etkinliğin veritabanında olup olmadığını kontrol eder
	_, err = models.GetEventById(eventId)

	if err != nil {
		// Eğer veri çekilirken hata olursa (örneğin: event yoksa), 500 Internal Server Error döner
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		// context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event"})
		return
	}

	// İstekten gelen JSON verisini `updatedEvent` nesnesine bağlamaya çalışır
	var updatedEvent models.Event
	err = context.ShouldBindJSON(&updatedEvent)

	if err != nil {
		// Eğer JSON verisi beklenen formatta değilse, 400 Bad Request hatası döner
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		// context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}

	// Güncellenmek istenen etkinliğe ait ID'yi ayarlar (gelen JSON'da ID olmasa bile doğru etkinlik güncellenir)
	updatedEvent.ID = eventId

	// Etkinlik verisini güncellemeye çalışır (veritabanı işlemi)
	err = updatedEvent.Update()

	if err != nil {
		// Güncelleme sırasında hata olursa, 500 Internal Server Error döner
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		// context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not update event"})
		return
	}

	// Her şey başarılıysa 200 OK döner ve başarı mesajı gönderilir
	context.JSON(http.StatusOK, gin.H{"message":"Event updated successfully!"})

}

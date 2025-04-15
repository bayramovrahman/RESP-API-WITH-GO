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
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch events: " + err.Error()})
		return
	}

	context.JSON(http.StatusOK, events)
}

func getEvent(context *gin.Context) {
	// URL'den gelen "id" parametresini alır ve int64 türüne çevirir (örneğin: /events/5 → id = 5)
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id: " + err.Error()})
		return
	}

	event, err := models.GetEventById(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event: " + err.Error()})
		return
	}

	context.JSON(http.StatusOK, event)
}

func createEvent(context *gin.Context) {
	var event models.Event
	
	err := context.ShouldBindJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data: " + err.Error()})
		return
	}

	userId := context.GetInt64("userId")
	event.UserID = userId

	err = event.Save()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not create event: " + err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "The event successfully created!", "event": event})
}

func updateEvent(context *gin.Context) {
	// URL'den gelen "id" parametresini alır ve int64 türüne çevirir (örneğin: /events/5 → id = 5)
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		// Eğer dönüşümde hata oluşursa (geçersiz id formatı gibi), 400 Bad Request hatası döner
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id: " + err.Error()})
		return
	}

	// Belirtilen ID'ye sahip etkinliğin veritabanında olup olmadığını kontrol eder
	_, err = models.GetEventById(eventId)

	if err != nil {
		// Eğer veri çekilirken hata olursa (örneğin: event yoksa), 500 Internal Server Error döner
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event: " + err.Error()})
		return
	}

	// İstekten gelen JSON verisini `updatedEvent` nesnesine bağlamaya çalışır
	var updatedEvent models.Event
	err = context.ShouldBindJSON(&updatedEvent)

	if err != nil {
		// Eğer JSON verisi beklenen formatta değilse, 400 Bad Request hatası döner
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data: " + err.Error()})
		return
	}

	// Güncellenmek istenen etkinliğe ait ID'yi ayarlar (gelen JSON'da ID olmasa bile doğru etkinlik güncellenir)
	updatedEvent.ID = eventId

	// Etkinlik verisini güncellemeye çalışır (veritabanı işlemi)
	err = updatedEvent.Update()

	if err != nil {
		// Güncelleme sırasında hata olursa, 500 Internal Server Error döner
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not update event: " + err.Error()})
		return
	}

	// Her şey başarılıysa 200 OK döner ve başarı mesajı gönderilir
	context.JSON(http.StatusOK, gin.H{"message":"Event updated successfully!"})

}

func deleteEvent(context *gin.Context) {
	// URL'den gelen "id" parametresini alır ve int64 türüne çevirir (örneğin: /events/5 → id = 5)
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		// Eğer dönüşümde hata oluşursa (geçersiz id formatı gibi), 400 Bad Request hatası döner
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id: " + err.Error()})
		return
	}

	// Belirtilen ID'ye sahip etkinliğin veritabanında olup olmadığını kontrol eder
	event, err := models.GetEventById(eventId)

	if err != nil {
		// Eğer veri çekilirken hata olursa (örneğin: event yoksa), 500 Internal Server Error döner
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event: " + err.Error()})
		return
	}

	err = event.Delete()
	if err != nil {
		// Eğer veri çekilirken hata olursa (örneğin: event yoksa), 500 Internal Server Error döner
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not delete the event: " + err.Error()})
		return
	}

	// Her şey başarılıysa 200 OK döner ve başarı mesajı gönderilir
	context.JSON(http.StatusOK, gin.H{"message":"Event deleted successfully!"})
}

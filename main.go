package main

import (
	"net/http"
	"example.com/rest-api/database"
	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()

	// Basit bir GET endpoint'i oluştur
	server.GET("/events", getEvents)
	server.POST("/events", createEvent)

	// 8080 portunda çalıştır
	server.Run(":8080")
}

func getEvents(context *gin.Context) {
	events := models.GetAllEvents()
	context.JSON(http.StatusOK, events)
}

func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		// context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}

	event.ID = 1
	event.UserID = 1
	event.Save()
	
	context.JSON(http.StatusCreated, gin.H{"message": "The event successfully created!", "event": event})
}

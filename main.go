package main

import (
	"GO_REST_API/db"
	"GO_REST_API/models"
	"net/http"

	"github.com/gin-gonic/gin"
)


func main(){
	db.InitDB()
	server := gin.Default()
	server.GET("/events", getEvents)
	server.POST("/events", createEvent)
	server.Run(":8080")
}


func getEvents(context *gin.Context) {
	events, err:= models.GetAllEvents()
	if err != nil{
		context.JSON(http.StatusInternalServerError, gin.H{"Message": "Couldnt fetch events. Try again later"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"event": events})
}

func createEvent(context *gin.Context){
	var event models.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Message": "Bad request"})
		return
	}

	event.ID = 1
	event.UserId = 1
	err = event.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"Message": "Couldnt create event. Try again later"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"event": event})

}
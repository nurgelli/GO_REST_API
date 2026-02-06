package routes

import (
	"GO_REST_API/models"
	"fmt"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
)

func registerForEvent(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Message": "Bad request."})
		return
	}	

	event, err := models.GetEventById(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"Message": "Couldnt fetch event."})
		return
	}

	err = event.Register(userId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"Message": "Couldnt register user."})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"Message": fmt.Sprintf("USER %v Successfully registered to '%v' Event", userId, event.Name)})
}

func getAllRegisteredUser(context *gin.Context) {

}

func cancelRegisterOfUser(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Message": "Bad request."})
		return
	}	

	var event models.Event
	event.ID = eventId

	
	err = event.Cancel(userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"Message": "Couldnt cancel user."})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"Message": "Cancelled"})

}
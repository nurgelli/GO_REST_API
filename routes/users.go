package routes

import (
	"GO_REST_API/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func signup(context *gin.Context) {

	var user models.User

	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Message": "Couldnt parse request data"})
		return
	}

	err = user.Save()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Message": "Couldnt save user"})
		return
	}


	context.JSON(http.StatusCreated, gin.H{"Message": "User Successfully created!"})
}

func getUsers(context *gin.Context) {
	users, err := models.GetAllUsers()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Message": "Bad request!"})
	}

	context.JSON(http.StatusOK, gin.H{"users": users})
}

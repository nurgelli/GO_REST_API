package routes

import (
	"GO_REST_API/models"
	"GO_REST_API/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)


// * This to signup new users
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
// * To get All users for admin
func getUsers(context *gin.Context) {
	users, err := models.GetAllUsers()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Message": "Bad request!"})
	}

	context.JSON(http.StatusOK, gin.H{"users": users})
}

// * This is to login already registered users and generating jwt token 
func login(context *gin.Context) {
	var user models.User

	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Message": "Couldnt parse request data"})
		return
	}

	err = user.ValidateCredentials()
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Credential errors!"})
		return
	}

	token, err := utils.GenerateToken(user.Email, user.ID)

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Credential errors!"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"Message": "Login Successfull!", "token": token})
}

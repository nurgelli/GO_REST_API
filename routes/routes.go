package routes

import (
	"GO_REST_API/middlewares"
	"github.com/gin-gonic/gin"
)
func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)

	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticator)
	authenticated.POST("/events", createEvent)
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.DELETE("/events/:id", deleteEvent)
	authenticated.POST("/events/:id/register", registerForEvent)
	authenticated.DELETE("/events/:id/register", cancelRegisterOfUser)
	
	
	
	server.GET("/users", getUsers)
	server.POST("/signup", signup)
	server.POST("/login", login)
}
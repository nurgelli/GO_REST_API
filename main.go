package main

import (

	"github.com/gin-gonic/gin"
)

func main()  {
	server := gin.Default()

	server.GET("/events",  getAllEvents)

	server.Run(":8080")
	
}

func getAllEvents(context *gin.Context) {
	
}

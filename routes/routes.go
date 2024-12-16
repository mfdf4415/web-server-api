package routes

import "github.com/gin-gonic/gin"

func RegisterRoute(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)
	server.PUT("/events/:id",updateEvent)
	server.DELETE("/events/:id",deleteEvent)
	server.POST("/events", createEvent)
	server.POST("/signup", signup)
	server.POST("/login", login)
}
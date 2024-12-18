package routes

import (
	"net/http"
	"strconv"

	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"
)

func getEvent(c *gin.Context) {
	eventId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "cannt convert to int"})
		return
	}

	event, err := models.GetEventById(int(eventId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "cannt get event"})
		return
	}
	c.JSON(http.StatusOK, event)

}

func getEvents(c *gin.Context) {

	events, err := models.GetAllEvents()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "internal server error"})
		return
	}
	c.JSON(http.StatusOK, events)
}

func createEvent(c *gin.Context) {
	var event models.Event
	err := c.ShouldBindJSON(&event)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}
	userId := c.GetInt64("userId")
	event.UserId = userId

	err = event.Save()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "internal server create error"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "event creaated", "event": event})

}

func updateEvent(c *gin.Context) {
	eventId, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "cannot convert to int"})
		return
	}
	event, err := models.GetEventById(int(eventId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "cannt get event"})
		return
	}
	userId := c.GetInt64("userId")

	if event.UserId != userId {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "not auth for up to date"})
		return
	}

	var updatedEvent models.Event
	err = c.ShouldBindJSON(&updatedEvent)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "errrorrr gooooooy saaaaaggg"})
		return
	}

	updatedEvent.ID = eventId
	err = updatedEvent.Update()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not update event."})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Event updated successfully!"})

}

func deleteEvent(c *gin.Context) {
	eventId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "cannot conver to int."})
		return
	}

	event, err := models.GetEventById(int(eventId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "cannt get event"})
		return
	}
	userId := c.GetInt64("userId")
	if event.UserId != userId {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "not auth for delete"})
		return
	}

	err = event.Delete()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not delete event."})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Event deleted successfully!"})

}

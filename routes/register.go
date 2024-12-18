package routes

import (
	"net/http"
	"strconv"

	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"
)

func registerForEvent(c *gin.Context) {
	userId := c.GetInt64("userId")
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

	err = event.Register(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "cannt register event"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "register event successfuly"})

}

func cancelRegistration(c *gin.Context) {
	userId := c.GetInt64("userId")
	eventId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "cannot conver to int."})
		return
	}

	var event models.Event
	event.ID = eventId

	err = event.CancelRegistration(userId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not cancel registration."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Cancelled!"})
}

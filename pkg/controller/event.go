package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mgr1054/go-ticket/pkg/db"
	"github.com/mgr1054/go-ticket/pkg/models"
)

type NewEvent struct {
	Band_Name	string		`json:"band_name" binding:"required"`
	Location	string		`json:"location" binding:"required"`
	City		string		`json:"city" binding:"required"`
	Capacity	int			`json:"capacity" binding:"required"`
	Date 		string		`json:"date" binding:"required"`
}

type EventUpdate struct {
	Band_Name	string		`json:"band_name"`
	Location	string		`json:"location"`
	City		string		`json:"city"`
	Capacity	int			`json:"capacity"`
	Date 		string		`json:"date"`
}

func GetEvents (c *gin.Context) {
	var events []models.Event
	if err := db.DB.Find(&events).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Could not get events"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": events})
}

func CreateEvent (c *gin.Context) {
	var event NewEvent
	
	if err := c.ShouldBindJSON(&event); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Could not create Event"})
		return
	}
	
	newEvent := models.Event{
		Band_Name: event.Band_Name, 
		Location: event.Location, 
		City: event.City, 
		Capacity: 
		event.Capacity, 
		Date: event.Date,
	}

	if err := db.DB.Create(&newEvent).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create Event"})
		return
	} 

	c.JSON(http.StatusOK, newEvent)
}

func GetEventByID (c *gin.Context) {
	var event models.Event

	if err := db.DB.Where("id = ?", c.Param("id")).First(&event).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
		return
	}

	c.JSON(http.StatusOK, event)
}

func GetEventByLocation (c *gin.Context) {
	var event []models.Event

	if err := db.DB.Where("location = ?", c.Param("location")).Find(&event).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
		return
	}

	if len(event) < 1 {
		c.JSON(http.StatusOK, gin.H{"info": "There are no events in this location"})
		return
	}

	c.JSON(http.StatusOK, event)
}

func GetEventByDate (c *gin.Context) {
	var event []models.Event

	if err := db.DB.Where("date = ?", c.Param("date")).Find(&event).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
		return
	}

	if len(event) < 1 {
		c.JSON(http.StatusOK, gin.H{"info": "There are no events at this date"})
		return
	}

	c.JSON(http.StatusOK, event)
}

func UpdateEventById (c *gin.Context) {
	
	var event models.Event

	if err := db.DB.Where("id = ?", c.Param("id")).First(&event).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
		return
	}

	var updateEvent EventUpdate

	if err:= c.ShouldBindJSON(&updateEvent); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Event could not be updated with provided data"})
        return
    }

	if err:= db.DB.Model(&event).Updates(models.Event{
		Band_Name: updateEvent.Band_Name, 
		Location: updateEvent.Location, 
		City: updateEvent.City, 
		Capacity: updateEvent.Capacity, 
		Date: updateEvent.Date,
	}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not update Event"})
        return
	}

	c.JSON(http.StatusOK, event)
}

func DeleteEventById (c *gin.Context) {
	
	var event models.Event

	if err := db.DB.Where("id = ?", c.Param("id")).First(&event).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Event not found!"})
        return
    }

	if err := db.DB.Delete(&event).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

	c.JSON(http.StatusOK, gin.H{"message": "Event deleted"})

}
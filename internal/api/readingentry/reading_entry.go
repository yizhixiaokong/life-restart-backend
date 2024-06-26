package readingentry

import (
	"life-restart-backend/internal/dao/models"
	readingEntryService "life-restart-backend/internal/services/readingentry"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Handler struct {
	readingEntrySvr *readingEntryService.ReadingEntryService
}

func NewReadingEntryHandler(readingEntrySvr *readingEntryService.ReadingEntryService) *Handler {
	return &Handler{
		readingEntrySvr: readingEntrySvr,
	}
}

func (api *Handler) GetAll(c *gin.Context) {
	readingEntries, err := api.readingEntrySvr.GetAllReadingEntries(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get reading entries"})
		return
	}

	c.JSON(http.StatusOK, readingEntries)
}

func (api *Handler) GetByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	readingEntry, err := api.readingEntrySvr.GetReadingEntryByID(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get reading entry"})
		return
	}

	c.JSON(http.StatusOK, readingEntry)
}

func (api *Handler) Create(c *gin.Context) {
	var readingEntry models.ReadingEntry
	if err := c.ShouldBindJSON(&readingEntry); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := api.readingEntrySvr.CreateReadingEntry(c, &readingEntry)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create reading entry"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Reading entry created successfully", "id": id})
}

func (api *Handler) Update(c *gin.Context) {
	idStr := c.Param("id")
	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var readingEntry models.ReadingEntry
	if err := c.ShouldBindJSON(&readingEntry); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = api.readingEntrySvr.UpdateReadingEntry(c, id, &readingEntry)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update reading entry"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Reading entry updated successfully"})
}

func (api *Handler) Delete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	err = api.readingEntrySvr.DeleteReadingEntry(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete reading entry"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Reading entry deleted successfully"})
}

package handler

import (
	"net/http"
	"strconv"
	"time"
	timetracker "time-tracker"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// @Summary CreateItem
// @Tags item
// @Description create item
// @ID create-item
// @Accept  json
// @Produce  json
// @Param input body timetracker.CreateItemInput true "item info" example({"name":"Завтрак", "userId":1})
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/item/ [post]
func (h *Handler) createItem(c *gin.Context) {
	logrus.Debug("createItem handler called")

	var item timetracker.TimeTrackerItem
	if err := c.BindJSON(&item); err != nil {
		logrus.WithError(err).Error("Failed to bind JSON for create item")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	id, err := h.services.TimeTrackerItem.CreateItem(item)
	if err != nil {
		logrus.WithError(err).Error("Failed to create item")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create item"})
		return
	}

	logrus.WithFields(logrus.Fields{
		"item_id": id,
	}).Info("Item created successfully")

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type itemsByIdResponse struct {
	Data []timetracker.TimeTrackerItem `json:"data"`
}

// @Summary GetItemsById
// @Tags item
// @Description get item by user id
// @ID get-item-by-id
// @Accept  json
// @Produce  json
// @Param id path int true "User ID"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/item/{id} [get]
func (h *Handler) getItemsById(c *gin.Context) {
	logrus.Debug("getItemsById handler called")

	userId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logrus.WithError(err).Error("Invalid user ID")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	items, err := h.services.TimeTrackerItem.GetItemsById(userId)
	if err != nil {
		logrus.WithError(err).Error("Failed to get items by ID")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get items by ID"})
		return
	}

	logrus.WithFields(logrus.Fields{
		"user_id": userId,
		"count":   len(items),
	}).Info("Retrieved items by ID successfully")

	c.JSON(http.StatusOK, itemsByIdResponse{
		Data: items,
	})
}

// @Summary UpdateItem
// @Tags item
// @Description update item
// @ID update-item
// @Accept  json
// @Produce  json
// @Param id path int true "Item ID"
// @Param input body timetracker.UpdateItemInput true "item info" example({"name":"Завтрак"})
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/item/{id} [put]
func (h *Handler) updateItem(c *gin.Context) {
	logrus.Debug("updateItem handler called")

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logrus.WithError(err).Error("Invalid item ID")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid item ID"})
		return
	}

	var input timetracker.UpdateItemInput
	if err := c.BindJSON(&input); err != nil {
		logrus.WithError(err).Error("Failed to bind JSON for update item")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	err = h.services.TimeTrackerItem.UpdateItem(id, input)
	if err != nil {
		logrus.WithError(err).Error("Failed to update item")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update item"})
		return
	}

	logrus.WithFields(logrus.Fields{
		"item_id": id,
	}).Info("Item updated successfully")

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}

// @Summary DeleteItem
// @Tags item
// @Description delete item
// @ID delete-item
// @Accept  json
// @Produce  json
// @Param id path int true "Item ID"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/item/{id} [delete]
func (h *Handler) deleteItem(c *gin.Context) {
	logrus.Debug("deleteItem handler called")

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logrus.WithError(err).Error("Invalid item ID")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid item ID"})
		return
	}

	err = h.services.TimeTrackerItem.DeleteItem(id)
	if err != nil {
		logrus.WithError(err).Error("Failed to delete item")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete item"})
		return
	}

	logrus.WithFields(logrus.Fields{
		"item_id": id,
	}).Info("Item deleted successfully")

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}

// @Summary UpdateItemTime
// @Tags item
// @Description update item time
// @ID update-item-time
// @Accept  json
// @Produce  json
// @Param id path int true "Item ID"
// @Param flag path int true "Start(1) or stop(0)"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/item/{id}/time/{flag} [put]
func (h *Handler) updateItemTime(c *gin.Context) {
	logrus.Debug("updateItemTime handler called")

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logrus.WithError(err).Error("Invalid item ID")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid item ID"})
		return
	}

	flag, err := strconv.Atoi(c.Param("flag"))
	if err != nil {
		logrus.WithError(err).Error("Invalid flag value")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid flag value"})
		return
	}

	err = h.services.TimeTrackerItem.UpdateItemTime(id, flag != 0)
	if err != nil {
		logrus.WithError(err).Error("Failed to update item time")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update item time"})
		return
	}

	logrus.WithFields(logrus.Fields{
		"item_id": id,
		"time":    time.Now().Format("2006.01.02 15:04:05"),
	}).Info("Item time updated successfully")

	c.JSON(http.StatusOK, statusResponse{
		Status: time.Now().Format("2006.01.02 15:04:05"),
	})
}

type itemsByDateResponse struct {
	Data []timetracker.GetItemsByDate `json:"data"`
}

// @Summary GetItemsByDate
// @Tags item
// @Description get item by date
// @ID get-item-by-date
// @Accept  json
// @Produce  json
// @Param id path int true "User ID"
// @Param timeStart query string true "Start time" example(2024-07-01T00:00:00Z)
// @Param timeStop query string true "Stop time" example(2024-07-01T59:59:59Z)
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/item/{id}/time [get]
func (h *Handler) getItemsByDate(c *gin.Context) {
	logrus.Debug("getItemsByDate handler called")

	userId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logrus.WithError(err).Error("Invalid user ID")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	var datePeriod timetracker.DatePeriod

	datePeriod.TimeStart, err = time.Parse(time.RFC3339, c.Query("timeStart"))
	if err != nil {
		logrus.WithError(err).Error("Failed to parse timeStart")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid timeStart"})
		return
	}
	datePeriod.TimeStop, err = time.Parse(time.RFC3339, c.Query("timeStop"))
	if err != nil {
		logrus.WithError(err).Error("Failed to parse timeStop")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid timeStop"})
		return
	}

	items, err := h.services.TimeTrackerItem.GetItemsByDate(userId, datePeriod)
	if err != nil {
		logrus.WithError(err).Error("Failed to get items by date")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get items by date"})
		return
	}

	logrus.WithFields(logrus.Fields{
		"user_id":   userId,
		"timeStart": datePeriod.TimeStart,
		"timeStop":  datePeriod.TimeStop,
		"count":     len(items),
	}).Info("Retrieved items by date successfully")

	c.JSON(http.StatusOK, itemsByDateResponse{
		Data: items,
	})
}

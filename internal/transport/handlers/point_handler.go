package handlers

import (
	"net/http"
	"strconv"
	"test-service-for-pick-up-points/internal/models"
	"test-service-for-pick-up-points/internal/services"

	"github.com/gin-gonic/gin"
)

type PointHandler struct {
	pointService *services.PointService
}

func NewPointHandler(pointService *services.PointService) *PointHandler {
	return &PointHandler{pointService: pointService}
}

func (h *PointHandler) CreatePoint(c *gin.Context) {
	var req models.CreatePointRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	point, err := h.pointService.CreatePoint(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, point)
}

func (h *PointHandler) GetPoints(c *gin.Context) {
	points, err := h.pointService.GetAllPoints()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, points)
}

func (h *PointHandler) GetPoint(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid point ID"})
		return
	}

	point, err := h.pointService.GetPoint(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, point)
}

func (h *PointHandler) UpdatePoint(c *gin.Context) {
	userIDStr := c.Param("id")
	userID, err := strconv.ParseUint(userIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid point ID"})
		return
	}

	user, err := h.pointService.GetPoint(uint(userID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	if err := h.pointService.UpdatePoint(user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "point updated successfully"})
}

func (h *PointHandler) DeletePoint(c *gin.Context) {
	userIDStr := c.Param("id")
	userID, err := strconv.ParseUint(userIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid point ID"})
		return
	}

	user, err := h.pointService.GetPoint(uint(userID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	if err := h.pointService.DeletePoint(user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "point deleted successfully"})
}

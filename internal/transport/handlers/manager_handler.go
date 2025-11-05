package handlers

import (
	"net/http"
	"strconv"
	"test-service-for-pick-up-points/internal/models"
	"test-service-for-pick-up-points/internal/services"

	"github.com/gin-gonic/gin"
)

type ManagerHandler struct {
	managerService *services.ManagerService
}

func NewManagerHandler(managerService *services.ManagerService) *ManagerHandler {
	return &ManagerHandler{managerService: managerService}
}

func (h *ManagerHandler) CreateManager(c *gin.Context) {
	var req models.CreateManagerRequests
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	manager, err := h.managerService.CreateManager(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, manager)
}

func (h *ManagerHandler) GetManager(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid manager ID"})
		return
	}

	orders, ordersErr := h.managerService.GetManagerOrders(uint(id))
	if ordersErr != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": ordersErr.Error()})
		return
	}

	point, pointErr := h.managerService.GetManagerPoint(uint(id))
	if pointErr != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": pointErr.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"orders": orders,
		"point":  point,
	})
}

func (h *ManagerHandler) GetManagerOrders(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid manager ID"})
		return
	}

	manager, err := h.managerService.GetManagerOrders(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, manager)
}

func (h *ManagerHandler) GetManagerPoint(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid manager ID"})
		return
	}

	manager, err := h.managerService.GetManagerPoint(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, manager)
}

func (h *ManagerHandler) UpdateManager(c *gin.Context) {
	managerIDStr := c.Param("id")
	managerID, err := strconv.ParseUint(managerIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user ID"})
		return
	}

	manager, err := h.managerService.GetManager(uint(managerID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	if err := h.managerService.UpdateManager(manager); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "manager updated successfully"})
}

func (h *ManagerHandler) DeleteManager(c *gin.Context) {
	managerIDStr := c.Param("id")
	managerID, err := strconv.ParseUint(managerIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user ID"})
		return
	}

	manager, err := h.managerService.GetManager(uint(managerID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	if err := h.managerService.DeleteManager(manager); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "manager deleted successfully"})
}

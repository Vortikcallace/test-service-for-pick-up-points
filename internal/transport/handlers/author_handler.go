package handlers

import (
	"net/http"
	"strconv"
	"test-service-for-pick-up-points/internal/models"
	"test-service-for-pick-up-points/internal/services"

	"github.com/gin-gonic/gin"
)

type AuthorHandler struct {
	authorService *services.AuthorService
}

func NewAuthorHandler(authorService *services.AuthorService) *AuthorHandler {
	return &AuthorHandler{authorService: authorService}
}

func (h *AuthorHandler) CreateAuthor(c *gin.Context) {
	var req models.CreateAuthorRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	author, err := h.authorService.CreateAuthor(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, author)
}

func (h *AuthorHandler) GetAuthor(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid author ID"})
		return
	}

	author, err := h.authorService.GetAuthorProducts(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, author)
}

func (h *AuthorHandler) UpdateAuthor(c *gin.Context) {
	authorIDStr := c.Param("id")
	authorID, err := strconv.ParseUint(authorIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid author ID"})
		return
	}

	author, err := h.authorService.GetAuthor(uint(authorID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	if err := h.authorService.UpdateAuthor(author); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "author updated successfully"})
}

func (h *AuthorHandler) DeleteAuthor(c *gin.Context) {
	authorIDStr := c.Param("id")
	authorID, err := strconv.ParseUint(authorIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid author ID"})
		return
	}

	author, err := h.authorService.GetAuthor(uint(authorID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	if err := h.authorService.DeleteAuthor(author); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "author deleted successfully"})
}

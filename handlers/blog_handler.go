package handlers

import (
	"alexnerotd/blog/models"
	"alexnerotd/blog/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BlogHandler struct {
	blogService *services.BlogService
}

func NewBlogHandler(blogService *services.BlogService) *BlogHandler {
	return &BlogHandler{blogService: blogService}
}

func (h *BlogHandler) GetBlogById(c *gin.Context) {
	blogID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid blog ID"})
		return
	}
	blog, err := h.blogService.GetBlogById(blogID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Blog not found"})
		return
	}

	c.JSON(http.StatusOK, blog)
}

func (h *BlogHandler) CreateBlog(c *gin.Context) {
	var blog models.Blog
	if err := c.BindJSON(&blog); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.blogService.CreateBlog(&blog); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failde to create blog"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": blog})
}

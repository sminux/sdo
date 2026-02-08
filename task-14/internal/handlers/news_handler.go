package handlers

import (
	"net/http"
	"task-14/internal/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type NewsHandler struct {
	db *gorm.DB
}

func NewNewsHandler(db *gorm.DB) *NewsHandler {
	return &NewsHandler{db: db}
}

func (h *NewsHandler) GetNews(c *gin.Context) {
	var news []models.News

	result := h.db.Where("is_active = ?", true).Order("published_at DESC").Find(&news)
	if result.Error != nil {
		c.HTML(http.StatusInternalServerError, "news.html", gin.H{
			"title": "Новости - Ошибка",
			"error": "Не удалось загрузить новости",
		})
		return
	}

	c.HTML(http.StatusOK, "news.html", gin.H{
		"title": "Новости",
		"news":  news,
	})
}

func (h *NewsHandler) GetNewsItem(c *gin.Context) {
	id := c.Param("id")
	var newsItem models.News

	result := h.db.Where("id = ? AND is_active = ?", id, true).First(&newsItem)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			c.HTML(http.StatusNotFound, "news.html", gin.H{
				"title": "Новость не найдена",
				"error": "Запрошенная новость не найдена",
			})
		} else {
			c.HTML(http.StatusInternalServerError, "news.html", gin.H{
				"title": "Ошибка",
				"error": "Не удалось загрузить новость",
			})
		}
		return
	}

	c.HTML(http.StatusOK, "news.html", gin.H{
		"title":    newsItem.Title,
		"newsItem": newsItem,
		"single":   true,
	})
}

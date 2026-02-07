package handlers

import (
	"net/http"
	"task-14/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CasesHandler struct {
	db *gorm.DB
}

func NewCasesHandler(db *gorm.DB) *CasesHandler {
	return &CasesHandler{db: db}
}

func (h *CasesHandler) GetCases(c *gin.Context) {
	var cases []models.Case

	result := h.db.Where("is_published = ?", true).Order("implemented_at DESC").Find(&cases)
	if result.Error != nil {
		c.HTML(http.StatusInternalServerError, "cases.html", gin.H{
			"title": "Кейсы - Ошибка",
			"error": "Не удалось загрузить кейсы",
		})
		return
	}

	c.HTML(http.StatusOK, "cases.html", gin.H{
		"title": "Кейсы",
		"cases": cases,
	})
}

func (h *CasesHandler) GetCase(c *gin.Context) {
	id := c.Param("id")
	var caseItem models.Case

	result := h.db.Where("id = ? AND is_published = ?", id, true).First(&caseItem)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			c.HTML(http.StatusNotFound, "cases.html", gin.H{
				"title": "Кейс не найден",
				"error": "Запрошенный кейс не найден",
			})
		} else {
			c.HTML(http.StatusInternalServerError, "cases.html", gin.H{
				"title": "Ошибка",
				"error": "Не удалось загрузить кейс",
			})
		}
		return
	}

	c.HTML(http.StatusOK, "cases.html", gin.H{
		"title":    caseItem.Title,
		"caseItem": caseItem,
		"single":   true,
	})
}

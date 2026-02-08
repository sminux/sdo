package handlers

import (
	"net/http"
	"task-14/internal/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type MainHandler struct {
	db *gorm.DB
}

type ContactForm struct {
	Organization string `form:"organization" binding:"required"`
	LastName     string `form:"lastName" binding:"required"`
	FirstName    string `form:"firstName" binding:"required"`
	Phone        string `form:"phone" binding:"required"`
	Email        string `form:"email" binding:"required,email"`
	Consent      string `form:"consent" binding:"required"`
}

func NewMainHandler(db *gorm.DB) *MainHandler {
	return &MainHandler{db: db}
}

func (h *MainHandler) MainPage(c *gin.Context) {
	c.HTML(http.StatusOK, "main.html", gin.H{
		"title": "Главная страница",
	})
}

func (h *MainHandler) SubmitContact(c *gin.Context) {
	var form ContactForm

	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Пожалуйста, заполните все поля правильно",
			"errors":  err.Error(),
		})
		return
	}

	// Проверка согласия
	if form.Consent != "on" {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Необходимо согласие на обработку персональных данных",
		})
		return
	}

	// SubmittedContact, который будет сохранен в БД
	contact := models.SubmittedContact{
		Organization: form.Organization,
		LastName:     form.LastName,
		FirstName:    form.FirstName,
		Phone:        form.Phone,
		Email:        form.Email,
		Consent:      form.Consent == "on",
		Processed:    false,
	}

	// Сохранение в БД
	result := h.db.Create(&contact)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Ошибка при сохранении данных. Попробуйте позже.",
			"errors":  result.Error.Error(),
		})
		return
	}

	// Отправка по почте
	go func() {
		// TODO: postfix, snmp
		_ = contact.ID
	}()

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Ваш запрос успешно отправлен! Мы свяжемся с вами в ближайшее время.",
	})
}

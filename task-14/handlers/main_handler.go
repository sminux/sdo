package handlers

import (
	"net/http"

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
	Consent      bool   `form:"consent" binding:"required"`
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
			"message": "Проверьте правильность заполнения полей",
			"errors":  err.Error(),
		})
		return
	}

	// проверка на согласие персональных данных
	if !form.Consent {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Необходимо согласие на обработку персональных данных",
		})
		return
	}

	// отправка письма
	go func() {
		// TODO: postfix или net/smtp
		_ = form
	}()

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Ваш запрос успешно отправлен! Мы свяжемся с вами в ближайшее время.",
	})
}

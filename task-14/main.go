package main

import (
	"log"
	"net/http"
	"task-14/config"
	"task-14/handlers"
	"task-14/models"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Загрузка конфигурации
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Ошибка загрузки конфигурации:", err)
	}

	// Подключение к базе данных
	dsn := cfg.Database.GetDSN()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Ошибка подключения к базе данных:", err)
	}

	// Автомиграция моделей
	err = db.AutoMigrate(&models.News{}, &models.Case{})
	if err != nil {
		log.Fatal("Ошибка миграции базы данных:", err)
	}

	// Инициализация Gin (как приятное дополнение к роутингу)
	router := gin.Default()

	// Загрузка шаблонов
	// router.LoadHTMLGlob(cfg.Paths.Templates + "/*.html")
	router.LoadHTMLGlob("pages/*.html")

	// Статические файлы
	router.Static("/static", cfg.Paths.Static)

	mainHandler := handlers.NewMainHandler(db)
	newsHandler := handlers.NewNewsHandler(db)
	casesHandler := handlers.NewCasesHandler(db)
	docsHandler := handlers.NewDocsHandler(cfg.Paths.Docs)

	// Middleware для установки текущей страницы
	router.Use(func(c *gin.Context) {
		currentPage := ""
		path := c.Request.URL.Path

		switch {
		case path == "/":
			currentPage = "home"
		case path == "/news" || len(path) > 6 && path[:6] == "/news/":
			currentPage = "news"
		case path == "/cases" || len(path) > 7 && path[:7] == "/cases/":
			currentPage = "cases"
		case path == "/docs" || len(path) > 6 && path[:6] == "/docs/":
			currentPage = "docs"
		}

		c.Set("current", currentPage)
		c.Next()
	})

	router.GET("/", mainHandler.MainPage)
	router.POST("/contact", mainHandler.SubmitContact)

	router.GET("/news", newsHandler.GetNews)
	router.GET("/news/:id", newsHandler.GetNewsItem)

	router.GET("/cases", casesHandler.GetCases)
	router.GET("/cases/:id", casesHandler.GetCase)

	router.GET("/docs", docsHandler.ListDocs)
	router.GET("/docs/:filename", docsHandler.GetDoc)

	srv := &http.Server{
		Addr:         cfg.Server.Port,
		Handler:      router,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	log.Printf("Сервер запущен на http://%s%s", cfg.Server.Host, cfg.Server.Port)
	log.Printf("База данных: %s", dsn)

	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal("Ошибка запуска сервера:", err)
	}
}

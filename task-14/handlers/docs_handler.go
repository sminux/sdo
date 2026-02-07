package handlers

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type DocsHandler struct {
	docsPath string
}

type DocFile struct {
	Name          string    `json:"name"`
	DisplayName   string    `json:"display_name"`
	Size          int64     `json:"size"`
	SizeFormatted string    `json:"size_formatted"`
	ModifiedTime  time.Time `json:"modified_time"`
	Extension     string    `json:"extension"`
}

func NewDocsHandler(docsPath string) *DocsHandler {
	return &DocsHandler{docsPath: docsPath}
}

func (h *DocsHandler) ListDocs(c *gin.Context) {
	files, err := h.getDocFiles()
	if err != nil {
		c.HTML(http.StatusInternalServerError, "docs.html", gin.H{
			"title": "Документация - Ошибка",
			"error": "Не удалось загрузить документы",
		})
		return
	}

	c.HTML(http.StatusOK, "docs.html", gin.H{
		"title": "Документация",
		"docs":  files,
	})
}

func (h *DocsHandler) GetDoc(c *gin.Context) {
	filename := c.Param("filename")

	// обработка path traversal (может быть уязвимо)
	if strings.Contains(filename, "..") || strings.Contains(filename, "/") {
		c.HTML(http.StatusForbidden, "docs.html", gin.H{
			"title": "Доступ запрещен",
			"error": "Недопустимое имя файла",
		})
		return
	}

	filePath := filepath.Join(h.docsPath, filename)

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		c.HTML(http.StatusNotFound, "docs.html", gin.H{
			"title": "Файл не найден",
			"error": "Запрошенный файл не найден",
		})
		return
	}

	c.File(filePath)
}

func (h *DocsHandler) getDocFiles() ([]DocFile, error) {
	var files []DocFile

	entries, err := os.ReadDir(h.docsPath)
	if err != nil {
		return nil, err
	}

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		info, err := entry.Info()
		if err != nil {
			continue
		}

		filename := entry.Name()
		ext := strings.ToLower(filepath.Ext(filename))

		// используем известные типы документов
		if !isDocumentType(ext) {
			continue
		}

		files = append(files, DocFile{
			Name:          filename,
			DisplayName:   h.getDisplayName(filename),
			Size:          info.Size(),
			SizeFormatted: h.formatFileSize(info.Size()),
			ModifiedTime:  info.ModTime(),
			Extension:     ext,
		})
	}

	return files, nil
}

func (h *DocsHandler) getDisplayName(filename string) string {
	// Без расширения
	name := filename[:len(filename)-len(filepath.Ext(filename))]

	// Без лишних спецсимволов
	name = strings.ReplaceAll(name, "_", " ")
	name = strings.ReplaceAll(name, "-", " ")

	// Заглавная первая буква
	if len(name) > 0 {
		name = strings.ToUpper(name[:1]) + name[1:]
	}

	return name
}

func (h *DocsHandler) formatFileSize(bytes int64) string {
	if bytes == 0 {
		return "0 Bytes"
	}

	const unit = 1024
	if bytes < unit {
		return fmt.Sprintf("%d B", bytes)
	}
	div, exp := int64(unit), 0
	for n := bytes / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}

	units := []string{"KB", "MB", "GB", "TB"}
	return fmt.Sprintf("%.1f %s", float64(bytes)/float64(div), units[exp])
}

func isDocumentType(ext string) bool {
	documentTypes := []string{
		".pdf", ".doc", ".docx", ".txt", ".rtf",
		".xls", ".xlsx", ".ppt", ".pptx",
	}
	for _, docType := range documentTypes {
		if ext == docType {
			return true
		}
	}
	return false
}

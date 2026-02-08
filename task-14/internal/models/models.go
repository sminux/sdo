package models

import (
	"time"

	"gorm.io/gorm"
)

type News struct {
	gorm.Model
	Title       string    `gorm:"size:255;not null" json:"title"`
	Content     string    `gorm:"type:text;not null" json:"content"`
	ImageURL    string    `gorm:"size:500" json:"image_url"`
	PublishedAt time.Time `json:"published_at"`
	Author      string    `gorm:"size:100" json:"author"`
	IsActive    bool      `gorm:"default:true" json:"is_active"`
}

type Case struct {
	gorm.Model
	Organization  string    `gorm:"size:255;not null" json:"organization"`
	Title         string    `gorm:"size:255;not null" json:"title"`
	Description   string    `gorm:"type:text;not null" json:"description"`
	LogoURL       string    `gorm:"size:500" json:"logo_url"`
	Content       string    `gorm:"type:text" json:"content"`
	Industry      string    `gorm:"size:100" json:"industry"`
	Employees     int       `json:"employees"`
	ImplementedAt time.Time `json:"implemented_at"`
	IsPublished   bool      `gorm:"default:true" json:"is_published"`
}

type SubmittedContact struct {
	gorm.Model
	Organization string    `gorm:"size:255;not null" json:"organization"`
	LastName     string    `gorm:"size:100;not null" json:"last_name"`
	FirstName    string    `gorm:"size:100;not null" json:"first_name"`
	Phone        string    `gorm:"size:20;not null" json:"phone"`
	Email        string    `gorm:"size:255;not null" json:"email"`
	Consent      bool      `gorm:"not null" json:"consent"`
	Processed    bool      `gorm:"default:false" json:"processed"`
	SubmittedAt  time.Time `gorm:"autoCreateTime" json:"submitted_at"`
}

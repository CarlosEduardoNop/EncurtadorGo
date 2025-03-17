package model

import "time"

type Link struct {
	ID           string `gorm:"primaryKey;size:64;"`
	OriginalLink string `gorm:"not null;size:50"`
	CreatedAt    *time.Time
	UpdatedAt    *time.Time
}

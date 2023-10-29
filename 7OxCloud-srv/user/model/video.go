package model

import (
	"time"

	"gorm.io/gorm"
)

type Video struct {
	ID        int64 `gorm:"primaryKey;index"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	AuthorId    int    `gorm:"index"`
	Author      User   `gorm:"foreignKey:AuthorId;AssociationForeignKey:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	PlayUrl     string `gorm:"not null"`
	CoverUrl    string `gorm:"not null"`
	Description string `gorm:"null;size:2048"`
}

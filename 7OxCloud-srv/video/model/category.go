package model

import "time"

type Category struct {
	ID        int64 `gorm:"primaryKey;index"`
	CreatedAt time.Time

	Name string `gorm:"not null;size:64"`
}

type VideoCategory struct {
	ID        int64 `gorm:"primaryKey;index"`
	CreatedAt time.Time

	VideoId    int64    `gorm:"index"`
	Video      Video    `gorm:"foreignKey:VideoId;AssociationForeignKey:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	CategoryId int64    `gorm:"index"`
	Category   Category `gorm:"foreignKey:CategoryId;AssociationForeignKey:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

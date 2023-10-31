package model

import "time"

type Topic struct {
	ID        int64 `gorm:"primaryKey"`
	CreatedAt time.Time

	Description string `gorm:"not null"`
}

type VideoTopic struct {
	ID        int64 `gorm:"primaryKey"`
	CreatedAt time.Time

	VideoId int64 `gorm:"index"`
	Video   Video `gorm:"foreignKey:VideoId;AssociationForeignKey:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	TopicId int64 `gorm:"index"`
	Topic   Topic `gorm:"foreignKey:TopicId;AssociationForeignKey:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

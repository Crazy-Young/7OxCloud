package model

import "time"

type Comment struct {
	ID        int64 `gorm:"primaryKey"`
	CreatedAt time.Time

	VideoId  int64    `gorm:"index"`
	Video    Video    `gorm:"foreignKey:VideoId;AssociationForeignKey:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	UserId   int64    `gorm:"index"`
	User     User     `gorm:"foreignKey:UserId;AssociationForeignKey:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	ParentId *int64   `gorm:"index;default:NULL"`
	Parent   *Comment `gorm:"foreignKey:ParentId;AssociationForeignKey:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Content  string   `gorm:"not null;size:1024"`
}

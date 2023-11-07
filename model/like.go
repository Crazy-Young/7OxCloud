package model

import "time"

type Like struct {
	ID        int64     `gorm:"primaryKey"`
	CreatedAt time.Time `gorm:"index"`

	UserId  int   `gorm:"index:idx_like_user;index:idx_like"`
	User    User  `gorm:"foreignKey:UserId;AssociationForeignKey:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	VideoId int   `gorm:"index:idx_like_video;index:idx_like"`
	Video   Video `gorm:"foreignKey:VideoId;AssociationForeignKey:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

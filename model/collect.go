package model

import "time"

type Collect struct {
	ID        int64 `gorm:"primaryKey"`
	CreatedAt time.Time

	UserId  int   `gorm:"index:idx_collect_user;index:idx_collect"`
	User    User  `gorm:"foreignKey:UserId;AssociationForeignKey:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	VideoId int   `gorm:"index:idx_collect_video;index:idx_collect"`
	Video   Video `gorm:"foreignKey:VideoId;AssociationForeignKey:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

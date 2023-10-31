package model

import "time"

type Follow struct {
	ID        int64 `gorm:"primaryKey"`
	CreatedAt time.Time

	UserId int64 `gorm:"index:idx_follow_user;index:idx_follow"`
	User   User  `gorm:"foreignKey:UserId;AssociationForeignKey:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	FanId  int64 `gorm:"index:idx_follow_fan;index:idx_follow"`
	Fan    User  `gorm:"foreignKey:FanId;AssociationForeignKey:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

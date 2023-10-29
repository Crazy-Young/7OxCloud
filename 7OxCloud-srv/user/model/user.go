package model

import (
	"time"
)

type User struct {
	ID        int64 `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time

	Username        string `gorm:"not null"`
	Age             int    `gorm:"null"`
	Location        string `gorm:"null"`
	Password        string `gorm:"not null"`
	Mobile          string `gorm:"not null"`
	Avatar          string `gorm:"not null;default:http://s38ddu7np.hn-bkt.clouddn.com/avatar.jpg"`
	BackgroundImage string `gorm:"not null;default:http://s38ddu7np.hn-bkt.clouddn.com/background.png"`
	Signature       string `gorm:"not null;default:这个人很懒，什么都没有留下"`
}

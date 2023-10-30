package dao

import (
	"github.com/palp1tate/7OxCloud/7OxCloud-srv/user/global"
	"github.com/palp1tate/7OxCloud/7OxCloud-srv/user/model"
	"gorm.io/gorm"
)

func FindUserByMobile(mobile string) (user model.User, err error) {
	err = global.DB.Where("mobile = ?", mobile).First(&user).Error
	return
}

func CreateUser(user *model.User) (err error) {
	err = global.DB.Create(&user).Error
	return
}

func FindUserById(uid int64) (user model.User, err error) {
	err = global.DB.Where("id = ?", uid).First(&user).Error
	return
}

func GetFollowCount(tx *gorm.DB, uid int64) (count int64, err error) {
	err = tx.Model(&model.Follow{}).Where("fan_id = ?", uid).Count(&count).Error
	return
}

func GetFanCount(tx *gorm.DB, uid int64) (count int64, err error) {
	err = tx.Model(&model.Follow{}).Where("user_id = ?", uid).Count(&count).Error
	return
}

func GetIsFollow(tx *gorm.DB, uid int64, currentUserId int64) (isFollow bool, err error) {
	var follow model.Follow
	isFollow = true
	err = tx.Where("user_id = ? AND fan_id = ?", uid, currentUserId).First(&follow).Error
	if err != nil {
		isFollow = false
		return
	}
	return
}

func GetLikeCount(tx *gorm.DB, uid int64) (count int64, err error) {
	err = tx.Model(&model.Like{}).Where("user_id = ?", uid).Count(&count).Error
	return
}

func GetLikedCount(tx *gorm.DB, uid int64) (count int64, err error) {
	err = tx.Model(&model.User{}).Where("user.id = ?", uid).
		Joins("join video ON user.id = video.author_id").
		Joins("join `like` ON video.id = `like`.video_id").
		Count(&count).Error
	return
}

func GetWorkCount(tx *gorm.DB, uid int64) (count int64, err error) {
	err = tx.Model(&model.Video{}).Where("author_id = ?", uid).Count(&count).Error
	return
}

func GetCollectCount(tx *gorm.DB, uid int64) (count int64, err error) {
	err = tx.Model(&model.Collect{}).Where("user_id = ?", uid).Count(&count).Error
	return
}

func UpdateUser(user *model.User) (err error) {
	err = global.DB.Save(&user).Error
	return
}

func UpdateUserPassword(user *model.User, password string) (err error) {
	err = global.DB.Model(&user).Update("password", password).Error
	return
}

package dao

import (
	"time"

	"gorm.io/gorm"

	"github.com/palp1tate/7OxCloud/7OxCloud-srv/social/global"
	"github.com/palp1tate/7OxCloud/model"
)

func Follow(userId int64, fanId int64) (err error) {
	follow := model.Follow{
		UserId: userId,
		FanId:  fanId,
	}
	err = global.DB.FirstOrCreate(&follow, follow).Error
	return
}

func CancelFollow(userId int64, fanId int64) (err error) {
	err = global.DB.Where("user_id = ? AND fan_id = ?", userId, fanId).Delete(&model.Follow{}).Error
	return
}

func GetFollowing(latestTime time.Time, userId int64) (followingList []model.User, count int64, err error) {
	tx := global.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	err = tx.Model(&model.Follow{}).Where("fan_id = ?", userId).Count(&count).Error
	if err != nil {
		tx.Rollback()
		return
	}
	err = tx.Joins("left join follow on follow.user_id = user.id").
		Where("follow.fan_id = ? AND follow.created_at < ?", userId, latestTime).
		Order("follow.id desc").
		Limit(15).
		Find(&followingList).Error
	if err != nil {
		tx.Rollback()
		return
	}
	tx.Commit()
	return
}

func GetFan(latestTime time.Time, userId int64) (fanList []model.User, count int64, err error) {
	tx := global.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	err = tx.Model(&model.Follow{}).Where("user_id = ?", userId).Count(&count).Error
	if err != nil {
		tx.Rollback()
		return
	}
	err = tx.Joins("left join follow on follow.fan_id = user.id").
		Where("follow.user_id = ? AND follow.created_at < ?", userId, latestTime).
		Order("follow.id desc").
		Limit(15).
		Find(&fanList).Error
	if err != nil {
		tx.Rollback()
		return
	}
	tx.Commit()
	return
}

func GetFollowFeed(latestTime time.Time, userId int64) (videoList []*model.Video, err error) {
	err = global.DB.Joins("left join follow on follow.user_id = video.author_id").
		Where("follow.fan_id = ? AND video.created_at < ?", userId, latestTime).
		Order("video.id desc").
		Limit(15).Find(&videoList).Error
	return
}

func GetFriendFeed(latestTime time.Time, userId int64) (videoList []*model.Video, err error) {
	var mutualFriendIDs []int64
	tx := global.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	err = tx.Model(&model.Follow{}).
		Joins("JOIN follow AS f2 ON follow.fan_id = f2.user_id AND follow.user_id = f2.fan_id").
		Where("follow.fan_id = ?", userId).
		Pluck("f2.fan_id", &mutualFriendIDs).Error
	if err != nil {
		tx.Rollback()
		return
	}
	mutualFriendIDs = append(mutualFriendIDs, userId)
	err = tx.Where("author_id IN (?) AND created_at < ?", mutualFriendIDs, latestTime).
		Order("id desc").
		Limit(15).
		Find(&videoList).Error
	if err != nil {
		tx.Rollback()
		return
	}
	tx.Commit()
	return
}

func GetVideoLikeCount(tx *gorm.DB, vid int64) (count int64, err error) {
	err = tx.Model(&model.Like{}).Where("video_id = ?", vid).Count(&count).Error
	return
}

func GetVideoCommentCount(tx *gorm.DB, vid int64) (count int64, err error) {
	err = tx.Model(&model.Comment{}).Where("video_id = ?", vid).Count(&count).Error
	return
}

func GetVideoCollectCount(tx *gorm.DB, vid int64) (count int64, err error) {
	err = tx.Model(&model.Collect{}).Where("video_id = ?", vid).Count(&count).Error
	return
}

func GetIsLike(tx *gorm.DB, vid int64, uid int64) (isLike bool, err error) {
	var like model.Like
	isLike = true
	err = tx.Where("video_id = ? and user_id = ?", vid, uid).First(&like).Error
	if err != nil {
		isLike = false
	}
	return
}

func GetIsCollect(tx *gorm.DB, vid int64, uid int64) (isCollect bool, err error) {
	var collect model.Collect
	isCollect = true
	err = tx.Where("video_id = ? and user_id = ?", vid, uid).First(&collect).Error
	if err != nil {
		isCollect = false
	}
	return
}

func FindAuthor(tx *gorm.DB, vid int64) (author model.User, err error) {
	err = tx.Joins("join video on video.author_id = user.id").
		Where("video.id = ?", vid).
		First(&author).Error
	return
}

func GetVideoTopics(tx *gorm.DB, vid int64) (topicList []*model.Topic, err error) {
	err = tx.Model(&model.Topic{}).
		Joins("join video_topic on topic.id = video_topic.topic_id").
		Where("video_topic.video_id = ?", vid).
		Order("topic.id desc").
		Find(&topicList).Error
	return
}

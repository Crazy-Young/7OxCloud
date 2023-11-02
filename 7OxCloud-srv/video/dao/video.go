package dao

import (
	"time"

	"github.com/palp1tate/7OxCloud/7OxCloud-srv/video/global"
	"github.com/palp1tate/7OxCloud/model"
	"gorm.io/gorm"
)

func GetVideos(latestTime time.Time) (videoList []*model.Video, err error) {
	err = global.DB.Where("created_at < ?", latestTime).
		Order("id desc").
		Limit(15).
		Find(&videoList).Error
	return
}

func GetHotVideos(latestTime time.Time) (videoList []*model.Video, err error) {
	err = global.DB.Model(&model.Video{}).
		Joins("left join `like` on video.id = `like`.video_id").
		Where("video.created_at < ?", latestTime).
		Group("video.id").
		Order("count(`like`.id) desc").
		Order("video.id desc").
		Limit(15).
		Find(&videoList).Error
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

func GetVideoTopics(tx *gorm.DB, vid int64) (topicList []*model.Topic, err error) {
	err = tx.Model(&model.Topic{}).
		Joins("join video_topic on topic.id = video_topic.topic_id").
		Where("video_topic.video_id = ?", vid).
		Order("topic.id desc").
		Find(&topicList).Error
	return
}

func GetVideosByUserId(latestTime time.Time, uid int64) (videoList []*model.Video, err error) {
	err = global.DB.Where("author_id = ? AND created_at < ?", uid, latestTime).
		Order("id desc").
		Limit(15).
		Find(&videoList).Error
	return
}

func GetVideo(vid int64) (video *model.Video, err error) {
	video = &model.Video{}
	err = global.DB.Where("id = ?", vid).First(&video).Error
	return
}

func CreateVideo(video *model.Video, topics []string, categoryId int64) (err error) {
	tx := global.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err = tx.Create(&video).Error; err != nil {
		tx.Rollback()
		return
	}
	for _, topic := range topics {
		topicModel := model.Topic{
			Description: topic,
		}
		if err = tx.FirstOrCreate(&topicModel, topicModel).Error; err != nil {
			tx.Rollback()
			return
		}
		videoTopic := model.VideoTopic{
			VideoId: video.ID,
			TopicId: topicModel.ID,
		}
		if err = tx.Create(&videoTopic).Error; err != nil {
			tx.Rollback()
			return
		}
	}
	videoCategory := model.VideoCategory{
		VideoId:    video.ID,
		CategoryId: categoryId,
	}
	if err = tx.Create(&videoCategory).Error; err != nil {
		tx.Rollback()
		return
	}
	return tx.Commit().Error
}

func GetVideosByTopicId(latestTime time.Time, topicId int64) (videoList []*model.Video, err error) {
	err = global.DB.Model(&model.Video{}).
		Joins("join video_topic on video.id = video_topic.video_id").
		Where("video_topic.topic_id = ? AND video.created_at < ?", topicId, latestTime).
		Order("video.id desc").
		Limit(15).
		Find(&videoList).Error
	return
}

func GetCategories() (categoryList []*model.Category, err error) {
	err = global.DB.Order("id desc").Find(&categoryList).Error
	return
}

func GetTopics() (topicList []*model.Topic, err error) {
	err = global.DB.Order("id desc").Find(&topicList).Error
	return
}

func GetVideosByCategoryId(latestTime time.Time, categoryId int64) (videoList []*model.Video, err error) {
	err = global.DB.Model(&model.Video{}).
		Joins("join video_category on video.id = video_category.video_id").
		Where("video_category.category_id = ? AND video.created_at < ?", categoryId, latestTime).
		Order("video.id desc").
		Limit(15).
		Find(&videoList).Error
	return
}

func GetVideosBySearch(latestTime time.Time, keyword string) (videoList []*model.Video, err error) {
	err = global.DB.Where("description like ? AND created_at < ?", "%"+keyword+"%", latestTime).
		Order("id desc").
		Limit(15).Find(&videoList).Error
	return
}

func DeleteVideo(vid int64, uid int64) (err error) {
	err = global.DB.Where("id = ? AND author_id = ?", vid, uid).Delete(&model.Video{}).Error
	return
}

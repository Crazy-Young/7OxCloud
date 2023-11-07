package dao

import (
	"time"

	"github.com/palp1tate/7OxCloud/7OxCloud-srv/interaction/global"
	"github.com/palp1tate/7OxCloud/model"
	interactionProto "github.com/palp1tate/7OxCloud/proto/interaction"
	"gorm.io/gorm"
)

func LikeVideo(uid int64, vid int64) (err error) {
	like := model.Like{
		UserId:  int(uid),
		VideoId: int(vid),
	}
	err = global.DB.FirstOrCreate(&like, like).Error
	return
}

func CancelLikeVideo(uid int64, vid int64) (err error) {
	err = global.DB.Where("user_id=? and video_id=?", uid, vid).Delete(&model.Like{}).Error
	return
}

func CollectVideo(uid int64, vid int64) (err error) {
	collect := model.Collect{
		UserId:  int(uid),
		VideoId: int(vid),
	}
	err = global.DB.FirstOrCreate(&collect, collect).Error
	return
}

func CancelCollectVideo(uid int64, vid int64) (err error) {
	err = global.DB.Where("user_id=? and video_id=?", uid, vid).Delete(&model.Collect{}).Error
	return
}

func GetUserLikedVideos(latestTime time.Time, uid int64) (videos []*model.Video, err error) {
	err = global.DB.
		Joins("join `like` on video.id=`like`.video_id").
		Where("`like`.created_at < ? AND `like`.user_id=?", latestTime, uid).
		Order("`like`.id desc").
		Limit(15).
		Find(&videos).Error
	return
}

func GetVideoLikeCount(tx *gorm.DB, vid int64) (count int64, err error) {
	err = tx.Model(&model.Like{}).Where("video_id = ?", vid).Count(&count).Error
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

func GetUserCollectedVideos(latestTime time.Time, uid int64) (videos []*model.Video, err error) {
	err = global.DB.
		Joins("join collect on video.id=collect.video_id").
		Where("collect.created_at < ? AND collect.user_id=?", latestTime, uid).
		Order("collect.id desc").
		Limit(15).
		Find(&videos).Error
	return
}

func CreateComment(comment *model.Comment) (err error) {
	err = global.DB.Create(&comment).Error
	return
}

func DeleteComment(uid int64, cid int64) (err error) {
	err = global.DB.Where("user_id=? and id=?", uid, cid).Delete(&model.Comment{}).Error
	return
}

func GetCommentList(vid int64) (commentList []*interactionProto.Comment, err error) {
	tx := global.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	var comments []model.Comment
	err = tx.Where("video_id = ? AND parent_id IS NULL", vid).Order("id desc").Find(&comments).Error
	if err != nil {
		return nil, err
	}
	for _, comment := range comments {
		protoComment, err := convertToProto(tx, comment)
		if err != nil {
			return nil, err
		}
		commentList = append(commentList, protoComment)
	}
	err = tx.Commit().Error
	return
}

func convertToProto(tx *gorm.DB, comment model.Comment) (*interactionProto.Comment, error) {
	childrenProto, err := GetChildren(tx, comment.ID)
	if err != nil {
		return nil, err
	}
	var author model.User
	tx.Where("id = ?", comment.UserId).First(&author)
	return &interactionProto.Comment{
		Cid:     comment.ID,
		Content: comment.Content,
		Author: &interactionProto.User{
			Id:       author.ID,
			Username: author.Username,
			Avatar:   author.Avatar,
			Location: author.Location,
		},
		CreatedAt: comment.CreatedAt.Unix(),
		Children:  childrenProto,
	}, nil
}

func GetChildren(tx *gorm.DB, parentId int64) ([]*interactionProto.Comment, error) {
	var comments []model.Comment
	err := tx.Where("parent_id = ?", parentId).Find(&comments).Error
	if err != nil {
		return nil, err
	}
	childrenProto := make([]*interactionProto.Comment, len(comments))
	for i, comment := range comments {
		childrenProto[i], err = convertToProto(tx, comment)
		if err != nil {
			return nil, err
		}
	}
	return childrenProto, nil
}

package form

type LikeVideoForm struct {
	VideoId string `form:"vid" json:"vid" binding:"required"`
	Type    int    `form:"type" json:"type" binding:"required,oneof=1 2"` // 1表示点赞, 2表示取消点赞
}

type CollectVideoForm struct {
	VideoId string `form:"vid" json:"vid" binding:"required"`
	Type    int    `form:"type" json:"type" binding:"required,oneof=1 2"` // 1表示收藏, 2表示取消收藏
}

type CommentForm struct {
	VideoId  string `form:"vid" json:"vid" binding:"required"`
	Content  string `form:"content" json:"content" binding:"required,max=1024"`
	ParentId int    `form:"pid" json:"pid" binding:"gte=0"`
}

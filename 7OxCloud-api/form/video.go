package form

type PublishForm struct {
	Description string   `form:"description" json:"description" binding:"required"`
	PlayUrl     string   `form:"playUrl" json:"playUrl" binding:"required"`
	CoverUrl    string   `form:"coverUrl" json:"coverUrl" binding:"required"`
	CategoryId  int      `form:"categoryId" json:"categoryId" binding:"required"`
	Topics      []string `form:"topics" json:"topics" binding:"omitempty,dive"`
}

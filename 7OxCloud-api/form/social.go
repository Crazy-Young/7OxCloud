package form

type FollowForm struct {
	UserId int `form:"userId" json:"userId" binding:"required"`
	Type   int `form:"type" json:"type" binding:"required,oneof=1 2"` // 1: follow, 2: cancel follow
}

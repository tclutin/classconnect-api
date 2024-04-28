package group

type CreateGroupRequest struct {
	Name string `json:"name" binding:"required,min=3"`
}

type JoinToGroupRequest struct {
	Code string `json:"code" binding:"required,min=4,max=4"`
}

package interfaces

type InviteRequest struct {
	ReceiverId int `json:"receiverId" example:"1"`
	SenderId   int `json:"senderId" example:"2"`
}

type InviteService interface {
	InviteFriend(InviteRequest) (interface{}, error)
}

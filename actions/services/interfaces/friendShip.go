package interfaces

type FriendRequest struct {
	ReceiverId int `json:"receiverId" example:"1"`
	SenderId   int `json:"senderId" example:"2"`
}

type FriendService interface {
	//InviteFriend(InviteRequest) (interface{}, error)
	//CheckFriendInvite(int) (interface{}, error)
	//RejectInvitation(InviteRequest) (interface{}, error)
	//AcceptInvitation(InviteRequest) (interface{}, error)
}

package interfaces

type InviteRequest struct {
	ReceiverEmail string `json:"targetMailAddress" example:"winner@mail.com"`
}

type InviteService interface {
	InviteFriend(string, InviteRequest) (interface{}, error)
	CheckFriendInvite(int) (interface{}, error)
	RejectInvitation(InviteRequest) (interface{}, error)
	AcceptInvitation(InviteRequest) (interface{}, error)
}

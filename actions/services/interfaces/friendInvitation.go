package interfaces

type InviteRequest struct {
	ReceiverEmail string `json:"targetMailAddress" example:"winner@mail.com"`
}

type InviteService interface {
	InviteFriend(string, InviteRequest) (interface{}, error)
	CheckFriendInvite(int) (interface{}, error)
	RejectInvitation(string, int) (interface{}, error)
	AcceptInvitation(InviteRequest) (interface{}, error)
}

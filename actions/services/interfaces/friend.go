package interfaces

import "go.mongodb.org/mongo-driver/bson/primitive"

type InviteRequest struct {
	TargetMailAddress string `json:"targetMailAddress" example:"winner@mail.com"`
}

type FriendShipResponse struct {
	ID     string `json:"_id"`
	Friend string `json:"friend_email" example:"winner@mail.com"`
}

type CheckInviteResponse struct {
	InviteId primitive.ObjectID `json:"inviteId" example:"1"`
	Username string             `json:"username" example:"winnerkypt"`
	Email    string             `json:"email" example:"winner@mail.com"`
	Image    string             `json:"image"`
}

type FriendService interface {
	InviteFriend(string, InviteRequest) (interface{}, error)
	CheckFriendInvite(string) (interface{}, error)
	RejectInvitation(string, string) (interface{}, error)
	RejectAllInvitation(string) (interface{}, error)
	AcceptInvitation(string, string) (interface{}, error)
	AcceptAllInvitations(string) (interface{}, error)
	GetFriend(string) (interface{}, error)
	DeleteFriend(token string, id string) (interface{}, error)
}

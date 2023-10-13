package interfaces

type FriendInvitation struct {
	ID         int `db:"id"`
	ReceiverId int `db:"receiver_id"`
	SenderId   int `db:"sender_id"`
}

type FriendInvitationRepository interface {
	Create(FriendInvitation) (*FriendInvitation, error)
	GetInvitationByReceiverId(int) ([]FriendInvitation, error)
	Delete(int) error
	GetByReceiverIdAndSenderId(int, int) (*FriendInvitation, error)
	GetInvitationByIdAndReceiverId(int, int) (*FriendInvitation, error)
}

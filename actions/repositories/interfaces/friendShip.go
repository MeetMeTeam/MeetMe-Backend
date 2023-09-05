package interfaces

type Friendship struct {
	ID      int `db:"id"`
	UserId1 int `db:"user_id1"`
	UserId2 int `db:"user_id2"`
}

type FriendshipRepository interface {
	Create(int, int) (*Friendship, error)
	GetFriendByReceiverAndSender(int, int) (*Friendship, error)
	//GetInvitationByReceiverId(int) ([]FriendInvitation, error)
	//Delete(int) error
}

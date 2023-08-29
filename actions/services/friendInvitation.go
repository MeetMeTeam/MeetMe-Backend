package services

import (
	"log"
	"meetme/be/actions/repositories"
	repoInt "meetme/be/actions/repositories/interfaces"
	"meetme/be/actions/services/interfaces"
	"meetme/be/errs"
	"meetme/be/utils"
)

type friendInvitationService struct {
	userRepo repositories.FriendInvitationRepository
}

func NewFriendInvitationService(userRepo repositories.FriendInvitationRepository) friendInvitationService {
	return friendInvitationService{userRepo: userRepo}
}

func (s friendInvitationService) InviteFriend(request interfaces.InviteRequest) (interface{}, error) {

	newUser := repoInt.FriendInvitation{
		ReceiverId: request.ReceiverId,
		SenderId:   request.SenderId,
	}

	_, err := s.userRepo.Create(newUser)
	if err != nil {
		log.Println(err)
		return nil, errs.NewInternalError(err.Error())
	}

	response := utils.ErrorResponse{
		Message: "Invite friend success",
	}

	return response, nil
}

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
	inviteRepo repositories.FriendInvitationRepository
	userRepo   repositories.UserRepository
}

func NewFriendInvitationService(inviteRepo repositories.FriendInvitationRepository, userRepo repositories.UserRepository) friendInvitationService {
	return friendInvitationService{inviteRepo: inviteRepo, userRepo: userRepo}
}

func (s friendInvitationService) InviteFriend(request interfaces.InviteRequest) (interface{}, error) {

	newUser := repoInt.FriendInvitation{
		ReceiverId: request.ReceiverId,
		SenderId:   request.SenderId,
	}

	_, err := s.inviteRepo.Create(newUser)
	if err != nil {
		log.Println(err)
		return nil, errs.NewInternalError(err.Error())
	}

	response := utils.ErrorResponse{
		Message: "Invite friend success",
	}

	return response, nil
}

func (s friendInvitationService) CheckFriendInvite(receiverId int) (interface{}, error) {
	result, err := s.inviteRepo.GetInvitationByReceiverId(receiverId)
	if err != nil {
		log.Println(err)
		return nil, errs.NewInternalError(err.Error())
	}

	if len(result) == 0 {
		return nil, errs.NewNotFoundError("Friend invitation is empty")
	}
	userResponses := []interfaces.RegisterResponse{}
	for _, receiver := range result {
		user, err := s.userRepo.GetById(receiver.SenderId)
		if err != nil {
			log.Println(err)
			return nil, errs.NewInternalError(err.Error())
		}

		userResponse := interfaces.RegisterResponse{
			ID:        user.ID,
			Firstname: user.Firstname,
			Lastname:  user.Lastname,
			Email:     user.Email,
			Birthday:  user.Birthday,
		}
		userResponses = append(userResponses, userResponse)
	}

	response := utils.DataResponse{
		Data:    userResponses,
		Message: "Get sender success.",
	}
	return response, nil
}

func (s friendInvitationService) RejectInvitation(req interfaces.InviteRequest) (interface{}, error) {
	err := s.inviteRepo.Delete(req.ReceiverId, req.SenderId)
	if err != nil {
		log.Println(err)
		return nil, errs.NewInternalError(err.Error())
	}
	return utils.ErrorResponse{
		Message: "Reject Friend Success",
	}, nil
}

func (s friendInvitationService) AcceptInvitation(req interfaces.InviteRequest) (interface{}, error) {
	return nil, nil
}

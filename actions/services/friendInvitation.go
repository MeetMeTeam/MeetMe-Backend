package services

import (
	"errors"
	"gorm.io/gorm"
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
	friendRepo repositories.FriendshipRepository
}

func NewFriendInvitationService(inviteRepo repositories.FriendInvitationRepository, userRepo repositories.UserRepository, friendRepo repositories.FriendshipRepository) friendInvitationService {
	return friendInvitationService{inviteRepo: inviteRepo, userRepo: userRepo, friendRepo: friendRepo}
}

func (s friendInvitationService) InviteFriend(token string, request interfaces.InviteRequest) (interface{}, error) {

	email, err := utils.IsTokenValid(token)
	if err != nil {
		return nil, err
	}

	if email == request.ReceiverEmail {
		return nil, errs.NewBadRequestError("Can not add yourself.")
	}

	sender, err := s.userRepo.GetByEmail(email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errs.NewNotFoundError("User not found.")
		}
		return nil, errs.NewInternalError(err.Error())
	}
	receiver, err := s.userRepo.GetByEmail(request.ReceiverEmail)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errs.NewNotFoundError("User not found.")
		}
		return nil, errs.NewInternalError(err.Error())
	}

	isInvite, err := s.inviteRepo.GetByReceiverIdAndSenderId(receiver.ID, sender.ID)
	if isInvite != nil {
		return nil, errs.NewBadRequestError("This email is already sent!")
	}

	isFriend, err := s.friendRepo.GetFriendByReceiverAndSender(receiver.ID, sender.ID)
	if isFriend != nil {
		return nil, errs.NewBadRequestError("They are friends now!")
	}
	newUser := repoInt.FriendInvitation{
		SenderId:   sender.ID,
		ReceiverId: receiver.ID,
	}

	_, err = s.inviteRepo.Create(newUser)
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

func (s friendInvitationService) RejectInvitation(token string, inviteId int) (interface{}, error) {
	email, err := utils.IsTokenValid(token)
	if err != nil {
		return nil, err
	}
	user, err := s.userRepo.GetByEmail(email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errs.NewNotFoundError("User not found.")
		}
		return nil, errs.NewInternalError(err.Error())
	}

	isInvite, err := s.inviteRepo.GetInvitationByIdAndReceiverId(inviteId, user.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errs.NewNotFoundError("Invitation not found.")
		}
		return nil, errs.NewInternalError(err.Error())
	} else if isInvite == nil {
		return nil, errs.NewNotFoundError("Invitation not found.")
	}
	err = s.inviteRepo.Delete(inviteId)
	if err != nil {
		log.Println(err)
		return nil, errs.NewInternalError(err.Error())
	}
	return utils.ErrorResponse{
		Message: "Reject Friend Success",
	}, nil
}

func (s friendInvitationService) AcceptInvitation(req interfaces.InviteRequest) (interface{}, error) {
	//result, err := s.friendRepo.Create(req.SenderId, req.ReceiverId)
	//if err != nil {
	//	log.Println(err)
	//	return nil, errs.NewInternalError(err.Error())
	//}
	//err = s.inviteRepo.Delete(req.ReceiverId, req.SenderId)
	//if err != nil {
	//	log.Println(err)
	//	return nil, errs.NewInternalError(err.Error())
	//}

	//return utils.DataResponse{
	//	Data:    result,
	//	Message: "Add friend success",
	//}, nil

	return nil, nil

}

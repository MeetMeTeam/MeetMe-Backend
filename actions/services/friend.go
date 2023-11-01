package services

import (
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
	"log"
	"meetme/be/actions/repositories"
	repoInt "meetme/be/actions/repositories/interfaces"
	"meetme/be/actions/services/interfaces"
	"meetme/be/errs"
	"meetme/be/utils"
)

type friendService struct {
	friendRepo repositories.FriendRepository
	userRepo   repositories.UserRepository
	//friendRepo repositories.FriendshipRepository
}

func NewFriendService(friendRepo repositories.FriendRepository, userRepo repositories.UserRepository) friendService {
	return friendService{friendRepo: friendRepo, userRepo: userRepo}
}

func (s friendService) InviteFriend(token string, request interfaces.InviteRequest) (interface{}, error) {

	email, err := utils.IsTokenValid(token)
	if err != nil {
		return nil, err
	}

	if email == request.TargetMailAddress {
		return nil, errs.NewBadRequestError("Can not add yourself.")
	}

	sender, err := s.userRepo.GetByEmail(email)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errs.NewNotFoundError("User not found.")
		}
		return nil, errs.NewInternalError(err.Error())
	}
	receiver, err := s.userRepo.GetByEmail(request.TargetMailAddress)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errs.NewNotFoundError("User not found.")
		}
		return nil, errs.NewInternalError(err.Error())
	}

	isInvite, err := s.friendRepo.GetByReceiverIdAndSenderId(receiver.ID, sender.ID)
	if isInvite != nil {
		if isInvite.Status == "PENDING" {
			return nil, errs.NewBadRequestError("This email is already sent!")
		} else if isInvite.Status == "FRIEND" {
			return nil, errs.NewBadRequestError("They are friends now!")
		}

	}

	newUser := repoInt.FriendRequest{
		Receiver: receiver.ID,
		Sender:   sender.ID,
	}

	_, err = s.friendRepo.Create(newUser)
	if err != nil {
		log.Println(err)
		return nil, errs.NewInternalError(err.Error())
	}

	response := utils.ErrorResponse{
		Message: "Invite friend success",
	}

	return response, nil
}

func (s friendService) CheckFriendInvite(token string) (interface{}, error) {
	email, err := utils.IsTokenValid(token)
	if err != nil {
		return nil, err
	}
	receiver, err := s.userRepo.GetByEmail(email)
	if err != nil {
		log.Println("User not found.")
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errs.NewNotFoundError("User not found.")
		}
		return nil, errs.NewInternalError(err.Error())
	}

	results, err := s.friendRepo.GetInvitationByReceiverId(receiver.ID)

	if err != nil {
		log.Println("Friend invitation is empty")
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errs.NewNotFoundError("Friend invitation is empty")
		}
		return nil, errs.NewInternalError(err.Error())
	}

	if len(results) == 0 {
		return nil, errs.NewNotFoundError("Friend invitation is empty")
	}
	userResponses := []interfaces.CheckInviteResponse{}
	for _, result := range results {
		user, err := s.userRepo.GetById(result.Sender)
		if err != nil {
			log.Println("err")
			log.Println(err)
			return nil, errs.NewInternalError(err.Error())
		}

		userResponse := interfaces.CheckInviteResponse{
			InviteId: result.ID,
			Username: user.Username,
			Email:    user.Email,
		}
		userResponses = append(userResponses, userResponse)
	}

	response := utils.DataResponse{
		Data:    userResponses,
		Message: "Get sender success.",
	}
	return response, nil
}

func (s friendService) AcceptInvitation(token string, inviteId string) (interface{}, error) {
	email, err := utils.IsTokenValid(token)
	if err != nil {
		return nil, err
	}
	_, err = s.userRepo.GetByEmail(email)
	if err != nil {
		log.Println("get from token")
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errs.NewNotFoundError("User not found.")
		}
		return nil, errs.NewInternalError(err.Error())
	}

	//isInvite, err := s.friendRepo.GetInvitationByReceiverId(inviteId, user.ID)
	//if err != nil {
	//	if errors.Is(err, gorm.ErrRecordNotFound) {
	//		return nil, errs.NewNotFoundError("Invitation not found.")
	//	}
	//	return nil, errs.NewInternalError(err.Error())
	//} else if isInvite == nil {
	//	return nil, errs.NewNotFoundError("Invitation not found.")
	//}
	id, err := primitive.ObjectIDFromHex(inviteId)
	if err != nil {
		log.Println(err)
		return nil, errs.NewInternalError(err.Error())
	}
	result, err := s.friendRepo.UpdateStatus(id)
	if err != nil {
		log.Println("Update Status")
		return nil, errs.NewInternalError(err.Error())
	}

	log.Print(result)
	user1, err := s.userRepo.GetById(result.Sender)
	if err != nil {
		log.Println("get user 1")
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errs.NewNotFoundError("User not found.")
		}
		return nil, errs.NewInternalError(err.Error())
	}

	user2, err := s.userRepo.GetById(result.Receive)
	if err != nil {
		log.Println("get user 2")
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errs.NewNotFoundError("User not found.")
		}
		return nil, errs.NewInternalError(err.Error())
	}
	return utils.DataResponse{
		Data: interfaces.FriendShipResponse{
			User1: user1.Email,
			User2: user2.Email,
		},
		Message: "Add friend success",
	}, nil

	return nil, nil
}

func (s friendService) RejectInvitation(token string, inviteId string) (interface{}, error) {
	email, err := utils.IsTokenValid(token)
	if err != nil {
		return nil, err
	}
	_, err = s.userRepo.GetByEmail(email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errs.NewNotFoundError("User not found.")
		}
		return nil, errs.NewInternalError(err.Error())
	}

	//isInvite, err := s.friendRepo.GetInvitationByIdAndReceiverId(inviteId, user.ID.Hex())
	//if err != nil {
	//	if errors.Is(err, gorm.ErrRecordNotFound) {
	//		return nil, errs.NewNotFoundError("Invitation not found.")
	//	}
	//	return nil, errs.NewInternalError(err.Error())
	//} else if isInvite == nil {
	//	return nil, errs.NewNotFoundError("Invitation not found.")
	//}
	id, err := primitive.ObjectIDFromHex(inviteId)
	if err != nil {
		log.Println(err)
		return nil, errs.NewInternalError(err.Error())
	}
	err = s.friendRepo.Delete(id)
	if err != nil {
		log.Println(err)
		return nil, errs.NewInternalError(err.Error())
	}
	return utils.ErrorResponse{
		Message: "Reject Friend Success",
	}, nil
}

package services

import (
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
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

	results, err := s.friendRepo.GetByReceiverId(receiver.ID, "PENDING")

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
	user, err := s.userRepo.GetByEmail(email)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errs.NewNotFoundError("User not found.")
		}
		return nil, errs.NewInternalError(err.Error())
	}

	id, err := primitive.ObjectIDFromHex(inviteId)
	if err != nil {
		log.Println(err)
		return nil, errs.NewInternalError(err.Error())
	}
	isInvite, err := s.friendRepo.GetByIdAndReceiverIdAndStatus(id, user.ID, "PENDING")
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errs.NewNotFoundError("Invitation not found.")
		}
		return nil, errs.NewInternalError(err.Error())
	} else if isInvite == nil {
		return nil, errs.NewNotFoundError("Invitation not found.")
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
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errs.NewNotFoundError("User not found.")
		}
		return nil, errs.NewInternalError(err.Error())
	}

	user2, err := s.userRepo.GetById(result.Receive)
	if err != nil {
		log.Println("get user 2")
		if errors.Is(err, mongo.ErrNoDocuments) {
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
	user, err := s.userRepo.GetByEmail(email)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errs.NewNotFoundError("User not found.")
		}
		return nil, errs.NewInternalError(err.Error())
	}

	id, err := primitive.ObjectIDFromHex(inviteId)
	if err != nil {
		log.Println(err)
		return nil, errs.NewInternalError(err.Error())
	}

	isInvite, err := s.friendRepo.GetByIdAndReceiverIdAndStatus(id, user.ID, "PENDING")
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errs.NewNotFoundError("Invitation not found.")
		}
		return nil, errs.NewInternalError(err.Error())
	} else if isInvite == nil {
		return nil, errs.NewNotFoundError("Invitation not found.")
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

func (s friendService) GetFriend(token string) (interface{}, error) {
	email, err := utils.IsTokenValid(token)
	if err != nil {
		return nil, err
	}
	user, err := s.userRepo.GetByEmail(email)
	if err != nil {
		log.Println("User not found.")
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errs.NewNotFoundError("User not found.")
		}
		return nil, errs.NewInternalError(err.Error())
	}

	results, err := s.friendRepo.GetByUserId(user.ID, "FRIEND")

	if err != nil {
		log.Println("Friend list is empty")
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errs.NewNotFoundError("Friend list is empty")
		}
		return nil, errs.NewInternalError(err.Error())
	}

	if len(results) == 0 {
		return nil, errs.NewNotFoundError("Friend list is empty")
	}
	userResponses := []interfaces.RegisterResponse{}
	for _, result := range results {
		id := primitive.ObjectID{}
		if result.Sender != user.ID {
			id = result.Sender
		} else if result.Receive != user.ID {
			id = result.Receive
		}

		user, err := s.userRepo.GetById(id)
		if err != nil {
			log.Println(err)
			return nil, errs.NewInternalError(err.Error())
		}

		userResponse := interfaces.RegisterResponse{
			Username:  user.Username,
			Firstname: user.Firstname,
			Lastname:  user.Lastname,
			Birthday:  user.Birthday,
			Email:     user.Email,
		}
		userResponses = append(userResponses, userResponse)
	}

	response := utils.DataResponse{
		Data:    userResponses,
		Message: "Get friend list success.",
	}
	return response, nil
}

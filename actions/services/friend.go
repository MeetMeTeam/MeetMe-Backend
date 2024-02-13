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

	if email.Email == request.TargetMailAddress {
		return nil, errs.NewBadRequestError("Can not add yourself.")
	}

	sender, err := s.userRepo.GetByEmail(email.Email)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errs.NewBadRequestError("User not found.")
		}
		return nil, errs.NewInternalError(err.Error())
	}
	receiver, err := s.userRepo.GetByEmail(request.TargetMailAddress)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errs.NewBadRequestError("User not found.")
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

	createUser, err := s.friendRepo.Create(newUser)
	if err != nil {
		log.Println(err)
		return nil, errs.NewInternalError(err.Error())
	}

	user, err := s.userRepo.GetById(createUser.Receiver)
	if err != nil {
		log.Println(err)
		return nil, errs.NewInternalError(err.Error())
	}
	response := utils.DataResponse{
		Data: interfaces.ListUserResponse{
			ID:          user.ID.Hex(),
			Username:    user.Username,
			DisplayName: user.DisplayName,
			Birthday:    user.Birthday,
			Email:       user.Email,
		},
		Message: "Invite friend success",
	}

	return response, nil
}

func (s friendService) CheckFriendInvite(token string) (interface{}, error) {
	email, err := utils.IsTokenValid(token)
	if err != nil {
		return nil, err
	}
	receiver, err := s.userRepo.GetByEmail(email.Email)
	if err != nil {
		log.Println("User not found.")
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errs.NewBadRequestError("User not found.")
		}
		return nil, errs.NewInternalError(err.Error())
	}

	results, err := s.friendRepo.GetByReceiverId(receiver.ID, "PENDING")

	if err != nil {
		log.Println("Friend invitation is empty")
		if errors.Is(err, mongo.ErrNoDocuments) {
			return utils.DataResponse{
				Data:    []int{},
				Message: "Get sender success.",
			}, nil
		}
		return nil, errs.NewInternalError(err.Error())
	}

	if len(results) == 0 {
		return utils.DataResponse{
			Data:    []int{},
			Message: "Get sender success.",
		}, nil
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
	user, err := s.userRepo.GetByEmail(email.Email)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errs.NewBadRequestError("User not found.")
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
			return nil, errs.NewBadRequestError("Invitation not found.")
		}
		return nil, errs.NewInternalError(err.Error())
	} else if isInvite == nil {
		return nil, errs.NewBadRequestError("Invitation not found.")
	}

	result, err := s.friendRepo.UpdateStatus(primitive.NilObjectID, id)
	if err != nil {
		log.Println("Update Status")
		return nil, errs.NewInternalError(err.Error())
	}

	log.Print(result)
	sender, err := s.userRepo.GetById(result[0].Sender)
	if err != nil {
		log.Println("get user 1")
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errs.NewBadRequestError("User not found.")
		}
		return nil, errs.NewInternalError(err.Error())
	}

	return utils.DataResponse{
		Data: interfaces.FriendShipResponse{
			ID:     user.ID.Hex(),
			Friend: sender.Email,
		},
		Message: "Add friend success",
	}, nil

	return nil, nil
}

func (s friendService) AcceptAllInvitations(token string) (interface{}, error) {
	email, err := utils.IsTokenValid(token)
	if err != nil {
		return nil, err
	}
	user, err := s.userRepo.GetByEmail(email.Email)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errs.NewBadRequestError("User not found.")
		}
		return nil, errs.NewInternalError(err.Error())
	}

	isInvite, err := s.friendRepo.GetByReceiverId(user.ID, "PENDING")
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errs.NewBadRequestError("Invitation not found.")
		}
		return nil, errs.NewInternalError(err.Error())
	} else if isInvite == nil {
		return nil, errs.NewBadRequestError("Invitation not found.")
	}

	results, err := s.friendRepo.UpdateStatus(user.ID, primitive.NilObjectID)
	if err != nil {
		log.Println("Update Status")
		return nil, errs.NewInternalError(err.Error())
	}

	friendResponses := []interfaces.FriendShipResponse{}
	for _, result := range results {
		sender, err := s.userRepo.GetById(result.Sender)
		if err != nil {
			log.Println("get user 1")
			if errors.Is(err, mongo.ErrNoDocuments) {
				return nil, errs.NewBadRequestError("User not found.")
			}
			return nil, errs.NewInternalError(err.Error())
		}
		friendResponse := interfaces.FriendShipResponse{
			ID:     user.ID.Hex(),
			Friend: sender.Email,
		}
		friendResponses = append(friendResponses, friendResponse)
	}

	return utils.DataResponse{
		Data:    friendResponses,
		Message: "Add friend success",
	}, nil

}

func (s friendService) RejectInvitation(token string, inviteId string) (interface{}, error) {
	email, err := utils.IsTokenValid(token)
	if err != nil {
		return nil, err
	}
	user, err := s.userRepo.GetByEmail(email.Email)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errs.NewBadRequestError("User not found.")
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
			return nil, errs.NewBadRequestError("Invitation not found.")
		}
		return nil, errs.NewInternalError(err.Error())
	} else if isInvite == nil {
		return nil, errs.NewBadRequestError("Invitation not found.")
	}

	err = s.friendRepo.Delete(primitive.NilObjectID, id, "PENDING")
	if err != nil {
		log.Println(err)
		return nil, errs.NewInternalError(err.Error())
	}
	return utils.ErrorResponse{
		Message: "Reject Invite Success",
	}, nil
}

func (s friendService) RejectAllInvitation(token string) (interface{}, error) {
	email, err := utils.IsTokenValid(token)
	if err != nil {
		return nil, err
	}
	user, err := s.userRepo.GetByEmail(email.Email)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errs.NewBadRequestError("User not found.")
		}
		return nil, errs.NewInternalError(err.Error())
	}

	isInvite, err := s.friendRepo.GetByReceiverId(user.ID, "PENDING")
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errs.NewBadRequestError("Invitation not found.")
		}
		return nil, errs.NewInternalError(err.Error())
	} else if isInvite == nil {
		return nil, errs.NewBadRequestError("Invitation not found.")
	}

	err = s.friendRepo.Delete(user.ID, primitive.NilObjectID, "PENDING")
	if err != nil {
		log.Println(err)
		return nil, errs.NewInternalError(err.Error())
	}
	return utils.ErrorResponse{
		Message: "Reject All Invite Success",
	}, nil
}

func (s friendService) GetFriend(token string) (interface{}, error) {
	email, err := utils.IsTokenValid(token)
	if err != nil {
		return nil, err
	}
	user, err := s.userRepo.GetByEmail(email.Email)
	if err != nil {
		log.Println("User not found.")
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errs.NewBadRequestError("User not found.")
		}
		return nil, errs.NewInternalError(err.Error())
	}

	results, err := s.friendRepo.GetByUserId(user.ID, "FRIEND")

	if err != nil {
		log.Println("Friend list is empty")
		if errors.Is(err, mongo.ErrNoDocuments) {
			return utils.DataResponse{
				Data:    []int{},
				Message: "Get friend list success.",
			}, nil
		}
		return nil, errs.NewInternalError(err.Error())
	}

	if len(results) == 0 {
		return utils.DataResponse{
			Data:    []int{},
			Message: "Get friend list success.",
		}, nil
	}
	userResponses := []interfaces.ListUserResponse{}
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

		userResponse := interfaces.ListUserResponse{
			ID:          user.ID.Hex(),
			Username:    user.Username,
			DisplayName: user.DisplayName,
			Birthday:    user.Birthday,
			Email:       user.Email,
		}
		userResponses = append(userResponses, userResponse)
	}

	response := utils.DataResponse{
		Data:    userResponses,
		Message: "Get friend list success.",
	}
	return response, nil
}

func (s friendService) DeleteFriend(token string, id string) (interface{}, error) {
	email, err := utils.IsTokenValid(token)
	if err != nil {
		return nil, err
	}
	user, err := s.userRepo.GetByEmail(email.Email)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errs.NewBadRequestError("User not found.")
		}
		return nil, errs.NewInternalError(err.Error())
	}

	friendId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println(err)
		return nil, errs.NewInternalError(err.Error())
	}

	isInvite, err := s.friendRepo.GetByReceiverIdAndSenderId(friendId, user.ID)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errs.NewBadRequestError("Invitation not found.")
		}
		return nil, errs.NewInternalError(err.Error())
	} else if isInvite == nil {
		return nil, errs.NewBadRequestError("Invitation not found.")
	}

	err = s.friendRepo.Delete(primitive.NilObjectID, isInvite.ID, "FRIEND")
	if err != nil {
		log.Println(err)
		return nil, errs.NewInternalError(err.Error())
	}

	friend, err := s.userRepo.GetById(friendId)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errs.NewBadRequestError("User not found.")
		}
		return nil, errs.NewInternalError(err.Error())
	}
	return utils.DataResponse{
		Data: interfaces.FriendShipResponse{
			ID:     friend.ID.Hex(),
			Friend: friend.Email,
		},
		Message: "Delete Friend Success",
	}, nil
}

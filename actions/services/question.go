package services

import (
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"meetme/be/actions/repositories"
	repoInt "meetme/be/actions/repositories/interfaces"
	"meetme/be/actions/services/interfaces"
	"meetme/be/errs"
	"meetme/be/utils"
)

type questionService struct {
	questionRepo repositories.QuestionRepository
}

func NewQuestionService(questionRepo repositories.QuestionRepository) interfaces.QuestionService {
	return questionService{questionRepo: questionRepo}
}

func (s questionService) GetQuestions(lang string, cate string) (interface{}, error) {
	var questions []repoInt.QuestionResponse
	var err error
	//var cateId primitive.ObjectID

	if cate == "" {
		questions, err = s.questionRepo.GetAll()
	} else {
		cateId, err := primitive.ObjectIDFromHex(cate)
		if err != nil {
			return nil, errs.NewInternalError(err.Error())
		}
		questions, err = s.questionRepo.GetByCategory(cateId)
	}

	if err != nil {
		return nil, errs.NewInternalError(err.Error())
	}
	if questions == nil {
		return utils.DataResponse{
			Data:    []string{},
			Message: "Get questions success.",
		}, nil
	}

	var response interface{}
	if lang == "thai" {
		questionRes := []interfaces.ThaiQuestionsResponse{}
		for _, ques := range questions {

			cate, err := s.questionRepo.GetCateById(ques.Category)
			if err != nil {
				if errors.Is(err, mongo.ErrNoDocuments) {
					return nil, errs.NewBadRequestError("Category not found.")
				}
				return nil, errs.NewInternalError(err.Error())
			}

			questionResponse := interfaces.ThaiQuestionsResponse{
				ID:       ques.ID.Hex(),
				Thai:     ques.Thai,
				Category: cate.Name,
			}
			questionRes = append(questionRes, questionResponse)

		}
		response = questionRes
	} else if lang == "eng" {
		questionRes := []interfaces.EngQuestionsResponse{}
		for _, ques := range questions {
			cate, err := s.questionRepo.GetCateById(ques.Category)
			if err != nil {
				if errors.Is(err, mongo.ErrNoDocuments) {
					return nil, errs.NewBadRequestError("Category not found.")
				}
				return nil, errs.NewInternalError(err.Error())
			}
			questionResponse := interfaces.EngQuestionsResponse{
				ID:       ques.ID.Hex(),
				Eng:      ques.Eng,
				Category: cate.Name,
			}
			questionRes = append(questionRes, questionResponse)

		}
		response = questionRes
	} else {
		questionRes := []interfaces.QuestionsResponse{}
		for _, ques := range questions {

			cate, err := s.questionRepo.GetCateById(ques.Category)
			if err != nil {
				if errors.Is(err, mongo.ErrNoDocuments) {
					return nil, errs.NewBadRequestError("Category not found.")
				}
				return nil, errs.NewInternalError(err.Error())
			}
			questionResponse := interfaces.QuestionsResponse{
				ID:       ques.ID.Hex(),
				Thai:     ques.Thai,
				Eng:      ques.Eng,
				Category: cate.Name,
			}
			questionRes = append(questionRes, questionResponse)

		}
		response = questionRes
	}

	return utils.DataResponse{
		Data:    response,
		Message: "Get questions success.",
	}, nil
}

func (s questionService) GetCategories() interface{} {
	categories, _ := s.questionRepo.GetCategoryAll()
	questionRes := []interfaces.CategoryResponse{}
	for _, cate := range categories {

		questionResponse := interfaces.CategoryResponse{
			ID:   cate.ID.Hex(),
			Name: cate.Name,
		}
		questionRes = append(questionRes, questionResponse)

	}

	return utils.DataResponse{
		Data:    questionRes,
		Message: "Get categories success.",
	}
}

func (s questionService) CreateQuestion(token string, question interfaces.QuestionRequest) (interface{}, error) {
	result, err := utils.IsTokenValid(token)
	if err != nil {
		return nil, err
	}
	if result.IsAdmin == false {
		return nil, errs.NewForbiddenError("You don't have permission.")
	}

	cateId, err := primitive.ObjectIDFromHex(question.Category)
	if err != nil {
		return nil, errs.NewInternalError(err.Error())
	}

	newQuestion := repoInt.Question{
		Eng:      question.Eng,
		Thai:     question.Thai,
		Category: cateId,
	}
	questions, err := s.questionRepo.CreateQuestions(newQuestion)
	if err != nil {
		return nil, err
	}
	return utils.DataResponse{Data: questions, Message: "Create question success"}, nil
}

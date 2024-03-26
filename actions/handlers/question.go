package handlers

import (
	"github.com/labstack/echo/v4"
	svInter "meetme/be/actions/services/interfaces"
	"meetme/be/errs"
	"meetme/be/utils"
	"net/http"
)

type QuestionHandler struct {
	questionService svInter.QuestionService
}

func NewQuestionHandler(questionService svInter.QuestionService) QuestionHandler {
	return QuestionHandler{questionService: questionService}
}

// GetQuestions godoc
//
//	@Summary		Get Questions.
//	@Description	Get questions by category or language .
//	@Tags			questions
//	@Accept			json
//	@Produce		json
//	@Param			lang	query	string	false	"Question's language [thai/eng]"
//	@Param			category	query	string	false	"ID Category of question"
//	@Success		200		{object}	utils.DataResponse
//	@Router			/questions [get]
//
// @Security BearerAuth
func (h QuestionHandler) GetQuestions(c echo.Context) error {
	lang := c.QueryParam("lang")
	cate := c.QueryParam("category")
	themes, err := h.questionService.GetQuestions(lang, cate)
	if err != nil {

		appErr, ok := err.(errs.AppError)
		if ok {
			return c.JSON(appErr.Code, utils.ErrorResponse{
				Message: appErr.Message,
			})
		}
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, themes)
}

// GetCategories godoc
//
//	@Summary		Get categories.
//	@Description	Get question's categories.
//	@Tags			questions
//	@Accept			json
//	@Produce		json
//	@Success		200		{object}	utils.DataResponse
//	@Router			/questions/categories [get]
func (h QuestionHandler) GetCategories(c echo.Context) error {

	themes := h.questionService.GetCategories()
	return c.JSON(http.StatusOK, themes)
}

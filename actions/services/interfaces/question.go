package interfaces

type QuestionsResponse struct {
	ID       string `json:"id"`
	Thai     string `json:"thai"`
	Eng      string `json:"eng"`
	Category string `json:"category"`
}

type ThaiQuestionsResponse struct {
	ID       string `json:"id"`
	Thai     string `json:"thai"`
	Category string `json:"category"`
}
type EngQuestionsResponse struct {
	ID       string `json:"id"`
	Eng      string `json:"eng"`
	Category string `json:"category"`
}
type CategoryResponse struct {
	ID   string `json:"_id"`
	Name string `json:"name"`
}
type QuestionService interface {
	GetQuestions(string, string) (interface{}, error)
	GetCategories() interface{}
}

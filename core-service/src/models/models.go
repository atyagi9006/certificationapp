package models

type Type struct {
	TypeName string `json:"string"`
}

type User struct {
	UserID   string `json:"userId"`
	UserName string `json:"userName"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	Type     string `json:"type"`
}

type Admin struct {
	UserID    string `json:"userId"`
	AdminID   string `json:"adminId"`
	AdminName string `json:"adminName"`
}

type Candidtate struct {
	CandidtateID string `json:"candidateId"`
	//EligibilityList []Eligibility `json:"eligibilityList"`
	ExamAttemptList []*ExamAttempt `json:"examAttemptList"`
}

type CurrentTest struct {
	CandidtateID string `json:"candidateId"`
	TestID       string `json:"testId"`
}

type Category struct {
	CategoryID   string `json:"categoryId"`
	CategoryName string `json:"categoryName"`
}

type Eligibility struct {
	EligibilityID    string   `json:"eligibilityId"`
	EligibilityLevel string   `json:"eligibilityLevel"`
	TestIDList       []string `json:"testIdList"`
}

type ExamAttempt struct {
	CategoryID string              `json:"categoryId"`
	Score      string              `json:"score"`
	TimeSpent  string              `json:"timeSpent"`
	Questions  []*QuestionAttempts `json:"questionsAttempted"`
	//Result     string              `json:"result"`
}
type QuestionAttempts struct {
	QuestionID      string `json:"questionId"`
	AttemptedAnswer string `json:"answer"`
	CorrectAnswer   string `json:"correct"`
	/* Category        string   `json:"category"`
	Type            string   `json:"type"`
	Difficulty      string   `json:"difficulty"`
	Question        string   `json:"question"`
	Option          []string `json:"option"` */
}

type TestResult struct {
	CategoryID string `json:"categoryId"`
	TestID     string `json:"testID"`
	Result     string `json:"result"`
	Score      string `json:"score"`
	TimeSpent  string `json:"timeSpent"`
}

type Test struct {
	Category       string   `json:"category"`
	TestID         string   `json:"testId"`
	TestName       string   `json:"testName"`
	QuestionIDList []string `json:"questionIDList"`
	TestDuration   string   `json:"testDuration"`
}

type Question_dao struct {
	QuestionID       string   `json:"questionId"`
	Category         string   `json:"category"`
	Type             string   `json:"type"`
	Difficulty       string   `json:"difficulty"`
	Question         string   `json:"question"`
	Answer           string   `json:"correct_answer"`
	IncorrectAnswers []string `json:"incorrect_answers"`
}

type Question struct {
	QuestionID string   `json:"questionId"`
	Category   string   `json:"category"`
	Type       string   `json:"type"`
	Difficulty string   `json:"difficulty"`
	Question   string   `json:"question"`
	Options    []string `json:"options"`
}

type Question_v1 struct {
	QuestionID   string `json:"questionId"`
	QuestionText string `json:"question-text"`
	IsMCQ        bool   `json:"question-mcq"`
}

type Option struct {
	QuestionID string `json:"questionId"`
	OptionID   string `json:"optionId"`
	OptionText string `json:"optionText"`
}

type QuestionToAnswerCorrectMapping struct {
	QuestionID          string   `json:"questionId"`
	CorrectOptionIDList []string `json:"correctOptionIdList"`
}

type AnswerToQuestion struct {
	QuestionID string `json:"questionId"`
	Answer     string `json:"answer"`
}
type CurrentQuestionOptionElement struct {
	QuestionID string
	OptionIDs  []string
}

type CandidateAttemptedQuestionList struct {
	CandidateID           string
	AttemptedQuestionList []CurrentQuestionOptionElement
}

package models

type Type struct {
	TypeName string `json:"string"`
}

type User struct {
	UserID   string
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

type Candidate struct {
	Key             string         `json:"_key"`
	CandidateID     string         `json:"candidateId"`
	ExamAttemptList []*ExamAttempt `json:"examAttemptList"`
}
type ExamAttempt struct {
	CategoryID string              `json:"categoryId"`
	Questions  []*QuestionAttempts `json:"questionsAttempted"`
	Score      string              `json:"score"`
	TimeSpent  string              `json:"timeSpent"`
	//Result     string              `json:"result"`
}
type QuestionAttempts struct {
	QuestionID      string `json:"questionId"`
	AttemptedAnswer string `json:"attemptedAnswer"`
	CorrectAnswer   string `json:"correct"`
	/* Category        string   `json:"category"`
	Type            string   `json:"type"`
	Difficulty      string   `json:"difficulty"`
	Question        string   `json:"question"`
	Option          []string `json:"option"` */
}

type CandidateUpdateRes struct {
	ID     string `json:"id"`
	Status bool   `json:"status"`
}

type CurrentTest struct {
	CandidateID string `json:"candidateId"`
	TestID      string `json:"testId"`
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

type TestResult struct {
	CategoryID string `json:"categoryId"`
	TestID     string `json:"testID"`
	Result     string `json:"result"`
	Score      string `json:"score"`
}

type Test struct {
	CategoryID     string   `json:"categoryId"`
	TestID         string   `json:"testId"`
	TestName       string   `json:"testName"`
	QuestionIDList []string `json:"questionIDList"`
	TestDuration   string   `json:testDuration`
}

type Question struct {
	QuestionID       string   `json:"questionId"`
	Category         string   `json:"category"`
	Type             string   `json:"type"`
	Difficulty       string   `json:"difficulty"`
	Question         string   `json:"question"`
	Answer           string   `json:"correct_answer"`
	IncorrectAnswers []string `json:"incorrect_answers"`
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

type CurrentQuestionOptionElement struct {
	QuestionID string
	OptionIDs  []string
}

type CandidateAttemptedQuestionList struct {
	CandidateID           string
	AttemptedQuestionList []CurrentQuestionOptionElement
}

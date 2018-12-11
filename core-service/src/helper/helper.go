package helper

import (
	"context"
	"errors"
	"log"
	"strings"

	"github.com/atyagi9006/certificationapp/core-service/src/config"

	"github.com/atyagi9006/certificationapp/core-service/src/cache"

	"github.com/atyagi9006/certificationapp/core-service/src/util"
	"github.com/atyagi9006/certificationapp/grpcproto"

	"github.com/atyagi9006/certificationapp/core-service/src/models"
)

func CreateUser(client grpcproto.UserServiceClient, user *models.User) (*grpcproto.UserCreateResponse, error) {
	funcName := "CreateUser"
	log.Printf("Enter - " + funcName)
	log.Printf("Initiating GRPC  client in  :" + funcName)
	ures, err := GetUser(client, user)
	if err != nil {
		log.Fatal(err)
	}
	if ures.Email != "" {
		return nil, errors.New(user.Email + " - " + user.Type + " - user already existing failed ... ")
	}
	res, err := client.CreateUser(context.Background(), &grpcproto.User{
		UserName: user.UserName,
		Password: util.GetMD5Hash(user.Password),
		Email:    user.Email,
		Name:     user.Name,
		Type:     user.Type,
	})
	if err != nil {
		log.Println(err)
	}
	if !res.Status {
		return nil, errors.New(user.Email + " - " + user.Type + " - user creation failed ... ")
	}
	log.Println(user.UserName, " - ", user.Type, " - user created sucessful ... ")
	log.Printf("Exit - " + funcName)
	return res, nil
}

func GetUser(client grpcproto.UserServiceClient, user *models.User) (*grpcproto.User, error) {
	funcName := "GetUser"
	log.Printf("Enter - " + funcName)
	log.Printf("Initiating GRPC  client in  :" + funcName)
	res, err := client.GetUser(context.Background(), &grpcproto.User{
		UserName: user.UserName,
		Email:    user.Email,
	})
	if err != nil {
		log.Println(err)
	}
	if res == nil {
		err = errors.New(" Error while reading user... uname: " + user.UserName + " orEmail:  " + user.Email)
	}
	log.Println(res.Name, " -- res get from DB sucessful ... ")

	log.Printf("Exit - " + funcName)
	return res, err
}

func GetAllUsers(client grpcproto.UserServiceClient, requestedBy *models.User) (*grpcproto.ListUserResponse, error) {
	funcName := "GetAllUsers"
	log.Printf("Enter - " + funcName)
	log.Printf("Initiating GRPC  client in  :" + funcName)
	res, err := client.GetUserList(context.Background(), &grpcproto.User{
		UserID: requestedBy.UserID,
	})
	if err != nil {
		log.Println(err)
	}
	if res == nil {
		err = errors.New(" Error while reading All users.......")
	}

	log.Printf("Exit - " + funcName)
	return res, err

}

func CreateCandidate(client grpcproto.CandidateServiceClient, candidate *models.Candidtate) (*grpcproto.CandidateCreateResponse, error) {
	funcName := "CreateCandidate"
	log.Printf("Enter - " + funcName)
	log.Printf("Initiating GRPC  client in  :" + funcName)
	res, err := client.CreateCandidate(context.Background(), &grpcproto.Candidate{
		CandidateID: candidate.CandidtateID,
	})
	if err != nil {
		log.Println(err)
	}
	if !res.Status {
		return nil, errors.New(candidate.CandidtateID + " - candiaiate creation failed  ... ")
	}
	log.Println(candidate.CandidtateID, " -- candiate created sucessful ... ")
	log.Printf("Exit - " + funcName)
	return res, nil

}

func UpdateCandidate(client grpcproto.CandidateServiceClient, candidate *models.Candidtate) (bool, error) {
	funcName := "CreateCandidate"
	log.Printf("Enter - " + funcName)
	log.Printf("Initiating GRPC  client in  :" + funcName)
	var examAttemptList []*grpcproto.ExamAttempt
	for _, examAttempt := range candidate.ExamAttemptList {
		var queAttempt []*grpcproto.QuestionAttempts
		for _, que := range examAttempt.Questions {
			queAttempt = append(queAttempt, &grpcproto.QuestionAttempts{
				QuestionID:      que.QuestionID,
				AttemptedAnswer: que.AttemptedAnswer,
				//Category:        que.Category,
				CorrectAnswer: que.CorrectAnswer,
				/* 	Difficulty:      que.Difficulty,
				Type:            que.Type,
				Question:        que.Question,
				Option:          que.Option, */
			})
		}
		examAttemptList = append(examAttemptList, &grpcproto.ExamAttempt{
			CategoryID: examAttempt.CategoryID,
			//Result:     examAttempt.Result,
			Score:     examAttempt.Score,
			TimeSpent: examAttempt.TimeSpent,
			Questions: queAttempt,
		})
	}
	reqCandidate := grpcproto.Candidate{
		CandidateID:     candidate.CandidtateID,
		ExamAttemptList: examAttemptList,
	}

	res, err := client.UpdateCandidate(context.Background(), &reqCandidate)
	if err != nil {
		log.Println(err)
	}
	if res == nil {
		return false, errors.New(candidate.CandidtateID + " - candiaiate creation failed  ... ")
	}
	log.Println(candidate.CandidtateID, " -- candiate created sucessful ... ")
	log.Printf("Exit - " + funcName)
	return true, nil

}

func GetQuestionsAndProcess(client grpcproto.TestServiceClient, test *models.Test) ([]*models.Question, error) {
	funcName := "GetQuestionsAndProcess"
	log.Printf("Enter - " + funcName)
	log.Printf("Initiating GRPC  client in  :" + funcName)
	res, err := client.TestLaunch(context.Background(), &grpcproto.Test{
		CategoryID: test.Category,
	})
	if err != nil {
		log.Println(err)
	}
	if res == nil {
		return nil, errors.New(" Error while Launching Test No Question found ......")
	}

	cacheClient := cache.RedisClient()
	defer cacheClient.RClient.Close()
	var responsePayload []*models.Question
	for _, que := range res.QuestionList {
		cacheClient.Write(que.QuestionID, que.Answer) //putting question and answer in cache..
		que.IncorrectAnswers = append(que.IncorrectAnswers, que.Answer)
		responsePayload = append(responsePayload, &models.Question{
			QuestionID: que.QuestionID,
			Category:   que.Category,
			Type:       que.Type,
			Difficulty: que.Difficulty,
			Question:   que.Question,
			Options:    que.IncorrectAnswers,
		})
	}
	log.Printf("Exit - " + funcName)
	return responsePayload, nil
}

func GetAnswers(Questions []string) (map[string]string, error) {
	funcName := "GetAnswers"
	log.Printf("Enter - " + funcName)
	log.Printf("Getting Answer from Cache   :" + funcName)
	responsePayload := make(map[string]string)
	cacheClient := cache.RedisClient()

	for _, question := range Questions {
		log.Println("=== Key--" + question + "---value" + cacheClient.Read(question))
		responsePayload[question] = cacheClient.Read(question)
	}
	log.Printf("Exit - " + funcName)
	defer cacheClient.RClient.Close()
	return responsePayload, nil
}
func GetCandidate(cclient grpcproto.CandidateServiceClient, candidate *models.Candidtate) (*models.Candidtate, error) {
	funcName := "GetCandidate"
	log.Printf("Enter - " + funcName)
	log.Printf("Getting Candidate from DB GRPC   :" + funcName)
	res, err := cclient.GetCandidate(context.Background(), &grpcproto.Candidate{
		CandidateID: candidate.CandidtateID,
	})
	if err != nil {
		log.Println(err)
	}
	var examAttemptList []*models.ExamAttempt
	for _, examAttempt := range res.ExamAttemptList {
		/*var quesionAttempts []*models.QuestionAttempts
		for _, question := range examAttempt.Questions {
		 	quesionAttempts = append(quesionAttempts, &models.QuestionAttempts{
				QuestionID:      question.QuestionID,
				AttemptedAnswer: question.AttemptedAnswer,
				CorrectAnswer:   question.CorrectAnswer, */
		/* Category:        question.Category,
		Difficulty:      question.Difficulty,
		Question:        question.Question,
		Type:            question.Type,
		Option:          question.Option, */
		/* })
		} */
		examAttemptList = append(examAttemptList, &models.ExamAttempt{
			CategoryID: examAttempt.CategoryID,
			//Result:     examAttempt.Result,
			Score:     examAttempt.Score,
			TimeSpent: examAttempt.TimeSpent,
			//	Questions: quesionAttempts,
		})
	}

	responsePayload := &models.Candidtate{
		CandidtateID:    res.CandidateID,
		ExamAttemptList: examAttemptList,
	}

	log.Println(candidate.CandidtateID, " -- candiate created sucessful ... ")
	log.Printf("Exit - " + funcName)
	return responsePayload, nil
}
func CreateAdminIfNotExist() *models.User {
	conn := util.GetGRPCConn()
	client := grpcproto.NewUserServiceClient(conn)
	admin := models.User{
		Name:     config.AdminName,
		Email:    config.AdminEmail,
		Password: config.AdminPassword,
		Type:     config.AdminType,
	}
	ures, err := GetUser(client, &admin)
	if err != nil {
		log.Fatal(err)
	}
	if ures.Email != "" && strings.EqualFold(util.GetMD5Hash(admin.Password), ures.Password) {
		log.Println("Admin exists in db...")
		return &admin
	}
	uresponse, err := CreateUser(client, &admin)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	log.Println("Admin Created status", uresponse.Status)
	return &admin
}

package handlers

import (
	"context"
	"errors"
	"log"

	"github.com/atyagi9006/certificationapp/data-service/src/dao"
	"github.com/atyagi9006/certificationapp/data-service/src/models"
	"github.com/atyagi9006/certificationapp/grpcproto"
)

func (s *Server) CreateCandidate(ctx context.Context, req *grpcproto.Candidate) (*grpcproto.CandidateCreateResponse, error) {
	log.Printf("CreateCandidate Function Invoked:%v\n", req)
	candidateID := req.GetCandidateID()
	candidate := &models.Candidate{
		CandidateID: candidateID,
	}
	status := dao.Writer.CreateCandidate(ctx, candidate)
	return &grpcproto.CandidateCreateResponse{
		CandidateID: req.GetCandidateID(),
		Status:      status,
	}, nil
}

func (s *Server) GetCandidate(ctx context.Context, req *grpcproto.Candidate) (*grpcproto.Candidate, error) {
	log.Println("GetCandidate Function Invoked")
	candidateID := req.GetCandidateID()
	candidate := dao.Reader.GetCandidate(ctx, candidateID)
	log.Println("At readservice candidate : " + candidateID)

	var examAttemptList []*grpcproto.ExamAttempt

	for _, examAttempt := range candidate.ExamAttemptList {
		var queAttempt []*grpcproto.QuestionAttempts
		for _, que := range examAttempt.Questions {
			queAttempt = append(queAttempt, &grpcproto.QuestionAttempts{
				QuestionID:      que.QuestionID,
				AttemptedAnswer: que.AttemptedAnswer,
				CorrectAnswer:   que.CorrectAnswer,
				/* Category:        que.Category,
				Difficulty:      que.Difficulty,
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
	responseCandidate := &grpcproto.Candidate{
		CandidateID:     candidate.CandidateID,
		ExamAttemptList: examAttemptList,
	}
	return responseCandidate, nil
}

func (s *Server) UpdateCandidate(ctx context.Context, req *grpcproto.Candidate) (*grpcproto.Candidate, error) {
	log.Printf("UpdateCandidate Function Invoked:%v\n", req)
	candidateID := req.GetCandidateID()
	ExamAttemptList := req.GetExamAttemptList()
	var examAttemptList []*models.ExamAttempt
	for _, examAttempt := range ExamAttemptList {
		var quesionAttempts []*models.QuestionAttempts
		for _, question := range examAttempt.Questions {
			quesionAttempts = append(quesionAttempts, &models.QuestionAttempts{
				QuestionID:      question.QuestionID,
				AttemptedAnswer: question.AttemptedAnswer,
				CorrectAnswer:   question.CorrectAnswer,
				/* Category:        question.Category,
				Difficulty:      question.Difficulty,
				Question:        question.Question,
				Type:            question.Type,
				Option:          question.Option, */
			})
		}
		examAttemptList = append(examAttemptList, &models.ExamAttempt{
			CategoryID: examAttempt.CategoryID,
			//Result:     examAttempt.Result,
			Score:     examAttempt.Score,
			TimeSpent: examAttempt.TimeSpent,
			Questions: quesionAttempts,
		})
	}

	candidate := &models.Candidate{
		CandidateID:     candidateID,
		ExamAttemptList: examAttemptList,
	}

	status := dao.Writer.UpdateCandidate(ctx, candidateID, candidate)
	if status {
		return req, nil
	}
	return nil, errors.New("Error while Updating Candidate")
}
func (s *Server) GetCandidateList(ctx context.Context, in *grpcproto.Candidate) (*grpcproto.ListCandidateResponse, error) {

	return nil, nil
}

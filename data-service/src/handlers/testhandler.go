package handlers

import (
	"context"
	"log"
	"strconv"

	"github.com/atyagi9006/certificationapp/data-service/src/dao"
	"github.com/atyagi9006/certificationapp/data-service/src/models"
	"github.com/atyagi9006/certificationapp/grpcproto"
)

func (s *Server) CreateTest(ctx context.Context, req *grpcproto.Test) (*grpcproto.Test, error) {
	return nil, nil
}
func (s *Server) GetTestList(ctx context.Context, req *grpcproto.Test) (*grpcproto.ListTestResponse, error) {
	return nil, nil
}
func (s *Server) TestLaunch(ctx context.Context, req *grpcproto.Test) (*grpcproto.TestLaunchResponse, error) {
	log.Println("DATA -SERVICE TestLaunch Function Invoked")
	test := &models.Test{
		CategoryID: req.CategoryID,
	}
	rques := dao.Reader.TestLaunch(ctx, *test)
	log.Println("At readservice got users: " + strconv.Itoa(len(rques)))

	var qSlice []*grpcproto.Question
	for _, rque := range rques {
		qSlice = append(qSlice, &grpcproto.Question{
			QuestionID:       rque.QuestionID,
			Category:         rque.Category,
			Difficulty:       rque.Difficulty,
			Question:         rque.Question,
			Type:             rque.Type,
			Answer:           rque.Answer,
			IncorrectAnswers: rque.IncorrectAnswers,
		})
	}

	return &grpcproto.TestLaunchResponse{
		QuestionList: qSlice,
	}, nil
}

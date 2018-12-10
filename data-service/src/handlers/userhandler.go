package handlers

import (
	"context"
	"log"
	"strconv"

	"github.com/atyagi9006/certificationapp/data-service/src/dao"
	"github.com/atyagi9006/certificationapp/data-service/src/models"
	grpcproto "github.com/atyagi9006/certificationapp/grpcproto"
)

func (s *Server) CreateUser(ctx context.Context, req *grpcproto.User) (*grpcproto.UserCreateResponse, error) {
	log.Println("CreateUser Function Invoked")
	userName := req.GetUserName()
	password := req.GetPassword()
	Type := req.GetType()
	email := req.GetEmail()
	name := req.GetName()
	user := &models.User{
		UserName: userName,
		Password: password,
		Type:     Type,
		Email:    email,
		Name:     name,
	}
	status := dao.Writer.CreateUser(ctx, user)
	return &grpcproto.UserCreateResponse{
		UserID: user.UserID,
		Status: status,
	}, nil
}

func (s *Server) GetUser(ctx context.Context, req *grpcproto.User) (*grpcproto.User, error) {

	log.Println("GetUser Function Invoked")
	userName := req.GetUserName()
	email := req.GetEmail()
	user := &models.User{
		UserName: userName,
		Email:    email,
	}
	ruser := dao.Reader.GetUser(ctx, *user)
	log.Println("At readservice user: " + ruser.Name)
	return &grpcproto.User{
		UserID:   ruser.UserID,
		UserName: ruser.UserName,
		Name:     ruser.Name,
		Password: ruser.Password,
		Email:    ruser.Email,
		Type:     ruser.Type,
	}, nil
}

func (s *Server) UpdateUser(ctx context.Context, req *grpcproto.User) (*grpcproto.User, error) {
	return nil, nil
}
func (s *Server) GetUserList(ctx context.Context, req *grpcproto.User) (*grpcproto.ListUserResponse, error) {
	log.Println("GetUserList Function Invoked")
	userId := req.GetUserID()
	user := &models.User{
		UserID: userId,
	}
	rusers := dao.Reader.GetAllUsers(ctx, *user)
	log.Println("At readservice got users: " + strconv.Itoa(len(rusers)))

	var uSlice []*grpcproto.User
	for _, ruser := range rusers {
		uSlice = append(uSlice, &grpcproto.User{
			UserID:   ruser.UserID,
			UserName: ruser.UserName,
			Name:     ruser.Name,
			Password: ruser.Password,
			Email:    ruser.Email,
			Type:     ruser.Type,
		})
	}

	return &grpcproto.ListUserResponse{
		Users: uSlice,
	}, nil

}

package dao

import (
	"context"

	"github.com/atyagi9006/certificationapp/data-service/src/models"
)

type WriteDataInterface interface {
	CreateUser(ctx context.Context, user *models.User) bool
	UpdateUser(ctx context.Context, user *models.User) bool
	DeleteUser(ctx context.Context, userName string) bool
	CreateCandidate(ctx context.Context, candidate *models.Candidate) bool
	UpdateCandidate(ctx context.Context, candidateID string, candidate *models.Candidate) bool
	DeleteCandidate(ctx context.Context, candidateID string) bool
	
	
}

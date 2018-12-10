package dao

import (
	"context"

	"github.com/atyagi9006/certificationapp/data-service/src/models"
)

//ReadDataInterface : Provides Read Functionalities
type ReadDataInterface interface {
	GetAdmin(ctx context.Context, adminID string) *models.Admin
	GetCandidate(ctx context.Context, candidateID string) *models.Candidate
	GetAllCandidates(ctx context.Context) []*models.Candidate
	GetUser(ctx context.Context, user models.User) *models.User
	GetAllUsers(ctx context.Context, requestedBy models.User) []*models.User
	TestLaunch(ctx context.Context, Test models.Test) []*models.Question
}

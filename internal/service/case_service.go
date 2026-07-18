package service

import (
	"context"

	"github.com/Zumbr3/case-service/internal/adapter/repository/models"
	"github.com/Zumbr3/case-service/internal/contracts"
)

type CaseService struct {
	repo contracts.CaseRepository
}

func NewCaseService(repo contracts.CaseRepository) *CaseService {
	return &CaseService{repo: repo}
}

func (s *CaseService) GetAllCases(ctx context.Context) ([]*models.Case, error) {
	return s.repo.GetAllCases(ctx)
}

func (s *CaseService) GetCaseByID(ctx context.Context, id string) (*models.Case, error) {
	return s.repo.GetCaseById(ctx, id)
}

func (s *CaseService) CreateCase(ctx context.Context, c *models.Case) (*models.Case, error) {
	return s.repo.CreateCase(ctx, c)
}

func (s *CaseService) UpdateCase(ctx context.Context, c *models.Case) (*models.Case, error) {
	return s.repo.UpdateCase(ctx, c)
}

func (s *CaseService) DeleteCaseByID(ctx context.Context, id string) error {
	return s.repo.DeleteCaseByID(ctx, id)
}

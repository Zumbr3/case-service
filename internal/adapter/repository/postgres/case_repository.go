package postgres

import (
	"context"

	"github.com/Zumbr3/case-service/internal/adapter/repository/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

type CaseRepository struct {
	db *pgxpool.Pool
}

func NewCaseRepository(db *pgxpool.Pool) *CaseRepository {
	return &CaseRepository{db: db}
}

func (r *CaseRepository) GetAllCases(ctx context.Context) ([]*models.Case, error) {
	return []*models.Case{}, nil
}

func (r *CaseRepository) GetCaseById(ctx context.Context, id string) (*models.Case, error) {
	return &models.Case{}, nil
}

func (r *CaseRepository) CreateCase(ctx context.Context, c *models.Case) (*models.Case, error) {
	return &models.Case{}, nil
}

func (r *CaseRepository) UpdateCase(ctx context.Context, c *models.Case) (*models.Case, error) {
	return &models.Case{}, nil
}

func (r *CaseRepository) DeleteCaseByID(ctx context.Context, id string) error {
	return nil
}

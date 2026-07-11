package contracts

import (
	"context"

	"github.com/Zumbr3/case-service/internal/models"
)

type CaseRepository interface {
	GetAllCases(ctx context.Context) ([]*models.Case, error)
	GetCaseById(ctx context.Context, id string) (*models.Case, error)
	CreateCase(ctx context.Context, c *models.Case) (*models.Case, error)
	UpdateCase(ctx context.Context, c *models.Case) (*models.Case, error)
	DeleteCaseByID(ctx context.Context, id string) error
}

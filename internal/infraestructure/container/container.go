package container

import (
	"github.com/Zumbr3/case-service/internal/adapter/repository/postgres"
	"github.com/Zumbr3/case-service/internal/contracts"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Container struct {
	CaseRepository contracts.CaseRepository
}

func New(db *pgxpool.Pool) *Container {
	return &Container{
		CaseRepository: postgres.NewCaseRepository(db),
	}
}

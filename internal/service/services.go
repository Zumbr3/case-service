package service

import "github.com/Zumbr3/case-service/internal/contracts"

// Services aggregates every domain service so callers only ever take one argument,
// regardless of how many services get added.
type Services struct {
	Case contracts.CaseService
}

func NewServices(caseRepo contracts.CaseRepository) *Services {
	return &Services{
		Case: NewCaseService(caseRepo),
	}
}

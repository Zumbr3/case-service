package handler

import "github.com/Zumbr3/case-service/internal/service"

// Handlers aggregates every domain handler so router only ever takes one argument,
// regardless of how many services/handlers get added.
type Handlers struct {
	Case *CaseHandler
}

func NewHandlers(s *service.Services) *Handlers {
	return &Handlers{
		Case: NewCaseHandler(s.Case),
	}
}

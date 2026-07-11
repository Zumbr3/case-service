package models

import (
	"time"

	"github.com/Zumbr3/case-service/internal/domain"
)

// Case is the persistence representation of a case, mirroring the `cases` table.
// Id values are uuid.NewV6(), so they are treated as string.
type Case struct {
	ID             string                 `json:"id" db:"id"`
	TransactionID  string                 `json:"transactionId" db:"transaction_id"`
	AccountID      string                 `json:"accountId" db:"account_id"`
	Score          float64                `json:"score" db:"score"`
	Origin         domain.Origin          `json:"origin" db:"origin"`
	Status         domain.Status          `json:"status" db:"status"`
	TriggeredRules []domain.TriggeredRule `json:"triggeredRules" db:"triggered_rules"`
	SLADeadline    time.Time              `json:"slaDeadline" db:"sla_deadline"`
	AnalystID      string                 `json:"analystId,omitempty" db:"analyst_id"`
	DecisionNotes  string                 `json:"decisionNotes,omitempty" db:"decision_notes"`
	DecidedAt      *time.Time             `json:"decidedAt,omitempty" db:"decided_at"`
	CreatedAt      time.Time              `json:"createdAt" db:"created_at"`
}

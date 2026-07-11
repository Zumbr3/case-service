package domain

import (
	"errors"
	"log/slog"
	"time"

	domainerrors "github.com/Zumbr3/case-service/internal/errors"
)

const (
	StatusHeld     Status = "held"
	StatusReleased Status = "released"
	StatusReversed Status = "reversed"
)

const (
	OriginManualReview Origin = "manual_review"
	OriginAutoBlock    Origin = "auto_block"
)

type Status string
type Origin string

type TriggeredRule struct {
	RuleID       string  `json:"ruleID"`
	Name         string  `json:"name"`
	PartialScore float64 `json:"partialScore"`
	Weight       string  `json:"weight"`
}

// Id values are uuid.NewV6(), so they are treated as string
type Case struct {
	TransactionID string        `json:"transactionId"`
	AccountID     string        `json:"accountId"`
	Score         float64       `json:"score"`
	Origin        Origin        `json:"origin"`
	Status        Status        `json:"status"`
	Triggered     TriggeredRule `json:"triggered"`
	SLADeadline   time.Time     `json:"slaDeadline"`
	AnalystID     string        `json:"analystId"`
	DecisionNotes string        `json:"decisionNotes"`
	DecidedAt     time.Time     `json:"decidedAt"`
	CreatedAt     time.Time     `json:"createdAt"`
}

func NewCase(transactionId, accountId, origin, status,
	ruleID, name, weight string,
	score, partialScore float64, slaDeadline, createdAt time.Time) (Case, error) {

	parsedStatus, err := parseStatus(status)
	if err != nil {
		return Case{}, err
	}

	parsedOrigin, err := parseOrigin(origin)
	if err != nil {
		return Case{}, err
	}

	if err := validateScore(score); err != nil {
		return Case{}, err
	}

	if err := validateScore(partialScore); err != nil {
		return Case{}, err
	}

	if slaDeadline.IsZero() {
		slog.Error("Invalid SLA deadline", "slaDeadline", slaDeadline)
		return Case{}, errors.New("slaDeadline is required")
	}

	triggered := TriggeredRule{
		RuleID:       ruleID,
		Name:         name,
		PartialScore: partialScore,
		Weight:       weight,
	}

	return Case{
		TransactionID: transactionId,
		AccountID:     accountId,
		Score:         score,
		Origin:        parsedOrigin,
		Status:        parsedStatus,
		Triggered:     triggered,
		SLADeadline:   slaDeadline,
		CreatedAt:     createdAt,
	}, nil
}

func validateScore(score float64) error {
	if score < 0 {
		slog.Error("Invalid score", "score", score)
		return errors.New("score must not be negative")
	}
	return nil
}

func validateStatus(s string) error {
	switch s {
	case string(StatusHeld), string(StatusReleased), string(StatusReversed):
		return nil
	default:
		slog.Error("Invalid status", "status", s)
		return errors.New("Invalid status")
	}
}

func parseStatus(s string) (Status, error) {
	if err := validateStatus(s); err != nil {
		return "", err
	}
	return Status(s), nil
}

func validateOrigin(o string) error {
	switch o {
	case string(OriginManualReview), string(OriginAutoBlock):
		return nil
	default:
		slog.Error("Invalid origin", "origin", o)
		return errors.New("Invalid origin")
	}
}

func parseOrigin(o string) (Origin, error) {
	if err := validateOrigin(o); err != nil {
		return "", err
	}
	return Origin(o), nil
}

func (c *Case) CanBeDecided() bool {
	return c.Status == StatusHeld
}

func (c *Case) ApplyDecision(decision Decision, analystID, notes string, now time.Time) error {
	if !c.CanBeDecided() {
		slog.Error("Case already decided", "transactionId", c.TransactionID, "status", c.Status)
		return domainerrors.ErrCaseAlreadyDecided
	}

	c.Status = decision.toStatus()
	c.AnalystID = analystID
	c.DecisionNotes = notes
	c.DecidedAt = now
	return nil
}

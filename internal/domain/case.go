package domain

import (
	"errors"
	"log/slog"
	"time"

	"github.com/Zumbr3/case-service/internal/dto"
)

const (
	StatusHeld     dto.Status = "held"
	StatusReleased dto.Status = "released"
	StatusReversed dto.Status = "reversed"
)

const (
	OriginManualReview dto.Origin = "manual_review"
	OriginAutoBlock    dto.Origin = "auto_block"
)

func NewCase(transactionId, accountId, origin, status,
	ruleID, name, weight, analystId, decisionNotes string,
	score, partialScore float64, slaDeadline,
	decisionAt, createdAt time.Time) (dto.Case, error) {

	parsedStatus, err := parseStatus(status)
	if err != nil {
		return dto.Case{}, err
	}

	parsedOrigin, err := parseOrigin(origin)
	if err != nil {
		return dto.Case{}, err
	}

	if err := validateScore(score); err != nil {
		return dto.Case{}, err
	}

	if err := validateScore(partialScore); err != nil {
		return dto.Case{}, err
	}

	if slaDeadline.IsZero() {
		slog.Error("Invalid SLA deadline", "slaDeadline", slaDeadline)
		return dto.Case{}, errors.New("slaDeadline is required")
	}

	triggered := dto.TriggeredRule{
		RuleID:       ruleID,
		Name:         name,
		PartialScore: partialScore,
		Weight:       weight,
	}

	return dto.NewCase(
		transactionId,
		accountId,
		score,
		parsedOrigin,
		parsedStatus,
		triggered,
		slaDeadline,
		analystId,
		decisionNotes,
		decisionAt,
		createdAt,
	), nil
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

func parseStatus(s string) (dto.Status, error) {
	if err := validateStatus(s); err != nil {
		return "", err
	}
	return dto.Status(s), nil
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

func parseOrigin(o string) (dto.Origin, error) {
	if err := validateOrigin(o); err != nil {
		return "", err
	}
	return dto.Origin(o), nil
}

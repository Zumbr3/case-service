package domain

import (
	"log/slog"

	decisionerrors "github.com/Zumbr3/case-service/internal/errors"
)

type Decision string

const (
	decisionReleased Decision = "released"
	decisionReversed Decision = "reversed"
)

func ReturnDecision(val string) (Decision, error) {
	d := Decision(val)

	if err := validateDecision(d); err != nil {
		return "", err
	}

	return d, nil
}

func (d Decision) toStatus() Status {
	switch d {
	case decisionReleased:
		return StatusReleased
	case decisionReversed:
		return StatusReversed
	default:
		return ""
	}
}

func validateDecision(d Decision) error {
	switch d {
	case decisionReleased, decisionReversed:
		return nil
	default:
		slog.Error("Invalid decision", "decision", d)
		return decisionerrors.ErrDecisionInvalid
	}
}

package domain

import (
	"errors"
	"log/slog"
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

func validateDecision(d Decision) error {
	switch d {
	case decisionReleased, decisionReversed:
		return nil
	default:
		slog.Error("Invalid decision", "decision", d)
		return errors.New("Invalid decision")
	}
}

package dto

import "time"

type Status string
type Origin string

type TriggeredRule struct {
	RuleID       string
	Name         string
	PartialScore float64
	Weight       string
}

// Id values are uuid.NewV6(), so they are treated as string
type Case struct {
	transactionId string
	accountId     string
	score         float64
	origin        Origin
	status        Status
	triggered     TriggeredRule
	slaDeadline   time.Time
	analystId     string
	decisionNotes string
	decisionAt    time.Time
	createdAt     time.Time
}

func NewCase(transactionId, accountId string, score float64, origin Origin, status Status,
	triggered TriggeredRule, slaDeadline time.Time, analystId, decisionNotes string,
	decisionAt, createdAt time.Time) Case {

	return Case{
		transactionId: transactionId,
		accountId:     accountId,
		score:         score,
		origin:        origin,
		status:        status,
		triggered:     triggered,
		slaDeadline:   slaDeadline,
		analystId:     analystId,
		decisionNotes: decisionNotes,
		decisionAt:    decisionAt,
		createdAt:     createdAt,
	}
}

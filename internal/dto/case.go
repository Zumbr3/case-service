package dto

type CaseID struct {
	CaseID string `uri:"id" json:"caseId" binding:"required"`
}

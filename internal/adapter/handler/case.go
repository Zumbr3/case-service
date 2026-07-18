package handler

import (
	"errors"
	"net/http"

	"github.com/Zumbr3/case-service/internal/contracts"
	"github.com/Zumbr3/case-service/internal/dto"
	repositoryerrors "github.com/Zumbr3/case-service/internal/errors"
	"github.com/gin-gonic/gin"
)

type CaseHandler struct {
	service contracts.CaseService
}

func NewCaseHandler(service contracts.CaseService) *CaseHandler {
	return &CaseHandler{service: service}
}

func (h *CaseHandler) GetCase(c *gin.Context) {
	var id dto.CaseID
	if err := c.ShouldBindUri(&id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid case id", "details": err.Error()})
		return
	}

	caseResult, err := h.service.GetCaseByID(c.Request.Context(), id.CaseID)
	if err != nil {
		if errors.Is(err, repositoryerrors.ErrCaseNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, caseResult)
}

func (h *CaseHandler) GetAllCases(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

func (h *CaseHandler) CreateCase(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

func (h *CaseHandler) UpdateCase(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

func (h *CaseHandler) DeleteCaseByID(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

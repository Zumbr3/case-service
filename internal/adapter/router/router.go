package router

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/Zumbr3/case-service/internal/adapter/handler"
	"github.com/gin-gonic/gin"
)

func Router(h *handler.Handlers) *http.Server {
	router := gin.Default()

	api := router.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			cases := v1.Group("/cases")
			{
				cases.GET("/:id", h.Case.GetCase)
				cases.GET("/", h.Case.GetAllCases)
				cases.POST("/", h.Case.CreateCase)
				cases.PUT("/:id", h.Case.UpdateCase)
				cases.DELETE("/:id", h.Case.DeleteCaseByID)
			}
		}
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	return &http.Server{
		Addr:              fmt.Sprintf(":%s", port),
		Handler:           router,
		ReadHeaderTimeout: 5 * time.Second,
	}
}

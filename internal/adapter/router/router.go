package router

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/Zumbr3/case-service/internal/adapter/handler"
	"github.com/gin-gonic/gin"
)

func Router() *http.Server {
	router := gin.Default()

	api := router.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			cases := v1.Group("/cases")
			{
				cases.GET("/", handler.GetCase)
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

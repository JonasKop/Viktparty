package main

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"log"
	"net/http"
	"os"
	"server/src/database"
	"server/src/handlers"
	"server/src/lib"
	"server/src/middlewares"
)

func makeRoutes(r *gin.Engine) {
	weight := r.Group("/weight")
	{
		weight.POST("", handlers.InsertWeight)
		today := weight.Group("/today")
		{
			today.DELETE("", handlers.DeleteTodaysWeight)
			today.GET("", handlers.GetTodaysWeight)
		}
	}
}

func getOIDCConfig() (*lib.OIDCConfig, error) {
	oidcConfig := lib.OIDCConfig{
		Issuer:   os.Getenv("OIDC_ISSUER"),
		Audience: os.Getenv("OIDC_AUDIENCE"),
	}

	if oidcConfig.Issuer == "" {
		return nil, errors.New("Missing configuration option \"OIDC_ISSUER\"")
	} else if oidcConfig.Audience == "" {
		return nil, errors.New("Missing configuration option \"OIDC_AUDIENCE\"")
	}
	return &oidcConfig, nil
}

func main() {
	db := database.ConnectToDb()

	oidcConfig, err := getOIDCConfig()
	if err != nil {
		log.Fatalln(errors.Wrap(err, "Cannot get OIDC config"))
	}

	r := gin.Default()
	r.GET("/healthz", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"healthy": true})
	})

	r.Use(middlewares.PanicRecoverer())
	r.Use(middlewares.AuthChecker(oidcConfig))
	r.Use(middlewares.DBContext(db))

	makeRoutes(r)

	port := "localhost:8080"
	if os.Getenv("GO_ENV") == "production" {
		port = ":8080"
	}
	err = r.Run(port)
	if err != nil {
		log.Fatalln("Could not start server on localhost:8080")
	}
}

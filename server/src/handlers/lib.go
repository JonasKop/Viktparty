package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"log"
)

func getNameFromContext(c *gin.Context) string {
	value, exists := c.Get("name")
	if !exists {
		log.Panicln("Cannot get name from context")
	}
	return value.(string)
}

func getUserIDFromContext(c *gin.Context) string {
	value, exists := c.Get("userID")
	if !exists {
		log.Panicln("Cannot get user id from context")
	}
	return value.(string)
}

func getDatabaseFromContext(c *gin.Context) *sqlx.DB {
	value, exists := c.Get("db")
	if !exists {
		log.Panicln("Cannot get database from context")
	}
	return value.(*sqlx.DB)
}

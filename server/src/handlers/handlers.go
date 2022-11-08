package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/src/database"
	"server/src/lib"
)

func InsertWeight(c *gin.Context) {
	userID := getUserIDFromContext(c)
	name := getNameFromContext(c)
	db := getDatabaseFromContext(c)

	var weight lib.NewWeight
	if err := c.BindJSON(&weight); err != nil {
		lib.ErrorMessage(c, err, http.StatusBadRequest, "Invalid JSON body")
		return
	}

	if weight.Weight <= 0 {
		lib.ErrorMessage(c, nil, http.StatusBadRequest, "Weight must be greater than zero")
		return
	}

	_, err := database.GetTodaysWeight(db, userID)
	if err == nil {
		lib.ErrorMessage(c, err, http.StatusBadRequest, "You have already checked in")
		return
	}

	err = database.InsertWeight(db, userID, name, weight.Weight)
	if err != nil {
		lib.ErrorMessage(c, err, http.StatusBadRequest, "Could not insert todays weight")
		return
	}
	c.JSON(http.StatusNoContent, nil)
}

func DeleteTodaysWeight(c *gin.Context) {
	userID := getUserIDFromContext(c)
	db := getDatabaseFromContext(c)
	err := database.DeleteTodaysWeight(db, userID)
	if err != nil {
		lib.ErrorMessage(c, err, http.StatusBadRequest, "Could not delete todays weight")
		return
	}
	c.JSON(http.StatusNoContent, nil)

	println("get some weights")
}

func GetTodaysWeight(c *gin.Context) {
	userID := getUserIDFromContext(c)
	db := getDatabaseFromContext(c)
	weight, err := database.GetTodaysWeight(db, userID)
	if err != nil {
		lib.ErrorMessage(c, err, http.StatusBadRequest, "Could not get todays weight")
		return
	}
	c.JSON(http.StatusOK, lib.Weight{
		UserID: userID,
		Weight: weight,
	})

	println("get some weights")
}

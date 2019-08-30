package handler

import (
	"github.com/AnotherCoolDude/proad/proadclient"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetEmployeesHandler returns all active proad employees
func GetEmployeesHandler(c *gin.Context) {
	emps := proadclient.GetEmployees()
	if len(emps.Employees) < 1 {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not get employees"})
		return
	}
	c.JSON(http.StatusOK, emps)
}

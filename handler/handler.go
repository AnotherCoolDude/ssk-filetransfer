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

// GetProjectsHandler returns a projectlist for the employee identified by the parameter urno
func GetProjectsHandler(c *gin.Context) {
	pp := proadclient.GetProjectsForEmployee(c.Param("urno"), proadclient.StatusNone)
	if len(pp.Projects) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "projectList is empty"})
	}
	c.JSON(http.StatusOK, pp)
}

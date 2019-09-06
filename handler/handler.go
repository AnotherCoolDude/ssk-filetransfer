package handler

import (
	"fmt"
	"github.com/AnotherCoolDude/ssk-filetransfer/filemanagement"
	"net/http"
	"strconv"
	"time"

	"github.com/AnotherCoolDude/ssk-filetransfer/proadclient"
	"github.com/gin-gonic/gin"
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

// GetProjectsByEmployeeHandler returns a projectlist for the employee identified by the parameter urno
func GetProjectsByEmployeeHandler(c *gin.Context) {
	pp := proadclient.GetProjectsForEmployee(c.Param("urno"), proadclient.StatusNone)
	if len(pp.Projects) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "projectList is empty"})
	}
	c.JSON(http.StatusOK, pp)
}

// GetFilteredProjects returns projects filtered by query params
func GetFilteredProjects(c *gin.Context) {
	pp := proadclient.GetProjects(extractFilterParameter(c))
	if len(pp.Projects) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "projectList is empty"})
	}
	c.JSON(http.StatusOK, pp)
}

// GetContentForProject returns the contents for a project
func GetContentForProject(c *gin.Context) {
	prnr := c.Query("projectnr")
	if prnr == "" {
		fmt.Println("couldn't extract Projectnumber")
		c.JSON(http.StatusBadRequest, gin.H{"error": "couldn't extract project number"})
	}
	paths := filemanagement.FindPathsForProject(filemanagement.PathDaten, prnr)
	files := filemanagement.ListContentForPaths(paths)
	c.JSON(http.StatusOK, files)
}

// GetFilteredTasks returns projects filtered by query params
func GetFilteredTasks(c *gin.Context) {
	tt := proadclient.GetTasks(extractFilterParameter(c))
	if len(tt.Tasks) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "taskList is empty"})
	}
	c.JSON(http.StatusOK, tt)
}

func extractFilterParameter(c *gin.Context) (code proadclient.StatusCode, startDate, endDate time.Time) {
	codeNr, err := strconv.Atoi(c.Query("status"))
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	code = proadclient.Code(codeNr)
	startDate, err = time.Parse(time.RFC3339, c.Query("startDate"))
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	endDate, err = time.Parse(time.RFC3339, c.Query("endDate"))
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	return code, startDate, endDate
}

package proadclient

import (
	"strconv"
	"time"
)

type queryPair struct {
	key, value string
}

const (
	proadTimeFormat = "2006-01-02"
)

// GetEmployees returns a map where the key is the emplyees shortname and the values his/her urno
func GetEmployees() EmployeeList {
	var el EmployeeList
	req := makeGETRequest("staffs", map[string]string{"active": "1"})
	unmarshalResponse(req, &el)
	return el
}

// GetProjectsForEmployee returns a ProjectList for emplyee specified by urno
func GetProjectsForEmployee(urno string, status StatusCode) *ProjectList {
	var pl ProjectList
	req := makeGETRequest("staffs/"+urno+"/projects", queryMap(status))
	unmarshalResponse(req, &pl)
	return &pl
}

// GetTasksForEmployee returns a list of Tasks für employee with identifier urno
func GetTasksForEmployee(urno string, status StatusCode) *TaskList {
	var tl TaskList
	req := makeGETRequest("staffs/"+urno+"/tasks", queryMap(status))
	unmarshalResponse(req, &tl)
	return &tl
}

// GetTasks returns tasks filtered by status and date
func GetTasks(code StatusCode, startDate, endDate time.Time) *TaskList {
	var tl TaskList
	req := makeFilteredGetRequest("tasks", code, startDate, endDate)
	unmarshalResponse(req, &tl)
	return &tl
}

// GetProjects returns projects filtered by status and date
func GetProjects(code StatusCode, startDate, endDate time.Time) *ProjectList {
	var pl ProjectList
	req := makeFilteredGetRequest("projects", code, startDate, endDate)
	unmarshalResponse(req, &pl)
	return &pl
}

// GetProject returns a projectlist with a single project defined by projectnr
func GetProject(projectnr string) *ProjectList {
	var pl ProjectList
	req := makeGETRequest("projects", queryMap(StatusNone, queryPair{key: "projectno", value: projectnr}))
	unmarshalResponse(req, &pl)
	return &pl
}

// GetTasksForProject returns a tasklist for the given project
func GetTasksForProject(project *Project) *TaskList {
	var tl TaskList
	req := makeGETRequest("tasks", queryMap(StatusNone, queryPair{key: "project", value: strconv.Itoa(project.Urno)}))
	unmarshalResponse(req, &tl)
	return &tl
}

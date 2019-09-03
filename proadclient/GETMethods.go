package proadclient

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
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

func makeFilteredGetRequest(path string, code StatusCode, startDate, endDate time.Time) *http.Request {
	q := queryPair{
		key:   "from_date",
		value: startDate.Format(proadTimeFormat) + "-" + endDate.Format(proadTimeFormat),
	}
	return makeGETRequest(path, queryMap(code, q))
}

func queryMap(code StatusCode, pairs ...queryPair) map[string]string {
	sMap := map[string]string{}
	if code != StatusNone {
		sMap["status"] = code.String()
	}
	for _, qP := range pairs {
		sMap[qP.key] = qP.value
	}
	return sMap
}

func unmarshalResponse(request *http.Request, v interface{}) {
	resp := Client.Do(request)
	respBodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("[proadclient/emplyees.go] error: %s\n", err.Error())
		return
	}
	defer resp.Body.Close()

	err = json.Unmarshal(respBodyBytes, &v)
	if err != nil {
		fmt.Printf("[proadclient/emplyees.go] error: %s\n", err.Error())
		return
	}
}

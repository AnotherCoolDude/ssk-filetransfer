package proadclient

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
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
	statusMap := map[string]string{}
	if status != StatusNone {
		statusMap["status"] = status.String()
	}
	req := makeGETRequest("staffs/"+urno+"/projects", statusMap)
	unmarshalResponse(req, &pl)
	return &pl
}

// GetTasksForEmployee returns a list of Tasks für employee with identifier urno
func GetTasksForEmployee(urno string, status StatusCode) *TaskList {
	var tl TaskList
	statusMap := map[string]string{}
	if status != StatusNone {
		statusMap["status"] = status.String()
	}
	req := makeGETRequest("staffs/"+urno+"/tasks", statusMap)
	unmarshalResponse(req, &tl)
	return &tl
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

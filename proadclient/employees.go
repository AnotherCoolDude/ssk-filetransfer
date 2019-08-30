package proadclient

// EmployeeList represents a list of employees
type EmployeeList struct {
	Employees []Employee `json:"person_list"`
}

// Employee represents an (active) employee
type Employee struct {
	Urno       int    `json:"urno"`
	Shortname  string `json:"shortname"`
	Firstname  string `json:"firstname"`
	Lastname   string `json:"lastname"`
	Department string `json:"department"`
	Active     bool   `json:"active"`
}

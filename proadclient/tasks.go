package proadclient

// TaskList represents a number of Tasks in Proad
type TaskList struct {
	Tasks []Task `json:"todo_list"`
}

// Task represents a task in Proad
type Task struct {
	Urno             int    `json:"urno"`
	UrnoManager      int    `json:"urno_manager"`
	UrnoCompany      int    `json:"urno_company"`
	UrnoProject      int    `json:"urno_project"`
	UrnoServiceCode  int    `json:"urno_service_code"`
	UrnoResponsible  int    `json:"urno_responsible"`
	Shortinfo        string `json:"shortinfo"`
	FromDatetime     string `json:"from_datetime"`
	UntilDatetime    string `json:"until_datetime"`
	ReminderDatetime string `json:"reminder_datetime"`
	Status           string `json:"status"`
	Priority         string `json:"priority"`
	Description      string `json:"description"`
}

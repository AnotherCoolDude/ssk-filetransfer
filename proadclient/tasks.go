package proadclient

// TaskList represents a number of Tasks in Proad
type TaskList struct {
	Tasks []Task `json:"todo_list"`
}

// Task represents a task in Proad
type Task struct {
	// "urno": 0,
	//     "shortinfo": "string",
	//     "from_datetime": "string",
	//     "until_datetime": "string",
	//     "reminder_datetime": "string",
	//     "status": "string",
	//     "priority": "string",
	//     "description": "string"
	Urno        int    `json:"urno"`
	Info        string `json:"shortinfo"`
	From        string `json:"from_datetime"`
	Until       string `json:"until_datetime"`
	Reminder    string `json:"reminder_datetime"`
	Status      string `json:"status"`
	Description string `json:"description"`
}

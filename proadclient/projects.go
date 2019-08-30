package proadclient

// ProjectList represents an array of projects
type ProjectList struct {
	Projects []Project `json:"project_list"`
}

// Project represents a single project
type Project struct {
	Urno        int    `json:"urno"`
	Status      string `json:"status"`
	Title       string `json:"project_name"`
	Jobnr       string `json:"projectno"`
	Description string `json:"description"`
}

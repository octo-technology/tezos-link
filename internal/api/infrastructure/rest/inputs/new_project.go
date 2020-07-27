package inputs

// NewProject is a new project sent to the api
type NewProject struct {
	Title   string `json:"title"`
	Network string `json:"network"`
}

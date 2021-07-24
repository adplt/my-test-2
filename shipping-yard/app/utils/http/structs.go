package http

// CommonJSON struct
type CommonJSON struct {
	Status     string      `json:"status"`
	Data       interface{} `json:"data"`
	Message    string      `json:"message"`
	StatusCode int
}

package presentase

type Meta struct {
	Total string `json:"total"`
	Page  string `json:"page"`
	Limit string `json:"limit"`
}

type GetAllResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

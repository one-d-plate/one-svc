package presentase

type Meta struct {
	Limit  string `json:"limit,omitempty"`
	Cursor string `json:"cursor,omitempty"`
}

type GetAllResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

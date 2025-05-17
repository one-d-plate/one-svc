package presentase

type GetAllHeader struct {
	Limit  int    `json:"limit,omitempty"`
	Search string `json:"search,omitempty"`
	Cursor string `json:"cursor,omitempty"`
}

type DeleteRequest struct {
	ID      []int `json:"id"`
	Include bool  `json:"include,omitempty"`
}

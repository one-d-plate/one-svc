package presentase

type GetAllHeader struct {
	Limit  int    `json:"limit,omitempty"`
	Search string `json:"search,omitempty"`
	Cursor string `json:"cursor,omitempty"`
}

package presentase

type GetAllHeader struct {
	Page   int    `json:"page,omitempty"`
	Limit  int    `json:"limit,omitempty"`
	Search string `json:"search,omitempty"`
}

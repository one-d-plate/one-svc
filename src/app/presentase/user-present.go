package presentase

import "github.com/one-d-plate/one-svc.git/src/app/entity"

type CreateUserReq struct {
	Username string `json:"username"`
	Nama     string `json:"nama"`
	Email    string `json:"email"`
	Hp       string `json:"hp"`
}

type GetUsersResponse struct {
	List []entity.User `json:"list"`
	Meta Meta          `json:"meta"`
}

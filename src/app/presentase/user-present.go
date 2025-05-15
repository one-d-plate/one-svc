package presentase

import "github.com/one-d-plate/one-svc.git/src/app/entity"

type GetUserResponse struct {
	List []entity.User `json:"list"`
	Meta Meta          `json:"meta"`
}

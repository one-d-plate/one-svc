package entity

type Status string

var (
	UserInActive Status = "00"
	UserActive   Status = "01"
	UserPending  Status = "02"
	UserBanned   Status = "09"
)

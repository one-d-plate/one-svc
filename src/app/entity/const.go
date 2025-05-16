package entity

type Status string
type StrStatus map[string]Status

const (
	UserInActive Status = "00"
	UserActive   Status = "01"
	UserPending  Status = "02"
	UserBanned   Status = "09"
)

var UserStatus = StrStatus{
	string(UserInActive): "INACTIVE",
	string(UserActive):   "ACTIVE",
	string(UserPending):  "PENDING",
	string(UserBanned):   "BANNED",
}

package models

import (
	"strconv"
)

type User_role struct {
	ID     uint   `gorm:"primary_key" json:"id"`
	UserId string `json:"user_id"`
	RoleId uint   `json:"role_id"`
}

func FindRoleByUserId(userId interface{}) User_role {
	var uRole User_role
	OldDB.Where("user_id = ?", userId).First(&uRole)
	return uRole
}
func CreateUserRole(userId uint, roleId uint) {
	uRole := &User_role{
		UserId: strconv.Itoa(int(userId)),
		RoleId: roleId,
	}
	OldDB.Create(uRole)
}
func DeleteRoleByUserId(userId interface{}) {
	OldDB.Where("user_id = ?", userId).Delete(User_role{})
}

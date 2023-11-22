package service

import "InterfaceDroch/internal/model"

func GetName(user *model.User) string {
	if user == nil {
		return ""
	}

	return user.Name
}

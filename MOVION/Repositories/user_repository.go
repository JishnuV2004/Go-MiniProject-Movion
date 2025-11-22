package repositories

import (
	config "movion/Config"
	models "movion/Models"
)

func GetUser(userID uint)(*models.User, error){
	var user models.User
	err := config.DB.First(&user, userID).Error
	return  &user, err
}
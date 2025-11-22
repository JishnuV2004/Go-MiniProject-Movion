package repositories

import (
	config "movion/Config"
	models "movion/Models"
)

func CreateUser(user *models.User) error{
	return config.DB.Create(&user).Error
}
func GetUserByEmail(email string)(*models.User, error){
	var user models.User
	err := config.DB.Where("email = ?", email).First(&user).Error
	return &user, err
}
func SaveUser(user *models.User)(*models.User, error){
	err := config.DB.Save(&user).Error
	return user, err
}
package repositories

import (
	"errors"
	"fmt"
	config "movion/Config"
	models "movion/Models"
)
func AdminLogin(email string) (*models.User, error){
	var admin models.User
	err := config.DB.Where("email = ?", email).First(&admin).Error
	if err != nil {
		return nil , err
	}
	return &admin, err
}

func GetAllUsers(limit, offset int)([]models.User, error){
	var users []models.User
	err := config.DB.Find(&users).Limit(limit).Offset(offset).Find(&users).Error
	if err != nil {
		return nil, err
	}
	if len(users) == 0 {
		return nil, fmt.Errorf("no users found")
	}
	return users, nil
}

func GetUserById(id string)(*models.User, error){
	var user models.User
	err := config.DB.First(&user, id).Error
	return &user, err
}

func DeleteUserById(id string) error{
	var user models.User
	if err := config.DB.First(&user, id).Error; err != nil {
		return errors.New("user not found")
	}
	return config.DB.Unscoped().Delete(&user).Error
}

func SearchUser(name string, limit, offset int)([]models.User, error){
	var user []models.User
	err := config.DB.Where("username LIKE ?", "%"+name+"%").Limit(limit).Offset(offset).Find(&user).Error
	return user, err
}
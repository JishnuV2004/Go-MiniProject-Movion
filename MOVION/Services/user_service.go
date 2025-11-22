package services

import (
	"errors"
	models "movion/Models"
	repositories "movion/Repositories"

	"golang.org/x/crypto/bcrypt"
)

func Profile(userID uint) (*models.User, error){
	user, err := repositories.GetUser(userID)
	if err != nil {
		return nil, err
	}
	return user, nil
}
func UpdateUser(userID uint, input *models.User)(*models.User, error){
	user, err := repositories.GetUser(userID)
	if err != nil {
		return nil, errors.New("user not found")
	}

	if input.Username != ""{
		user.Username = input.Username
	}

	if input.Password != ""{
		hashpassword,err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)

		if err != nil {
			return nil, errors.New("password hashing faild")
		}
		user.Password = string(hashpassword)
	}
	updateUser, err := repositories.SaveUser(user)
	if err != nil {
		return nil, errors.New("update failed")
	}
	return updateUser, nil 
}
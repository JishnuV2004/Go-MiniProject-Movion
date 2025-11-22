package services

import (
	"errors"
	models "movion/Models"
	repositories "movion/Repositories"
	utils "movion/Utils"
)
// Signup Users
func Signup(user *models.User) error{
	existingUser, _ := repositories.GetUserByEmail(user.Email)
	if existingUser.Email != "" {
		return errors.New("user alredy exists")
	}
	hashedpassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashedpassword

	if user.Role == "" {
		user.Role = "user"
	}

	return repositories.CreateUser(user)
}

// Login Users
func Login(email, password string) (string, string, error){
	user, err := repositories.GetUserByEmail(email)
	if err != nil {
		return "","", errors.New("invalid username or password")
	}

	if user.IsBlocked {
		return "", "", errors.New("user is blocked. access denied")
	}

	err = utils.ComparePassword(user.Password, password)
	if err != nil {
		return "","", errors.New("invalid username or password")
	}

	accessToken, refreshToken, err := utils.GenerateToken(user.ID, user.Email, user.Role)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil

}
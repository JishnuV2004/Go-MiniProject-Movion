package services

import (
	"errors"
	"fmt"
	models "movion/Models"
	repositories "movion/Repositories"
	utils "movion/Utils"
	constants "movion/const"

	"golang.org/x/crypto/bcrypt"
)

func AdminLogin(email, password string) (*models.User, error) {
	admin, err := repositories.AdminLogin(email)
	if err != nil {
		return nil, errors.New("invalid email")
	}
	if err := utils.ComparePassword(admin.Password, password); err != nil {
		return nil, errors.New("invalid password")
	}
	if admin.Role != constants.Admin {
		return nil, errors.New("admin not found")
	}
	fmt.Println("admin", admin)
	return admin, err
}

// GetUsers
func GetAllUsers(page, limit int) ([]models.User, error) {
	offset := (page - 1) * limit

	return repositories.GetAllUsers(limit, offset)
}

// Get User ById
func GetUser(id string) (*models.User, error) {
	user, err := repositories.GetUserById(id)
	if err != nil {
		return nil, errors.New("user not found")
	}
	return user, nil
}

// Edit Users
func EditUser(userID string, input *models.EditUserInput) (*models.User, error) {
	user, err := repositories.GetUserById(userID)
	if err != nil {
		return nil, errors.New("user not found")
	}
	if input.Username != "" {
		user.Username = input.Username
	}
	if input.Email != "" {
		user.Email = input.Email
	}
	if input.Role != "" {
		user.Role = input.Role
	}
	if input.Password != "" {
		hashed, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
		if err != nil {
			return nil, errors.New("password hashing failed")
		}
		user.Password = string(hashed)
	}
	updateduser, err := repositories.SaveUser(user)
	if err != nil {
		return nil, errors.New("failed to save user")
	}
	return updateduser, nil
}

// Create Users
func CreateUser(newUser *models.User) (*models.User, error) {
	_, err := repositories.GetUserByEmail(newUser.Email)
	if err == nil {
		return nil, errors.New("user alredy exists")
	}
	hashedpassword, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New("failed password hashing")
	}
	newUser.Password = string(hashedpassword)

	if newUser.Role == "" {
		newUser.Role = "user"
	}

	if err := repositories.CreateUser(newUser); err != nil {
		return nil, err
	}

	newUser.Password = "" // hide before returning

	return newUser, nil
}

// Delete Users
func DeleteUser(id string) error {
	err := repositories.DeleteUserById(id)
	if err != nil {
		return err
	}
	return nil
}

// Search Users
func SearchUser(name string, page, limit int) ([]models.User, error) {
	offset := (page - 1) * limit

	user, err := repositories.SearchUser(name, limit, offset)
	if err != nil {
		return nil, errors.New("database error")
	}
	if len(user) == 0 {
		return nil, errors.New("users not found")
	}
	return user, nil
}

func BlockUser(id string, Block bool) (*models.User, error) {
	user, err := repositories.GetUserById(id)
	if err != nil {
		return nil, errors.New("user not found")
	}
	user.IsBlocked = Block
	user, err = repositories.SaveUser(user)
	return user, err
}

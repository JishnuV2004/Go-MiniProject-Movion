package utils

import "golang.org/x/crypto/bcrypt"

func ComparePassword(oldpassword string, newpassword string) error{
	err := bcrypt.CompareHashAndPassword([]byte(oldpassword), []byte(newpassword))
	return err
}
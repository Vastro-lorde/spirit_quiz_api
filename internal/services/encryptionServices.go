package services

import "golang.org/x/crypto/bcrypt"

// HashPassword returns the bcrypt hash of the password
func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

// ComparePassword compares a plain text password with a hashed password
func ComparePassword(plainPwd string, hashedPwd string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hashedPwd), []byte(plainPwd))
    return err == nil
}
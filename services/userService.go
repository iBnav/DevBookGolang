package services

import (
	"errors"
	models "modulo/models"
	repository "modulo/repository"
	"net/mail"
)

func CreateUser(user models.UserModel) (models.UserModel, error) {
	var err error
	emailValid := validateEmail(user.Email)
	if !emailValid || user.Name == "" {
		return models.UserModel{}, errors.New("Email or Name is invalid")
	}

	user, err = repository.CreateUser(user)
	if err != nil {
		return models.UserModel{}, err
	}

	return user, nil
}

func GetAllUsers() ([]models.UserModel, error) {
	resp, err := repository.GetAllusers()
	if err != nil {
		return []models.UserModel{}, err
	}

	return resp, nil
}

func GetUser(id uint) (models.UserModel, error) {
	resp, err := repository.GetUser(id)
	if err != nil {
		return models.UserModel{}, err
	}

	return resp, nil
}

func UpdateUser(ID uint32, user models.UserModel) error {
	return repository.UpdateUser(user, ID)
}

func DeleteUser(ID uint64) error {
	return repository.DeleteUser(ID)
}

func validateEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

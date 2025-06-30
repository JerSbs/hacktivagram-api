package service

import (
	"errors"

	"p2-livecode-3-JerSbs/config"
	"p2-livecode-3-JerSbs/dto"
	"p2-livecode-3-JerSbs/models"
	"p2-livecode-3-JerSbs/repository"
	"p2-livecode-3-JerSbs/utils"
)

func RegisterUser(input dto.UserRegisterRequest) (*models.User, error) {
	// Validasi input wajib
	if input.Email == "" || input.Password == "" || input.FullName == "" || input.Username == "" {
		return nil, errors.New("email, username, full name, and password must not be empty")
	}

	db := config.GetDB()
	userRepo := repository.NewUserRegisterRepository(db)

	// Cek apakah email/full name sudah ada
	_, err := userRepo.FindByEmailOrFullName(input.Email, input.FullName)
	if err == nil {
		return nil, ErrEmailAlreadyExists
	}

	// Hash password
	hashedPassword, err := utils.HashPassword(input.Password)
	if err != nil {
		return nil, errors.New("failed to hash password")
	}

	user := models.User{
		FullName: input.FullName,
		Username: input.Username,
		Email:    input.Email,
		Password: hashedPassword,
		Address:  input.Address,
		Age:      input.Age,
	}

	// Simpan ke database
	if err := userRepo.Create(&user); err != nil {
		return nil, err
	}

	// Jangan tampilkan password ke response
	user.Password = ""
	return &user, nil
}

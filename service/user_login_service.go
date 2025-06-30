package service

import (
	"p2-livecode-3-JerSbs/config"
	"p2-livecode-3-JerSbs/dto"
	"p2-livecode-3-JerSbs/repository"
	"p2-livecode-3-JerSbs/utils"
)

func LoginUser(input dto.UserLoginRequest) (string, error) {
	if input.Email == "" || input.Password == "" {
		return "", ErrInvalidInput
	}

	db := config.GetDB()
	userRepo := repository.NewUserLoginRepository(db)

	user, err := userRepo.FindByEmail(input.Email)
	if err != nil {
		return "", ErrUserNotFound
	}

	if !utils.CheckPasswordHash(input.Password, user.Password) {
		return "", ErrWrongPassword
	}

	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		return "", ErrInternal
	}

	return token, nil
}

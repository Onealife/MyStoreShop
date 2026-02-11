package services

import (
	"errors"

	"github.com/Onealife/Nutchapholshop/internal/core/domain/entities"
	"github.com/Onealife/Nutchapholshop/internal/core/domain/ports/repositories"
	"github.com/Onealife/Nutchapholshop/pkg/utils"
)

type AuthServiceImpl struct {
	userRepo repositories.UserRepository
}

func NewAuthService(userRepo repositories.UserRepository) *AuthServiceImpl {
	return &AuthServiceImpl{
		userRepo: userRepo,
	}
}

func (s *AuthServiceImpl) Register(req entities.RegisterRequest) (*entities.User, error) {
	existingUser, _ := s.userRepo.GetByEmail(req.Email)
	if existingUser != nil {
		return nil, errors.New("user already exists")
	}

	if err := utils.ValidatePassword(req.Password); err != nil {
		return nil, err
	}

	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	user := &entities.User{
		Email:     req.Email,
		Password:  hashedPassword,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Role:      entities.RoleUser,
		IsActive:  true,
	}
	err = s.userRepo.Create(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *AuthServiceImpl) Login(req entities.LoginRequest) (*entities.LoginResponse, error) {
	user, err := s.userRepo.GetByEmail(req.Email)
	if err != nil {
		return nil, errors.New("invalid email or password")
	}
	if !user.IsActive {
		return nil, errors.New("account is deactivated")
	}
	if !utils.CheckPassword(req.Password, user.Password) {
		return nil, errors.New("invalid email or password")
	}

	token, err := utils.GenerateJWT(user.ID, string(user.Role))
	if err != nil {
		return nil, errors.New("failed to generate token")
	}
	return &entities.LoginResponse{
		Token: token,
		User:  *user,
	}, nil

}

func (s *AuthServiceImpl) AdminRegister(req entities.AdminRegisterRequest) (*entities.User, error) {

	existingUser, _ := s.userRepo.GetByEmail(req.Email)
	if existingUser != nil {
		return nil, errors.New("user already exissts")
	}

	if err := utils.ValidatePassword(req.Password); err != nil {
		return nil, err
	}

	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	user := &entities.User{
		Email:     req.Email,
		Password:  hashedPassword,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Role:      req.Role,
		IsActive:  true,
	}

	err = s.userRepo.Create(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *AuthServiceImpl) GetUserByID(id uint) (*entities.User, error) {
	return s.userRepo.GetByID(id)
}

func (s *AuthServiceImpl) UpdateUser(user *entities.User) error {
	return s.userRepo.Update(user)
}

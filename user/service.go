package user

import "golang.org/x/crypto/bcrypt"

type Service interface {
	RegisterUser(RegisterUserInput) (User, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) RegisterUser(input RegisterUserInput) (User, error) {
	user := User{}
	user.Name = input.Name
	user.Email = input.Email
	Password, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return user, err
	}

	user.Password = string(Password)
	user.Occupation = input.Occupation
	user.Roles = "user"

	newUser, err := s.repository.Save(user) 
	if err != nil {
		return newUser, err
	}

	return newUser, nil
}
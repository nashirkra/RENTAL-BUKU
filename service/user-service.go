package service

import (
	"log"

	"github.com/mashingan/smapping"
	"github.com/nashirkra/RENTAL-BUKU/dto"
	"github.com/nashirkra/RENTAL-BUKU/entity"
	"github.com/nashirkra/RENTAL-BUKU/repository"
)

type UserService interface {
	Update(user dto.UserUpdate) entity.User
	Profile(userID string) entity.User
	All() []entity.User
	UserRole(userID string) string
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{
		userRepository: userRepo,
	}
}

func (serv *userService) Update(user dto.UserUpdate) entity.User {
	userToUpdate := entity.User{}
	err := smapping.FillStruct(&userToUpdate, smapping.MapFields(&user))
	if err != nil {
		log.Fatalf("Failed to map %v", err)
	}
	updateUser := serv.userRepository.UpdateUser(userToUpdate)
	return updateUser
}

func (serv *userService) Profile(userID string) entity.User {
	return serv.userRepository.ProfileUser(userID)
}

func (serv *userService) All() []entity.User {
	return serv.userRepository.GetAllUser()
}
func (serv *userService) UserRole(userID string) string {
	res := serv.userRepository.UserRole(userID)
	return res
}

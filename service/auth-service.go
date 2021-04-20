package service

import (
	"log"

	"github.com/mashingan/smapping"
	"github.com/nashirkra/RENTAL-BUKU/dto"
	"github.com/nashirkra/RENTAL-BUKU/entity"
	"github.com/nashirkra/RENTAL-BUKU/repository"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	VerifyCredential(email string, password string) interface{}
	CreateUser(user dto.Register) entity.User
	FindByEmail(email string) entity.User
	IsDuplicateEmail(email string) bool
	UserRole(userID string) string
}

type authService struct {
	userRepository repository.UserRepository
}

func NewAuthService(userRep repository.UserRepository) AuthService {
	return &authService{
		userRepository: userRep,
	}
}

func (serv *authService) VerifyCredential(email string, password string) interface{} {
	res := serv.userRepository.VerifyCredential(email, password)
	if v, ok := res.(entity.User); ok {
		comparedPassword := comparedPassword(v.Password, []byte(password))
		if v.Email == email && comparedPassword {
			return res
		}
		return nil
	}
	return nil
}

func (serv *authService) CreateUser(user dto.Register) entity.User {
	userToCreate := entity.User{}
	err := smapping.FillStruct(&userToCreate, smapping.MapFields(&user))
	if err != nil {
		log.Fatalf("Failed map %v", err)
	}
	res := serv.userRepository.InsertUser(userToCreate)
	return res
}

func (serv *authService) FindByEmail(email string) entity.User {
	return serv.userRepository.FindByEmail(email)
}

func (serv *authService) IsDuplicateEmail(email string) bool {
	res := serv.userRepository.IsDuplicateEmail(email)
	return !(res.Error == nil)
}

func (serv *authService) UserRole(userID string) string {
	res := serv.userRepository.ProfileUser(userID)
	return res.Role
}

func comparedPassword(hashedPwd string, plainPwd []byte) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

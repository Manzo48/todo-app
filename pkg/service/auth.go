package service

import (
	"crypto/sha1"
	"fmt"

	"github.com/Manzo48/todo-app"
	"github.com/Manzo48/todo-app/pkg/repository"
)
const salt = "hiejwlkr2pokda;w,12"
type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService{
	return &AuthService{repo: repo}
}
func (s *AuthService) GetAllUsers() ([]todo.User, error) {
	return s.repo.GetAllUsers()
}

func (s *AuthService) CreateUser(user todo.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)

}


func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
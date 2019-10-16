package services

import (
	"gitlab.com/my-app/infrastructures"

	"gitlab.com/my-app/models"
	"gitlab.com/my-app/repositories"
)

type IUsersService interface {
	GetUsers() ([]models.Users, error)
	StoreUser(models.UserInput) (models.Users, error)
	GetUserByID(int) (models.Users, error)
}

type UsersService struct {
	UsersRepository repositories.IUsersRepository
}

func InitUsersService() *UsersService {
	usersRepository := new(repositories.UsersRepository)
	usersRepository.DB = &infrastructures.SQLConnection{}

	usersService := new(UsersService)
	usersService.UsersRepository = usersRepository

	return usersService
}

func (r *UsersService) StoreUser(data models.UserInput) (result models.Users, err error) {
	var user models.Users
	user.Email = data.Email
	user.Name = data.Name
	user.Password = data.Password
	user.Phone = data.Phone
	user.Status = data.Status
	result, err = r.UsersRepository.StoreUser(user)
	return
}

//list users
func (p *UsersService) GetUsers() (users []models.Users, err error) {
	users, err = p.UsersRepository.GetUsers()
	return
}

//user by id
func (p *UsersService) GetUserByID(ID int) (result models.Users, err error) {
	result, err = p.UsersRepository.GetUserByID(ID)
	return
}

package usecase

import (
	"github.com/Rawipass/golang-test-service/internal/user/repository"
	"github.com/Rawipass/golang-test-service/models"
)

type UserUseCase struct {
	repo repository.UserRepository
}

func NewUserUseCase(repo *repository.UserRepository) *UserUseCase {
	return &UserUseCase{repo: *repo}
}

func (u *UserUseCase) ListUsers(limit, page int) ([]models.User, int, error) {
	offset := (page - 1) * limit
	return u.repo.GetAllUsers(limit, offset)
}

func (u *UserUseCase) GetUserDetail(id string) (*models.User, error) {
	return u.repo.GetUserByID(id)
}

func (u *UserUseCase) DeductBalance(id string, amount float64) error {
	return u.repo.UpdateUserBalance(id, -amount)
}

func (u *UserUseCase) AddBalance(id string, amount float64) error {
	return u.repo.UpdateUserBalance(id, amount)
}

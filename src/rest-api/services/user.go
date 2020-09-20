package services

import (
	"rest-api/models"
)

type userService struct{}

var (
	allUsers          = map[int64]*models.User{}
	userID      int64 = 1
	UserService       = userService{}
)

func (service userService) Create(user models.User) (*models.User, *models.HttpError) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	user.ID = userID
	userID++
	allUsers[user.ID] = &user
	return &user, nil
}

func (service userService) Get(userID int64) (*models.User, *models.HttpError) {
	if user := allUsers[userID]; user != nil {
		return user, nil
	}
	return nil, models.NotFoundError("User not found")
}

func (service userService) GetAll() (map[int64]*models.User, *models.HttpError) {
	if len(allUsers) > 0 {
		return allUsers, nil
	}
	return nil, models.NotFoundError("No users found")
}

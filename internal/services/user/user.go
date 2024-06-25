package user

import (
	"context"
	"life-restart-backend/internal/dao"
	"life-restart-backend/internal/dao/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserService struct {
	userDAO *dao.UserDAO
}

func NewUserService(userDAO *dao.UserDAO) *UserService {
	return &UserService{
		userDAO: userDAO,
	}
}

func (s *UserService) RegisterUser(ctx context.Context, user *models.User) (primitive.ObjectID, error) {
	return s.userDAO.CreateUser(ctx, user)
}

func (s *UserService) GetUserByID(ctx context.Context, id primitive.ObjectID) (*models.User, error) {
	return s.userDAO.GetUserByID(ctx, id)
}

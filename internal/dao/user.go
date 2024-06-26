package dao

import (
	"context"

	"life-restart-backend/internal/dao/models"
	"life-restart-backend/internal/pkg/database"
	"log"

	"github.com/qiniu/qmgo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserDAO struct {
	collection *qmgo.Collection
}

func NewUserDAO() *UserDAO {
	return &UserDAO{
		collection: database.GetCollection("users"),
	}
}

func (dao *UserDAO) CreateUser(ctx context.Context, user *models.User) (primitive.ObjectID, error) {
	result, err := dao.collection.InsertOne(ctx, user)
	if err != nil {
		log.Printf("Failed to create user: %v", err)
		return primitive.NilObjectID, err
	}
	// 提取并返回插入的 ID
	id := result.InsertedID.(primitive.ObjectID)
	return id, nil
}

func (dao *UserDAO) GetUserByID(ctx context.Context, id primitive.ObjectID) (*models.User, error) {
	var user models.User
	err := dao.collection.Find(ctx, bson.M{"_id": id}).One(&user)
	if err != nil {
		log.Printf("Failed to get user by ID: %v", err)
		return nil, err
	}
	return &user, nil
}

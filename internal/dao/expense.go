package dao

import (
	"context"
	"life-restart-backend/internal/dao/models"
	"life-restart-backend/internal/pkg/database"

	"github.com/qiniu/qmgo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ExpenseDAO struct {
	collection *qmgo.Collection
}

func NewExpenseDAO() *ExpenseDAO {
	return &ExpenseDAO{
		collection: database.GetCollection("expenses"),
	}
}

func (dao *ExpenseDAO) CreateExpense(ctx context.Context, expense *models.Expense) (primitive.ObjectID, error) {
	result, err := dao.collection.InsertOne(ctx, expense)
	if err != nil {
		return primitive.NilObjectID, err
	}
	return result.InsertedID.(primitive.ObjectID), nil
}

func (dao *ExpenseDAO) GetExpenseByID(ctx context.Context, id primitive.ObjectID) (*models.Expense, error) {
	var expense models.Expense
	err := dao.collection.Find(ctx, bson.M{"_id": id}).One(&expense)
	return &expense, err
}

func (dao *ExpenseDAO) GetAllExpenses(ctx context.Context) ([]*models.Expense, error) {
	var expenses []*models.Expense
	err := dao.collection.Find(ctx, bson.M{}).All(&expenses)
	return expenses, err
}

func (dao *ExpenseDAO) UpdateExpense(ctx context.Context, id primitive.ObjectID, expense *models.Expense) error {
	return dao.collection.UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": expense})
}

func (dao *ExpenseDAO) DeleteExpense(ctx context.Context, id primitive.ObjectID) error {
	return dao.collection.Remove(ctx, bson.M{"_id": id})
}

package expense

import (
	"context"
	"life-restart-backend/internal/dao"
	"life-restart-backend/internal/dao/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ExpenseService struct {
	dao *dao.ExpenseDAO
}

func NewExpenseService(dao *dao.ExpenseDAO) *ExpenseService {
	return &ExpenseService{dao: dao}
}

func (s *ExpenseService) CreateExpense(ctx context.Context, expense *models.Expense) (primitive.ObjectID, error) {
	return s.dao.CreateExpense(ctx, expense)
}

func (s *ExpenseService) GetExpenseByID(ctx context.Context, id primitive.ObjectID) (*models.Expense, error) {
	return s.dao.GetExpenseByID(ctx, id)
}

func (s *ExpenseService) GetAllExpenses(ctx context.Context) ([]*models.Expense, error) {
	return s.dao.GetAllExpenses(ctx)
}

func (s *ExpenseService) UpdateExpense(ctx context.Context, id primitive.ObjectID, expense *models.Expense) error {
	return s.dao.UpdateExpense(ctx, id, expense)
}

func (s *ExpenseService) DeleteExpense(ctx context.Context, id primitive.ObjectID) error {
	return s.dao.DeleteExpense(ctx, id)
}

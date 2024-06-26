package plan

import (
	"context"
	"life-restart-backend/internal/dao"
	"life-restart-backend/internal/dao/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PlanService struct {
	dao *dao.PlanDAO
}

func NewPlanService(dao *dao.PlanDAO) *PlanService {
	return &PlanService{dao: dao}
}

func (s *PlanService) CreatePlan(ctx context.Context, plan *models.Plan) (primitive.ObjectID, error) {
	return s.dao.CreatePlan(ctx, plan)
}

func (s *PlanService) GetPlanByID(ctx context.Context, id primitive.ObjectID) (*models.Plan, error) {
	return s.dao.GetPlanByID(ctx, id)
}

func (s *PlanService) GetAllPlans(ctx context.Context) ([]*models.Plan, error) {
	return s.dao.GetAllPlans(ctx)
}

func (s *PlanService) UpdatePlan(ctx context.Context, id primitive.ObjectID, plan *models.Plan) error {
	return s.dao.UpdatePlan(ctx, id, plan)
}

func (s *PlanService) DeletePlan(ctx context.Context, id primitive.ObjectID) error {
	return s.dao.DeletePlan(ctx, id)
}

package dao

import (
	"context"
	"life-restart-backend/internal/dao/models"
	"life-restart-backend/internal/pkg/database"

	"github.com/qiniu/qmgo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PlanDAO struct {
	collection *qmgo.Collection
}

func NewPlanDAO() *PlanDAO {
	return &PlanDAO{
		collection: database.GetCollection("plans"),
	}
}

func (dao *PlanDAO) CreatePlan(ctx context.Context, plan *models.Plan) (primitive.ObjectID, error) {
	result, err := dao.collection.InsertOne(ctx, plan)
	if err != nil {
		return primitive.NilObjectID, err
	}
	return result.InsertedID.(primitive.ObjectID), nil
}

func (dao *PlanDAO) GetPlanByID(ctx context.Context, id primitive.ObjectID) (*models.Plan, error) {
	var plan models.Plan
	err := dao.collection.Find(ctx, bson.M{"_id": id}).One(&plan)
	return &plan, err
}

func (dao *PlanDAO) GetAllPlans(ctx context.Context) ([]*models.Plan, error) {
	var plans []*models.Plan
	err := dao.collection.Find(ctx, bson.M{}).All(&plans)
	return plans, err
}

func (dao *PlanDAO) UpdatePlan(ctx context.Context, id primitive.ObjectID, plan *models.Plan) error {
	return dao.collection.UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": plan})
}

func (dao *PlanDAO) DeletePlan(ctx context.Context, id primitive.ObjectID) error {
	return dao.collection.Remove(ctx, bson.M{"_id": id})
}

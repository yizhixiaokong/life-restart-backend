package dao

import (
	"context"
	"life-restart-backend/internal/dao/models"
	"life-restart-backend/internal/pkg/database"

	"github.com/qiniu/qmgo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ReadingEntryDAO struct {
	collection *qmgo.Collection
}

func NewReadingEntryDAO() *ReadingEntryDAO {
	return &ReadingEntryDAO{
		collection: database.GetCollection("reading_entries"),
	}
}

func (dao *ReadingEntryDAO) CreateReadingEntry(ctx context.Context, entry *models.ReadingEntry) (primitive.ObjectID, error) {
	result, err := dao.collection.InsertOne(ctx, entry)
	if err != nil {
		return primitive.NilObjectID, err
	}
	return result.InsertedID.(primitive.ObjectID), nil
}

func (dao *ReadingEntryDAO) GetReadingEntryByID(ctx context.Context, id primitive.ObjectID) (*models.ReadingEntry, error) {
	var entry models.ReadingEntry
	err := dao.collection.Find(ctx, bson.M{"_id": id}).One(&entry)
	return &entry, err
}

func (dao *ReadingEntryDAO) GetAllReadingEntries(ctx context.Context) ([]*models.ReadingEntry, error) {
	var entries []*models.ReadingEntry
	err := dao.collection.Find(ctx, bson.M{}).All(&entries)
	return entries, err
}

func (dao *ReadingEntryDAO) UpdateReadingEntry(ctx context.Context, id primitive.ObjectID, entry *models.ReadingEntry) error {
	return dao.collection.UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": entry})
}

func (dao *ReadingEntryDAO) DeleteReadingEntry(ctx context.Context, id primitive.ObjectID) error {
	return dao.collection.Remove(ctx, bson.M{"_id": id})
}

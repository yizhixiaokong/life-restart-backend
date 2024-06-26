package readingentry

import (
	"context"
	"life-restart-backend/internal/dao"
	"life-restart-backend/internal/dao/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ReadingEntryService struct {
	dao *dao.ReadingEntryDAO
}

func NewReadingEntryService(dao *dao.ReadingEntryDAO) *ReadingEntryService {
	return &ReadingEntryService{dao: dao}
}

func (s *ReadingEntryService) CreateReadingEntry(ctx context.Context, entry *models.ReadingEntry) (primitive.ObjectID, error) {
	return s.dao.CreateReadingEntry(ctx, entry)
}

func (s *ReadingEntryService) GetReadingEntryByID(ctx context.Context, id primitive.ObjectID) (*models.ReadingEntry, error) {
	return s.dao.GetReadingEntryByID(ctx, id)
}

func (s *ReadingEntryService) GetAllReadingEntries(ctx context.Context) ([]*models.ReadingEntry, error) {
	return s.dao.GetAllReadingEntries(ctx)
}

func (s *ReadingEntryService) UpdateReadingEntry(ctx context.Context, id primitive.ObjectID, entry *models.ReadingEntry) error {
	return s.dao.UpdateReadingEntry(ctx, id, entry)
}

func (s *ReadingEntryService) DeleteReadingEntry(ctx context.Context, id primitive.ObjectID) error {
	return s.dao.DeleteReadingEntry(ctx, id)
}

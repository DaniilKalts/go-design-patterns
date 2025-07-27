package repository

import (
	"context"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/DaniilKalts/go-design-patterns/abstract-factory/internal/domain"
)

// ItemRepo : Abstract Product B
type ItemRepo interface {
	CreateItem(ctx context.Context, item *domain.Item) error
	GetItemsByOwnerID(ctx context.Context, ownerID uuid.UUID) (
		[]*domain.Item, error,
	)
}

// itemRepo - Concrete Product B
type itemRepo struct {
	col *mongo.Collection
}

func NewItemRepo(col *mongo.Collection) ItemRepo {
	return &itemRepo{col}
}

func (r *itemRepo) CreateItem(ctx context.Context, item *domain.Item) error {
	_, err := r.col.InsertOne(ctx, item)

	return err
}

func (r *itemRepo) GetItemsByOwnerID(
	ctx context.Context, ownerID uuid.UUID,
) ([]*domain.Item, error) {
	filter := bson.M{"owner_id": ownerID}

	cursor, err := r.col.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer func() {
		if cursorErr := cursor.Close(ctx); cursorErr != nil && err == nil {
			err = cursorErr
		}
	}()

	var items []*domain.Item
	if err = cursor.All(ctx, &items); err != nil {
		return nil, err
	}

	return items, nil
}

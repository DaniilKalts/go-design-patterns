package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/DaniilKalts/go-design-patterns/abstract-factory/internal/domain"
)

// UserRepo : Abstract Product A
type UserRepo interface {
	CreateUser(ctx context.Context, user *domain.User) (*domain.User, error)
	GetUserByEmail(ctx context.Context, email string) (*domain.User, error)
}

// userRepo - Concrete Product A
type userRepo struct {
	col *mongo.Collection
}

func NewUserRepo(col *mongo.Collection) UserRepo {
	return &userRepo{col}
}

func (r *userRepo) CreateUser(ctx context.Context, user *domain.User) (
	*domain.User, error,
) {
	_, err := r.col.InsertOne(ctx, user)

	return user, err
}

func (r *userRepo) GetUserByEmail(ctx context.Context, email string) (
	*domain.User, error,
) {
	var user domain.User

	err := r.col.FindOne(ctx, bson.M{"email": email}).Decode(&user)

	return &user, err
}

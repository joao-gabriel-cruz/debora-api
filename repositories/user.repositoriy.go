package repository

import (
	"context"

	"github.com/joao-gabriel-cruz/debora-api/model"
	"github.com/joao-gabriel-cruz/debora-api/prisma/db"
)

type UserRepository interface {
	Create(ctx context.Context, user model.User) error
	FindAll(ctx context.Context) ([]db.UserModel, error)
	Update(ctx context.Context, id string, user model.User) error
	Delete(ctx context.Context, id string) error
	FindByID(ctx context.Context, id string) (*db.UserModel, error)
}

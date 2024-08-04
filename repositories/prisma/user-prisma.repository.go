package prisma_repository

import (
	"context"

	"github.com/joao-gabriel-cruz/debora-api/lib"
	"github.com/joao-gabriel-cruz/debora-api/model"
	"github.com/joao-gabriel-cruz/debora-api/prisma/db"
	repository "github.com/joao-gabriel-cruz/debora-api/repositories"
)

type UserPrismaRepository struct {
	db *db.PrismaClient
}

// FindByID implements repository.UserRepository.
func (u *UserPrismaRepository) FindByID(ctx context.Context, id string) (*db.UserModel, error) {
	user, err := u.db.User.FindUnique(
		db.User.ID.Equals(id),
	).Exec(ctx)

	return lib.ErrorData(user, err)
}

// Create implements repository.UserRepository.
func (u *UserPrismaRepository) Create(ctx context.Context, user model.User) error {
	_, err := u.db.User.CreateOne(
		db.User.Email.Set(user.Email),
		db.User.Name.Set(user.Name),
		db.User.Password.Set(user.Password),
	).Exec(ctx)

	return lib.Error(err)
}

// Delete implements repository.UserRepository.
func (u *UserPrismaRepository) Delete(ctx context.Context, id string) error {
	_, err := u.db.User.FindUnique(
		db.User.ID.Equals(id),
	).Delete().Exec(ctx)

	return lib.Error(err)
}

// FindAll implements repository.UserRepository.
func (u *UserPrismaRepository) FindAll(ctx context.Context) ([]db.UserModel, error) {

	users, err := u.db.User.FindMany().Exec(ctx)

	return lib.ErrorData(users, err)

}

// Update implements repository.UserRepository.
func (u *UserPrismaRepository) Update(ctx context.Context, id string, user model.User) error {
	_, err := u.db.User.FindUnique(
		db.User.ID.Equals(id),
	).Update(
		db.User.Email.Set(user.Email),
		db.User.Name.Set(user.Name),
		db.User.Password.Set(user.Password),
	).Exec(ctx)

	return lib.Error(err)

}

func NewUserPrismaRepository(db *db.PrismaClient) repository.UserRepository {
	return &UserPrismaRepository{
		db: db,
	}
}

package prisma_repository

import (
	"context"

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

	if err != nil {
		return nil, err
	}

	return user, nil
}

// Create implements repository.UserRepository.
func (u *UserPrismaRepository) Create(ctx context.Context, user model.User) error {
	_, err := u.db.User.CreateOne(
		db.User.Email.Set(user.Email),
		db.User.Name.Set(user.Name),
		db.User.Password.Set(user.Password),
	).Exec(ctx)

	if err != nil {
		println("Error creating user")
		return err

	}
	return nil
}

// Delete implements repository.UserRepository.
func (u *UserPrismaRepository) Delete(ctx context.Context, id string) error {
	_, err := u.db.User.FindUnique(
		db.User.ID.Equals(id),
	).Delete().Exec(ctx)

	if err != nil {
		return err
	}
	return nil
}

// FindAll implements repository.UserRepository.
func (u *UserPrismaRepository) FindAll(ctx context.Context) ([]db.UserModel, error) {

	users, err := u.db.User.FindMany().Exec(ctx)

	if err != nil {
		return nil, err
	}

	return users, nil
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

	if err != nil {
		return err
	}

	return nil
}

func NewUserPrismaRepository(db *db.PrismaClient) repository.UserRepository {
	return &UserPrismaRepository{
		db: db,
	}
}

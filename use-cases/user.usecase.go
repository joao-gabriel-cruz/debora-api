package usecases

import (
	"context"

	"github.com/joao-gabriel-cruz/debora-api/model"
	"github.com/joao-gabriel-cruz/debora-api/prisma/db"
	repository "github.com/joao-gabriel-cruz/debora-api/repositories"
)

type UserUseCase struct {
	userRepository repository.UserRepository
}

func NewUserUseCase(userPrismaRepository repository.UserRepository) UserUseCase {
	return UserUseCase{
		userRepository: userPrismaRepository,
	}
}

func (u *UserUseCase) CreateUser(ctx context.Context, user model.User) error {
	err := u.userRepository.Create(ctx, user)

	if err != nil {
		return err
	}
	return nil
}

func (u *UserUseCase) UpdateUser(ctx context.Context, id string, user model.User) error {
	err := u.userRepository.Update(ctx, id, user)

	if err != nil {
		return err
	}
	return nil
}

func (u *UserUseCase) DeleteUser(ctx context.Context, id string) error {
	err := u.userRepository.Delete(ctx, id)

	if err != nil {
		return err
	}
	return nil
}

func (u *UserUseCase) GetUserByID(ctx context.Context, id string) (user *db.UserModel, err error) {
	userById, err := u.userRepository.FindByID(ctx, id)

	if err != nil {
		return nil, err
	}
	return userById, nil
}

func (u *UserUseCase) GetUser(ctx context.Context) (user []db.UserModel, err error) {
	// Get user from database
	users, err := u.userRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	return users, nil
}

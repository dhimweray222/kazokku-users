package repository

import (
	"context"

	"github.com/dhimweray222/users/exception"
	"github.com/dhimweray222/users/model/domain"
	"github.com/jackc/pgx/v5"
)

type UserRepositoryImpl struct {
	DB Store
}
type UserRepository interface {
	CreateUserImpl(ctx context.Context, user domain.User) (domain.User, error)
	FindUserByIdTx(ctx context.Context, value string) (domain.User, error)
	FindAllUserTx(ctx context.Context, search, page, limit, searchName, searchEmail string) ([]domain.User, error)
	UpdateUserImpl(ctx context.Context, user domain.User) error
}

func NewUserRepository(db Store) UserRepository {
	return &UserRepositoryImpl{
		DB: db,
	}
}

func (repository *UserRepositoryImpl) CreateUserImpl(ctx context.Context, user domain.User) (domain.User, error) {
	var err error

	err = repository.DB.WithTransaction(ctx, func(tx pgx.Tx) error {
		err = repository.CreateUser(ctx, tx, user)
		if err != nil {
			return exception.ErrorBadRequest(err.Error())
		}
		return nil
	})
	if err != nil {
		return domain.User{}, err
	}

	return user, err
}

func (repository *UserRepositoryImpl) FindUserByIdTx(ctx context.Context, id string) (domain.User, error) {

	var data domain.User
	var err error

	err = repository.DB.WithTransaction(ctx, func(tx pgx.Tx) error {

		data, err = repository.FindUserById(ctx, tx, id)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return domain.User{}, err
	}

	return data, err
}

func (repository *UserRepositoryImpl) FindAllUserTx(ctx context.Context, search, page, limit, searchName, searchEmail string) ([]domain.User, error) {

	var data []domain.User
	var err error

	err = repository.DB.WithTransaction(ctx, func(tx pgx.Tx) error {

		data, err = repository.FindAllUser(ctx, tx, search, page, limit, searchName, searchEmail)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return []domain.User{}, err
	}

	return data, err
}

func (repository *UserRepositoryImpl) UpdateUserImpl(ctx context.Context, user domain.User) error {
	var err error

	err = repository.DB.WithTransaction(ctx, func(tx pgx.Tx) error {

		err = repository.CreateUser(ctx, tx, user)
		if err != nil {
			return exception.ErrorBadRequest(err.Error())
		}
		return nil
	})

	if err != nil {
		return err
	}

	return err
}

package repository

import (
	"context"
	"database/sql"
	"errors"
	"golang_jsonplaceholder_without_framework/helper"
	"golang_jsonplaceholder_without_framework/model"
)

type UserRepositoryImpl struct {
	Db *sql.DB
}

func NewUserRepository(Db *sql.DB) UserRepository {
	return &UserRepositoryImpl{Db: Db}
}

// Delete implements UserRepository
func (u *UserRepositoryImpl) Delete(ctx context.Context, userId int) {
	tx, err := u.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	SQL := "DELETE FROM users WHERE id = $1"

	_, errExec := tx.ExecContext(ctx, SQL, userId)
	helper.PanicIfError(errExec)
}

// FindAll implements UserRepository
func (u *UserRepositoryImpl) FindAll(ctx context.Context) []model.User {
	tx, err := u.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	SQL := "SELECT * FROM users"
	result, errQuery := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(errQuery)
	defer result.Close()

	var users []model.User

	for result.Next() {
		user := model.User{}
		err := result.Scan(&user.Id, &user.Name, &user.UserName, &user.Email)
		helper.PanicIfError(err)

		users = append(users, user)
	}

	return users
}

// FindById implements UserRepository
func (u *UserRepositoryImpl) FindById(ctx context.Context, userId int) (model.User, error) {
	tx, err := u.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	SQL := "SELECT id, name, username, email FROM users WHERE id = $1"

	result, errQuery := tx.QueryContext(ctx, SQL, userId)
	helper.PanicIfError(errQuery)
	defer result.Close()

	user := model.User{}

	if result.Next() {
		err := result.Scan(&user.Id, &user.Name, &user.UserName, &user.Email)
		helper.PanicIfError(err)
		return user, nil
	} else {
		return user, errors.New("user id not found")
	}
}

// Save implements UserRepository
func (u *UserRepositoryImpl) Save(ctx context.Context, user model.User) {
	tx, err := u.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	SQL := "INSERT INTO users(name, username, email) VALUES ($1, $2, $3)"
	_, err = tx.ExecContext(ctx, SQL, user.Name, user.UserName, user.Email)
	helper.PanicIfError(err)
}

// Update implements UserRepository
func (u *UserRepositoryImpl) Update(ctx context.Context, user model.User) {
	tx, err := u.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	SQL := "UPDATE users SET name = $1, username = $2, email = $3 WHERE id = $4"
	_, err = tx.ExecContext(ctx, SQL, user.Name, user.UserName, user.Email, user.Id)
	helper.PanicIfError(err)
}

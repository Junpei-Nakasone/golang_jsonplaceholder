package repository

import "database/sql"

type UserRepositoryImpl struct {
	Db *sql.DB
}

func NewUserRepository(Db *sql.DB) {

}

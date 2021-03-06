package services

import (
	"database/sql"
	"github.com/pngouin/pinshorter/db"
	"github.com/pngouin/pinshorter/models"
)

func NewUser(sql *sql.DB) User {
	return User{db.NewUser(sql)}
}

type User struct {
	database db.User
}

func (u User) Connection(user models.UserConn) (models.UserInfo, bool, error) {
	return u.database.Connection(user)
}

func (u User) Add(conn models.UserConn) (models.UserInfo, error) {
	return u.database.Add(conn)
}

func (u User) GetById(user models.UserInfo) (models.UserInfo, error) {
	return u.database.GetById(user)
}

func (u User) IsUserExist(user models.UserInfo) bool {
	return u.database.IsUserExist(user)
}

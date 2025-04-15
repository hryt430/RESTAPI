package database

import (
	"errors"
	"fmt"

	"github.com/hryt430/RESTAPI/api/internal/domain/entity"
)

type UserServiceRepository struct {
	SqlHandler
}

func (repository *UserServiceRepository) FindUser() (users []*entity.User, err error) {
	rows, err := repository.Query("SELECT id, username FROM users")
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		var username string
		if err := rows.Scan(&id, &username); err != nil {
			continue
		}
		user := entity.User{
			ID:       id,
			Username: username,
		}
		users = append(users, &user)
	}
	return
}

func (repository *UserServiceRepository) FindUserById(identifier int) (user *entity.User, err error) {
	row, err := repository.Query("SELECT id, username, password FROM users WHERE id = ?", identifier)
	if err != nil {
		return
	}
	defer row.Close()

	if row.Next() {
		var id int
		var username string
		var password string
		if err = row.Scan(&id, &username, &password); err != nil {
			return
		}

		user = &entity.User{
			ID:       id,
			Username: username,
			Password: password,
		}
		return
	}

	// データが見つからなかった場合
	err = fmt.Errorf("user with id %d not found", identifier)
	return
}

func (repository *UserServiceRepository) Save(user *entity.User) (id int, err error) {
	result, err := repository.Execute(
		"INSERT INTO users (username, password) VALUES (?, ?)", user.Username, user.Password,
	)
	if err != nil {
		return
	}
	id64, err := result.LastInsertId()
	if err != nil {
		return
	}
	id = int(id64)
	return
}

func (repository *UserServiceRepository) Edit(identifier int, user *entity.User) (id int, err error) {
	result, err := repository.Execute(
		"UPDATE users SET username = ?, password = ? WHERE id = ?",
		user.Username, user.Password, identifier,
	)
	if err != nil {
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return
	}

	if rowsAffected == 0 {
		err = errors.New("no user found with given id")
		return
	}

	id = identifier
	return
}

func (repo *UserServiceRepository) Delete(id int) (err error) {
	result, err := repo.Execute("DELETE FROM users WHERE id = ?", id)
	if err != nil {
		return
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return
	}

	if affected == 0 {
		err = errors.New("no user found with given id")
	}
	return
}

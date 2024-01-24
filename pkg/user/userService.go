package user

import (
	"database/sql"
)

type UserService struct {
	DB *sql.DB
}

func NewUserService(db *sql.DB) *UserService {
	return &UserService{DB: db}
}

func (us *UserService) GetAllUsers() ([]User, error) {
	rows, err := us.DB.Query("SELECT id, name FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var u User
		if err := rows.Scan(&u.ID, &u.Name); err != nil {
			return nil, err
		}
		users = append(users, u)
	}

	return users, nil
}

func (us *UserService) CreateUser(name string) (int64, error) {
	result, err := us.DB.Exec("INSERT INTO users (name) VALUES (?)", name)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (us *UserService) GetUser(id int) (*User, error) {
	row := us.DB.QueryRow("SELECT id, name FROM users WHERE id = ?", id)

	var u User
	err := row.Scan(&u.ID, &u.Name)
	if err != nil {
		return nil, err
	}

	return &u, nil
}

func (us *UserService) DeleteUser(id int) error {
	_, err := us.DB.Exec("DELETE FROM users WHERE id = ?", id)
	return err
}

package repository

import (
	"database/sql"
	"lmd-func/models"
)

// UserRepository は、ユーザー情報を操作するためのリポジトリです
type UserRepository struct {
	DB *sql.DB
}

// NewUserRepository は、新しいUserRepositoryを作成します
func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{DB: db}
}

// GetUsers は、すべてのユーザーを取得します
func (r *UserRepository) GetUsers() ([]models.User, error) {
	rows, err := r.DB.Query("SELECT id, name, email FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

// GetUser は、指定されたIDに一致するユーザーを取得します
func (r *UserRepository) GetUser(id string) (models.User, error) {
	row := r.DB.QueryRow("SELECT id, name, email FROM users WHERE id = ?", id)

	var user models.User
	if err := row.Scan(&user.ID, &user.Name, &user.Email); err != nil {
		return models.User{}, err
	}

	return user, nil
}

// CreateUser は、新しいユーザーを作成します
func (r *UserRepository) CreateUser(user *models.User) error {
	_, err := r.DB.Exec("INSERT INTO users (id, name, email) VALUES (?, ?, ?)", user.ID, user.Name, user.Email)
	return err
}

// UpdateUser は、指定されたIDのユーザーを更新します
func (r *UserRepository) UpdateUser(user *models.User) error {
	_, err := r.DB.Exec("UPDATE users SET name = ?, email = ? WHERE id = ?", user.Name, user.Email, user.ID)
	return err
}

// DeleteUser は、指定されたIDのユーザーを削除します
func (r *UserRepository) DeleteUser(id string) error {
	_, err := r.DB.Exec("DELETE FROM users WHERE id = ?", id)
	return err
}

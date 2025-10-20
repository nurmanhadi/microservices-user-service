package repository

import (
	"fmt"
	"strings"
	"user-service/src/internal/entity"

	"github.com/jmoiron/sqlx"
)

type UserRepository interface {
	Insert(user entity.User) error
	FindByEmail(email string) (*entity.User, error)
	CountByEmail(email string) (int64, error)
	CountByPhone(phone string) (int, error)
	GetUserByID(id string) (*entity.User, error)
	GetAllUsers() ([]entity.User, error)
	UpdateStatusByID(id string, status string) error
	UpdateProfile(id string, user entity.User) error
	CountByID(id string) (int, error)
}
type userRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}
func (r *userRepository) Insert(user entity.User) error {
	_, err := r.db.Exec("INSERT INTO users(id, email, password, role, status) VALUES($1, $2, $3, $4, $5)", user.Id, user.Email, user.Password, user.Role, user.Status)
	if err != nil {
		return err
	}
	return nil
}
func (r *userRepository) FindByEmail(email string) (*entity.User, error) {
	var user entity.User
	err := r.db.Get(&user, "SELECT * FROM users WHERE email = $1", email)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
func (r *userRepository) CountByEmail(email string) (int64, error) {
	var total int64
	err := r.db.Get(&total, "SELECT COUNT(email) AS total FROM users WHERE email = $1", email)
	if err != nil {
		return 0, err
	}
	return total, err
}
func (r *userRepository) CountByPhone(phone string) (int, error) {
	var count int
	query := `SELECT COUNT(*) FROM users WHERE phone = $1`
	err := r.db.Get(&count, query, phone)
	if err != nil {
		return 0, err
	}
	return count, nil
}
func (r *userRepository) GetUserByID(id string) (*entity.User, error) {
	var user entity.User
	query := `SELECT * FROM users WHERE id = $1`
	err := r.db.Get(&user, query, id)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
func (r *userRepository) GetAllUsers() ([]entity.User, error) {
	var users []entity.User
	query := `SELECT * FROM users ORDER BY created_at DESC`
	err := r.db.Select(&users, query)
	if err != nil {
		return nil, err
	}
	return users, nil
}
func (r *userRepository) UpdateStatusByID(id string, status string) error {
	query := `
		UPDATE users 
		SET status = $1, updated_at = NOW()
		WHERE id = $2
	`
	_, err := r.db.Exec(query, status, id)
	return err
}
func (r *userRepository) UpdateProfile(id string, user entity.User) error {
	query := "UPDATE users SET "
	if user.FirstName != nil {
		query += fmt.Sprintf("first_name = '%s',", *user.FirstName)
	}
	if user.LastName != nil {
		query += fmt.Sprintf("last_name = '%s',", *user.LastName)
	}
	if user.Phone != nil {
		query += fmt.Sprintf("phone = '%s',", *user.Phone)
	}
	newQuery := strings.TrimSuffix(query, ",")
	finalQuery := fmt.Sprintf("%s WHERE id = $1", newQuery)

	stmt, err := r.db.Prepare(finalQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}
	return nil
}
func (r *userRepository) CountByID(id string) (int, error) {
	var count int
	query := `SELECT COUNT(*) FROM users WHERE id = $1`
	err := r.db.Get(&count, query, id)
	if err != nil {
		return 0, err
	}
	return count, nil
}

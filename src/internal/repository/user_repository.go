package repository

import (
	"user-service/src/internal/entity"

	"github.com/jmoiron/sqlx"
)

type UserRepository interface {
	Insert(user entity.User) error
	FindByEmail(email string) (*entity.User, error)
	CountByEmail(email string) (int64, error)
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
	_, err := r.db.Exec("INSERT INTO users(id, email, password) VALUES($1, $2, $3)", user.Id, user.Email, user.Password)
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

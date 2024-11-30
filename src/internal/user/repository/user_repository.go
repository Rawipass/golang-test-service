package repository

import (
	"context"

	"github.com/Rawipass/golang-test-service/models"
	"github.com/jackc/pgx/v4/pgxpool"
)

type UserRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) GetAllUsers(limit, offset int) ([]models.User, int, error) {
	var users []models.User
	var total int

	query := `SELECT id, username, balance, affiliate_id FROM users LIMIT $1 OFFSET $2`
	rows, err := r.db.Query(context.Background(), query, limit, offset)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.Username, &user.Balance, &user.AffiliateID); err != nil {
			return nil, 0, err
		}
		users = append(users, user)
	}

	err = r.db.QueryRow(context.Background(), `SELECT COUNT(*) FROM users`).Scan(&total)
	if err != nil {
		return nil, 0, err
	}

	return users, total, nil
}

func (r *UserRepository) GetUserByID(id string) (*models.User, error) {
	var user models.User
	query := `SELECT id, username, balance, affiliate_id FROM users WHERE id = $1`
	err := r.db.QueryRow(context.Background(), query, id).Scan(&user.ID, &user.Username, &user.Balance, &user.AffiliateID)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) UpdateUserBalance(id string, amount float64) error {
	query := `UPDATE users SET balance = balance + $1 WHERE id = $2`
	_, err := r.db.Exec(context.Background(), query, amount, id)
	return err
}

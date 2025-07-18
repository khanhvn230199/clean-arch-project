package repository

import (
	"clean-arch-project/internal/domain/entity"
	"clean-arch-project/internal/domain/repository"
	"context"
	"database/sql"

	"github.com/google/uuid"
)

type userRepositoryImpl struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) repository.UserRepository {
	return &userRepositoryImpl{db: db}
}

func (r *userRepositoryImpl) Create(ctx context.Context, user *entity.User) error {
	query := `
        INSERT INTO users (id, email, name, created_at, updated_at)
        VALUES ($1, $2, $3, $4, $5)
    `

	_, err := r.db.ExecContext(ctx, query, user.ID, user.Email, user.Name, user.CreatedAt, user.UpdatedAt)
	return err
}

func (r *userRepositoryImpl) GetByID(ctx context.Context, id uuid.UUID) (*entity.User, error) {
	query := `
        SELECT id, email, name, created_at, updated_at
        FROM users
        WHERE id = $1
    `

	user := &entity.User{}
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&user.ID, &user.Email, &user.Name, &user.CreatedAt, &user.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *userRepositoryImpl) GetByEmail(ctx context.Context, email string) (*entity.User, error) {
	query := `
        SELECT id, email, name, created_at, updated_at
        FROM users
        WHERE email = $1
    `

	user := &entity.User{}
	err := r.db.QueryRowContext(ctx, query, email).Scan(
		&user.ID, &user.Email, &user.Name, &user.CreatedAt, &user.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *userRepositoryImpl) Update(ctx context.Context, user *entity.User) error {
	query := `
        UPDATE users
        SET name = $1, updated_at = $2
        WHERE id = $3
    `

	_, err := r.db.ExecContext(ctx, query, user.Name, user.UpdatedAt, user.ID)
	return err
}

func (r *userRepositoryImpl) Delete(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM users WHERE id = $1`

	_, err := r.db.ExecContext(ctx, query, id)
	return err
}

func (r *userRepositoryImpl) GetAll(ctx context.Context) ([]*entity.User, error) {
	query := `
        SELECT id, email, name, created_at, updated_at
        FROM users
        ORDER BY created_at DESC
    `

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*entity.User
	for rows.Next() {
		user := &entity.User{}
		err := rows.Scan(&user.ID, &user.Email, &user.Name, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

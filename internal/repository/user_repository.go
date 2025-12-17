package repository

import (
	"context"
	"time"

	"github.com/user-api/db/sqlc"
)

type UserRepository struct {
	db *sqlc.Queries
}

func NewUserRepository(q *sqlc.Queries) *UserRepository {
	return &UserRepository{db: q}
}

// create new user
func (r *UserRepository) Create(ctx context.Context, name string, dob time.Time) (sqlc.User, error) {
	user, err := r.db.CreateUser(ctx, name, dob)
	return user, err
}

// get user by id
func (r *UserRepository) GetByID(ctx context.Context, id int32) (sqlc.User, error) {
	user, err := r.db.GetUserByID(ctx, id)
	return user, err
}

// get all users
func (r *UserRepository) List(ctx context.Context) ([]sqlc.User, error) {
	users, err := r.db.ListUsers(ctx)
	return users, err
}

// update user
func (r *UserRepository) Update(ctx context.Context, id int32, name string, dob time.Time) (sqlc.User, error) {
	user, err := r.db.UpdateUser(ctx, name, dob, id)
	return user, err
}

// delete user
func (r *UserRepository) Delete(ctx context.Context, id int32) error {
	err := r.db.DeleteUser(ctx, id)
	return err
}

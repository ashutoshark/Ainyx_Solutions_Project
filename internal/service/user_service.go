package service

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/user-api/internal/models"
	"github.com/user-api/internal/repository"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(r *repository.UserRepository) *UserService {
	return &UserService{repo: r}
}

// create user
func (s *UserService) Create(ctx context.Context, req models.CreateUserRequest) (*models.UserResponse, error) {
	// parse date
	dob, err := time.Parse("2006-01-02", req.DOB)
	if err != nil {
		return nil, errors.New("invalid date format use YYYY-MM-DD")
	}

	// save to db
	user, err := s.repo.Create(ctx, req.Name, dob)
	if err != nil {
		log.Println("error creating user:", err)
		return nil, errors.New("failed to create user")
	}

	log.Println("user created id:", user.ID)

	// return response
	res := &models.UserResponse{
		ID:   user.ID,
		Name: user.Name,
		DOB:  user.Dob.Format("2006-01-02"),
	}
	return res, nil
}

// get user by id
func (s *UserService) GetByID(ctx context.Context, id int32) (*models.UserWithAge, error) {
	user, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, errors.New("user not found")
	}

	// make response with age
	res := &models.UserWithAge{
		ID:   user.ID,
		Name: user.Name,
		DOB:  user.Dob.Format("2006-01-02"),
		Age:  models.CalculateAge(user.Dob),
	}
	return res, nil
}

// list all users
func (s *UserService) List(ctx context.Context) ([]models.UserWithAge, error) {
	users, err := s.repo.List(ctx)
	if err != nil {
		return nil, errors.New("failed to get users")
	}

	// convert to response
	var result []models.UserWithAge
	for i := 0; i < len(users); i++ {
		u := models.UserWithAge{
			ID:   users[i].ID,
			Name: users[i].Name,
			DOB:  users[i].Dob.Format("2006-01-02"),
			Age:  models.CalculateAge(users[i].Dob),
		}
		result = append(result, u)
	}

	return result, nil
}

// update user
func (s *UserService) Update(ctx context.Context, id int32, req models.UpdateUserRequest) (*models.UserResponse, error) {
	// parse date
	dob, err := time.Parse("2006-01-02", req.DOB)
	if err != nil {
		return nil, errors.New("invalid date format use YYYY-MM-DD")
	}

	// update in db
	user, err := s.repo.Update(ctx, id, req.Name, dob)
	if err != nil {
		return nil, errors.New("user not found")
	}

	log.Println("user updated id:", user.ID)

	res := &models.UserResponse{
		ID:   user.ID,
		Name: user.Name,
		DOB:  user.Dob.Format("2006-01-02"),
	}
	return res, nil
}

// delete user
func (s *UserService) Delete(ctx context.Context, id int32) error {
	// check if exists
	_, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return errors.New("user not found")
	}

	// delete
	err = s.repo.Delete(ctx, id)
	if err != nil {
		return errors.New("failed to delete user")
	}

	log.Println("user deleted id:", id)
	return nil
}

package usecase

import (
	"clean-arch-project/internal/domain/entity"
	"clean-arch-project/internal/domain/repository"
	"clean-arch-project/internal/domain/service"
	"context"
	"errors"

	"github.com/google/uuid"
)

type UserUseCase struct {
	userRepo    repository.UserRepository
	userService *service.UserService
}

func NewUserUseCase(userRepo repository.UserRepository) *UserUseCase {
	return &UserUseCase{
		userRepo:    userRepo,
		userService: service.NewUserService(),
	}
}

func (uc *UserUseCase) CreateUser(ctx context.Context, email, name string) (*entity.User, error) {
	// Check if user already exists
	existingUser, _ := uc.userRepo.GetByEmail(ctx, email)
	if existingUser != nil {
		return nil, errors.New("user with this email already exists")
	}

	// Sanitize input
	name = uc.userService.SanitizeName(name)

	// Create new user
	user := entity.NewUser(email, name)

	// Validate user
	if err := uc.userService.ValidateUser(user); err != nil {
		return nil, err
	}

	// Save to repository
	if err := uc.userRepo.Create(ctx, user); err != nil {
		return nil, err
	}

	return user, nil
}

func (uc *UserUseCase) GetUserByID(ctx context.Context, id uuid.UUID) (*entity.User, error) {
	user, err := uc.userRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, errors.New("user not found")
	}

	return user, nil
}

func (uc *UserUseCase) UpdateUser(ctx context.Context, id uuid.UUID, name string) (*entity.User, error) {
	user, err := uc.userRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, errors.New("user not found")
	}

	// Sanitize input
	name = uc.userService.SanitizeName(name)

	// Update user
	user.Update(name)

	// Validate user
	if err := uc.userService.ValidateUser(user); err != nil {
		return nil, err
	}

	// Save to repository
	if err := uc.userRepo.Update(ctx, user); err != nil {
		return nil, err
	}

	return user, nil
}

func (uc *UserUseCase) DeleteUser(ctx context.Context, id uuid.UUID) error {
	user, err := uc.userRepo.GetByID(ctx, id)
	if err != nil {
		return err
	}

	if user == nil {
		return errors.New("user not found")
	}

	return uc.userRepo.Delete(ctx, id)
}

func (uc *UserUseCase) GetAllUsers(ctx context.Context) ([]*entity.User, error) {
	return uc.userRepo.GetAll(ctx)
}

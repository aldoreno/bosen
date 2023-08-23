package user

import (
	"bosen/manifest"
	"bosen/pkg/domain"
	errs "bosen/pkg/errors"
	"context"
	"errors"

	"go.opentelemetry.io/otel"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var _ UserRepository = (*UserRepositoryImpl)(nil)

type (
	UserRepository interface {
		FindOne(context.Context, FindCriteria, *domain.UserModel) error
	}

	UserRepositoryImpl struct {
		db *gorm.DB
	}

	FindCriteria struct {
		Username string
	}
)

func NewUserRepositoryImpl(db *gorm.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{db}
}

func (c FindCriteria) Map() map[string]any {
	output := make(map[string]any)

	if c.Username != "" {
		output["username"] = c.Username
	}

	return output
}

func (r *UserRepositoryImpl) FindOne(ctx context.Context, criteria FindCriteria, user *domain.UserModel) error {
	_, span := otel.Tracer(manifest.AppName).Start(ctx, "UserRepository.FindOne")
	defer span.End()

	result := r.db.Model(user).Where(criteria.Map()).First(user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return errs.ErrUserNotFound
		}

		zap.S().Errorf("userRepo.FindOne error: %w", result.Error)
		return errs.WrapDbError(result.Error)
	}

	return nil
}

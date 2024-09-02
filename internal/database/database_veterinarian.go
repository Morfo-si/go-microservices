package database

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/morfo-si/go-microservices/internal/dberrors"
	"github.com/morfo-si/go-microservices/internal/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (c Client) GetAllVeterinarians(ctx context.Context) ([]models.Veterinarian, error) {
	var veterinarians []models.Veterinarian
	result := c.DB.WithContext(ctx).
		Find(&veterinarians)
	return veterinarians, result.Error
}

func (c Client) AddVeterinarian(ctx context.Context, veterinarian *models.Veterinarian) (*models.Veterinarian, error) {
	veterinarian.VeterinarianID = uuid.NewString()
	result := c.DB.WithContext(ctx).
		Create(&veterinarian)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
			return nil, &dberrors.ConflictError{}
		}
		return nil, result.Error
	}
	return veterinarian, nil
}

func (c Client) GetVeterinarianById(ctx context.Context, ID string) (*models.Veterinarian, error) {
	veterinarian := &models.Veterinarian{}
	result := c.DB.WithContext(ctx).
		Where(&models.Veterinarian{VeterinarianID: ID}).
		First(veterinarian)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, &dberrors.NotFoundError{Entity: "veterinarian", ID: ID}
		}
		return nil, result.Error
	}
	return veterinarian, nil
}

func (c Client) UpdateVeterinarian(ctx context.Context, veterinarian *models.Veterinarian) (*models.Veterinarian, error) {
	var veterinarians []models.Veterinarian
	result := c.DB.WithContext(ctx).
		Model(&veterinarians).
		Clauses(clause.Returning{}).
		Where(&models.Veterinarian{VeterinarianID: veterinarian.VeterinarianID}).
		Updates(models.Veterinarian{
			FirstName:    veterinarian.FirstName,
			LastName: veterinarian.LastName,
			Phone:   veterinarian.Phone,
			Email:   veterinarian.Email,
			Specialty: veterinarian.Specialty,
		})
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
			return nil, &dberrors.ConflictError{}
		}
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, &dberrors.NotFoundError{Entity: "veterinarian", ID: veterinarian.VeterinarianID}
	}
	return &veterinarians[0], nil
}

func (c Client) DeleteVeterinarian(ctx context.Context, ID string) error {
	return c.DB.WithContext(ctx).Delete(&models.Veterinarian{VeterinarianID: ID}).Error
}

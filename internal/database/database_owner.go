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

func (c Client) GetAllOwners(ctx context.Context, emailAddress string) ([]models.Owner, error) {
	var owners []models.Owner
	result := c.DB.WithContext(ctx).
		Where(models.Owner{Email: emailAddress}).
		Find(&owners)
	return owners, result.Error
}

func (c Client) AddOwner(ctx context.Context, owner *models.Owner) (*models.Owner, error) {
	owner.OwnerID = uuid.NewString()
	result := c.DB.WithContext(ctx).
		Create(&owner)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
			return nil, &dberrors.ConflictError{}
		}
		return nil, result.Error
	}
	return owner, nil
}

func (c Client) GetOwnerById(ctx context.Context, ID string) (*models.Owner, error) {
	owner := &models.Owner{}
	result := c.DB.WithContext(ctx).
		Where(&models.Owner{OwnerID: ID}).
		First(owner)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, &dberrors.NotFoundError{Entity: "owner", ID: ID}
		}
		return nil, result.Error
	}
	return owner, nil
}

func (c Client) UpdateOwner(ctx context.Context, owner *models.Owner) (*models.Owner, error) {
	var owners []models.Owner
	result := c.DB.WithContext(ctx).
		Model(&owners).
		Clauses(clause.Returning{}).
		Where(&models.Owner{OwnerID: owner.OwnerID}).
		Updates(models.Owner{
			FirstName: owner.FirstName,
			LastName:  owner.LastName,
			Email:     owner.Email,
			Phone:     owner.Phone,
			Address:   owner.Address,
		})
	if result.Error != nil {
		if errors.Is(result.Error, &dberrors.ConflictError{}) {
			return nil, &dberrors.ConflictError{}
		}
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, &dberrors.NotFoundError{Entity: "owner", ID: owner.OwnerID}
	}
	return &owners[0], nil
}

func (c Client) DeleteOwner(ctx context.Context, ID string) error {
	return c.DB.WithContext(ctx).Delete(&models.Owner{OwnerID: ID}).Error
}

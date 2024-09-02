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

func (c Client) GetAllPets(ctx context.Context) ([]models.Pet, error) {
	var services []models.Pet
	result := c.DB.WithContext(ctx).
		Find(&services)
	return services, result.Error
}

func (c Client) AddPet(ctx context.Context, pet *models.Pet) (*models.Pet, error) {
	pet.PetID = uuid.NewString()
	result := c.DB.WithContext(ctx).
		Create(&pet)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
			return nil, &dberrors.ConflictError{}
		}
		return nil, result.Error
	}
	return pet, nil
}

func (c Client) GetPetById(ctx context.Context, ID string) (*models.Pet, error) {
	pet := &models.Pet{}
	result := c.DB.WithContext(ctx).
		Where(&models.Pet{PetID: ID}).
		First(pet)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, &dberrors.NotFoundError{Entity: "pet", ID: ID}
		}
		return nil, result.Error
	}
	return pet, nil
}

func (c Client) UpdatePet(ctx context.Context, pet *models.Pet) (*models.Pet, error) {
	var pets []models.Pet
	result := c.DB.WithContext(ctx).
		Model(&pets).
		Clauses(clause.Returning{}).
		Where(&models.Pet{PetID: pet.PetID}).
		Updates(models.Pet{
			Name:  pet.Name,
			Species: pet.Species,
			Breed: pet.Breed,
			Age: pet.Age,
			OwnerID: pet.OwnerID,
		})
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
			return nil, &dberrors.ConflictError{}
		}
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, &dberrors.NotFoundError{Entity: "pet", ID: pet.PetID}
	}
	return &pets[0], nil
}

func (c Client) DeletePet(ctx context.Context, ID string) error {
	return c.DB.WithContext(ctx).Delete(&models.Pet{PetID: ID}).Error
}

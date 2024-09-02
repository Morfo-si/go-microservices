package database

import (
	"context"
	"fmt"
	"time"

	"github.com/morfo-si/go-microservices/internal/configuration"
	"github.com/morfo-si/go-microservices/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	Host     string
	User     string
	Password string
	DBName   string
	Port     string
	SSLMode  string
}

type DatabaseClient interface {
	Ready() bool

	GetAllOwners(ctx context.Context, emailAddress string) ([]models.Owner, error)
	AddOwner(ctx context.Context, owner *models.Owner) (*models.Owner, error)
	GetOwnerById(ctx context.Context, ID string) (*models.Owner, error)
	UpdateOwner(ctx context.Context, owner *models.Owner) (*models.Owner, error)
	DeleteOwner(ctx context.Context, ID string) error

	GetAllAppointments(ctx context.Context, appointmentID string) ([]models.Appointment, error)
	AddAppointment(ctx context.Context, appointment *models.Appointment) (*models.Appointment, error)
	GetAppointmentById(ctx context.Context, ID string) (*models.Appointment, error)
	UpdateAppointment(ctx context.Context, appointment *models.Appointment) (*models.Appointment, error)
	DeleteAppointment(ctx context.Context, ID string) error

	GetAllPets(ctx context.Context) ([]models.Pet, error)
	AddPet(ctx context.Context, pet *models.Pet) (*models.Pet, error)
	GetPetById(ctx context.Context, ID string) (*models.Pet, error)
	UpdatePet(ctx context.Context, pet *models.Pet) (*models.Pet, error)
	DeletePet(ctx context.Context, ID string) error

	GetAllVeterinarians(ctx context.Context) ([]models.Veterinarian, error)
	AddVeterinarian(ctx context.Context, veterinarian *models.Veterinarian) (*models.Veterinarian, error)
	GetVeterinarianById(ctx context.Context, ID string) (*models.Veterinarian, error)
	UpdateVeterinarian(ctx context.Context, veterinarian *models.Veterinarian) (*models.Veterinarian, error)
	DeleteVeterinarian(ctx context.Context, ID string) error
}

type Client struct {
	DB     *gorm.DB
	Config *configuration.Config
}

func NewDatabaseClient(config *configuration.Config) (DatabaseClient, error) {
	// Load environment variables from .env file
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		config.Host,
		config.User,
		config.Password,
		config.DBName,
		config.Port,
		config.SSLMode,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NowFunc: func() time.Time {
			return time.Now().UTC()
		},
		QueryFields: true,
	})

	if err != nil {
		return nil, err
	}
	client := Client{
		DB:     db,
		Config: config,
	}
	return client, nil
}

func (c Client) Ready() bool {
	var ready string
	tx := c.DB.Raw("SELECT 1 as ready").Scan(&ready)
	if tx.Error != nil {
		return false
	}
	if ready == "1" {
		return true
	}
	return false
}

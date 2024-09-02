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

func (c Client) GetAllAppointments(ctx context.Context, appointmentID string) ([]models.Appointment, error) {
	var appointments []models.Appointment
	result := c.DB.WithContext(ctx).
		Where(models.Appointment{AppointmentID: appointmentID}).
		Find(&appointments)
	return appointments, result.Error
}

func (c Client) AddAppointment(ctx context.Context, appointment *models.Appointment) (*models.Appointment, error) {
	appointment.AppointmentID = uuid.NewString()
	result := c.DB.WithContext(ctx).
		Create(&appointment)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
			return nil, &dberrors.ConflictError{}
		}
		return nil, result.Error
	}
	return appointment, nil
}

func (c Client) GetAppointmentById(ctx context.Context, ID string) (*models.Appointment, error) {
	appointment := &models.Appointment{}
	result := c.DB.WithContext(ctx).
		Where(&models.Appointment{AppointmentID: ID}).
		First(appointment)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, &dberrors.NotFoundError{Entity: "appointment", ID: ID}
		}
		return nil, result.Error
	}
	return appointment, nil
}

func (c Client) UpdateAppointment(ctx context.Context, appointment *models.Appointment) (*models.Appointment, error) {
	var appointments []models.Appointment
	result := c.DB.WithContext(ctx).
		Model(&appointments).
		Clauses(clause.Returning{}).
		Where(&models.Appointment{AppointmentID: appointment.AppointmentID}).
		Updates(models.Appointment{
			AppointmentDate:     appointment.AppointmentDate,
			Reason:    appointment.Reason,
		})
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
			return nil, &dberrors.ConflictError{}
		}
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, &dberrors.NotFoundError{Entity: "appointment", ID: appointment.AppointmentID}
	}
	return &appointments[0], nil
}

func (c Client) DeleteAppointment(ctx context.Context, ID string) error {
	return c.DB.WithContext(ctx).Delete(&models.Appointment{AppointmentID: ID}).Error
}

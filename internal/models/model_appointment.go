package models

import "time"

type Appointment struct {
	AppointmentID   string    `gorm:"primaryKey" json:"appointmentId"`
	AppointmentDate time.Time `gorm:"type:timestamp;not null" json:"appointment_Date"`
	Reason          string    `json:"reason"`
	PetID           string    `json:"petId"`
	VeterinarianID  string    `json:"veterinarianId"`
}

package models

type Veterinarian struct {
	VeterinarianID string `gorm:"primaryKey" json:"veterinarianId"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Specialty        string `json:"specialty"`
	Phone          string `json:"phoneNumber"`
	Email          string `json:"emailAddress"`
}

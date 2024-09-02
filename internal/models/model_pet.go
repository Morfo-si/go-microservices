package models

type Pet struct {
	PetID string  `gorm:"primaryKey" json:"petId"`
	Name  string  `json:"name"`
	Species string `json:"species"`
	Breed string `json:"breed"`
	Age int32 `gorm:"type:numeric" json:"age"`
	OwnerID string `json:"ownerId"`
}

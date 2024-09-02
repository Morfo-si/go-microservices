package models

type Owner struct {
	OwnerID   string `gorm:"primaryKey" json:"ownerId"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"emailAddress"`
	Phone     string `json:"phoneNumber"`
	Address   string `json:"address"`
}

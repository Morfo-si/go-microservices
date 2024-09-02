package models

type Vendor struct {
	VendorID string `gorm:"primaryKey" json:"vendorId"`
	Name     string `json:"name"`
	Contact  string `json:"contact"`
	Phone    string `json:"phoneNumber"`
	Email    string `json:"emailAddress"`
	Address  string `json:"address"`
}

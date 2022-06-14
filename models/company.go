package models

type CompanyType string

const (
	DiamondType   CompanyType = "Diamond"
	GoldType      CompanyType = "Gold"
	SilverType    CompanyType = "Silver"
	CopperType    CompanyType = "Copper"
	CompanionType CompanyType = "CompanyType"
)

type Company struct {
	ID          int         `json:"id" gorm:"primaryKey"`
	Name        string      `json:"name"`
	Logo        string      `json:"logo"`
	Type        CompanyType `json:"companyType"`
	Description string      `json:"description"`
	PhoneNumber string      `json:"phoneNumber"`
	Email       string      `json:"email"`
	Address     string      `json:"address"`
	Website     string      `json:"website"`
	Facebook    string      `json:"facebook"`
	Linkedin    string      `json:"linkedin"`
}

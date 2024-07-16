package models

type City struct {
	BaseModel
	Name      string `gorm:"size:10;type:string;not null"`
	CountryID uint
	Country   Country `gorm:"foreignKey:CountryID"`
}

package entity

import(

	"gorm.io/gorm"
)

type User struct{
	gorm.Model
	StudentID string `gorm:"uniqueIndex" valid:"required~StudentID is required, matches(^[B]\\d{7}$)~StudentID is invalid"`
	Username string `gorm:"uniqueIndex" valid:"required~Username is required"`
	Password string `valid:"required~Password is required"`
	Email string `valid:"required~Email is required, email~Email is invalid"`
	Phone string `valid:"required~Phone is required, stringlength(10|10)~Phone is invalid"`
	
	GenderID uint `valid:"required~GenderID is required"`
	Gender Gender `gorm:"foreignKey:GenderID"`
}

type Gender struct{
	gorm.Model
	Name string
}
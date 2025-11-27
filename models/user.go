package models

type User struct {
	ID    uint   `json:"id" gorm:"primaryKey"`
	Name  string `json:"name" binding:"required"`
	Todos []Todo `json:"todos,omitempty" gorm:"foreignKey:UserID"`
}

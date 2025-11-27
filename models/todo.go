package models

type Todo struct {
	ID        uint   `json:"id" gorm:"primarykey"`
	Title     string `json:"title" binding:"required,min=2,max=200"`
	Completed bool   `json:"completed" gorm:"default:false"`
	UserID    uint   `json:"user_id"`
	User      *User  `json:"user,omitzero" gorm:"foreignKey:UserID" binding:"-"`
}

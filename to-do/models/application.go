package models

import "gorm.io/gorm"

// Helps to create the users table
type Users struct {
	gorm.Model  `json:"-"`
	UserID      string `gorm:"column:user_id; uniqueIndex; primaryKey; type:varchar;" json:"user_id" validate:"required"`
	UserName    string `gorm:"column:name; type:varchar;" json:"user_name" validate:"required"`
	Email       string `gorm:"column:email; type:varchar; " json:"email" validate:"email,required"`
	Password    string `gorm:"column:password; type:varchar;" json:"password,omitempty" validate:"required"`
	OldPassword string `gorm:"-" json:"old_password,omitempty"`
}

// Helps to create the tasks table
type Tasks struct {
	gorm.Model `json:"-"`
	TaskID     string `gorm:"column:task_id; uniqueIndex; primaryKey; type:varchar;" json:"task_id" validate:"required"`
	UsersID    string `gorm:"column:user_id; type:varchar;" json:"user_id" validate:"required"`
	Users      Users  `gorm:"references:user_id" json:"-" validate:"omitempty,uuid4"`
	Task       string `gorm:"column:task; type:varchar;" json:"task" validate:"required"`
	Active     bool   `gorm:"column:active; type:bool;" json:"active" validate:"required"`
}

// Helps to access the claims
type Claims struct {
	UUID string `json:"uuid"`
}

// Helps to get the login credentials
type Login struct {
	Name     string `json:"user_name" validate:"required_without=Email"`
	Email    string `json:"email" validate:"required_without=Name"`
	Password string `json:"password" validate:"required"`
}

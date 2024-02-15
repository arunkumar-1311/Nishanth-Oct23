package models

// Helps to create the users table
type Users struct {
	UserID      string `gorm:"column:user_id; uniqueIndex; primaryKey; type:varchar;" json:"user_id" validate:"required"`
	UserName    string `gorm:"column:name; type:varchar; unique" json:"user_name" validate:"required"`
	Email       string `gorm:"column:email; type:varchar; unique" json:"email" validate:"email,required"`
	Password    string `gorm:"column:password; type:varchar;" json:"password,omitempty" validate:"required"`
	OldPassword string `gorm:"-" json:"old_password,omitempty"`
}

// Helps to create the tasks table
type Tasks struct {
	TaskID string `gorm:"column:task_id; uniqueIndex; primaryKey; type:varchar;" json:"task_id" validate:"required"`
	Task   string `gorm:"column:task; type:varchar;" json:"task" validate:"required"`
	Active bool   `gorm:"column:active; type:bool;" json:"active" validate:"required"`
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

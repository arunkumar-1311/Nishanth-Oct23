package models

import (
	"time"

	"gorm.io/gorm"
)

type Roles struct {
	Role_ID string `gorm:"column:role_id; uniqueIndex;primaryKey; type:varchar" json:"role_id"`
	Role    string `gorm:"column:role; type:varchar;" json:"role"`
}

type Country struct {
	CountryID string `gorm:"column:country_id; uniqueIndex; primaryKey; type:varchar" json:"country_id"`
	Country   string `gorm:"column:country; uniqueIndex; type:varchar;" json:"country"`
}

type JobType struct {
	JobTypeID string `gorm:"column:job_type_id; uniqueIndex; primaryKey; type:varchar" json:"job_type_id"`
	JobType   string `gorm:"column:job_type; uniqueIndex; type:varchar" json:"job_type"`
}

type Users struct {
	UserID   string `gorm:"column:user_id; uniqueIndex; primaryKey; type:varchar;" json:"user_id" validate:"required"`
	UserName string `gorm:"column:name; type:varchar;" json:"user_name" validate:"required"`
	Email    string `gorm:"column:email; type:varchar;" json:"email" validate:"email,required"`
	Password string `gorm:"column:password; type:varchar;" json:"password,omitempty" validate:"required"`
	RolesID  string `gorm:"column:role_id; type:varchar;" json:"role_id" validate:"required"`
	Roles    Roles  `gorm:"references:role_id" validate:"omitempty,uuid4" json:"-"`
}

type Post struct {
	gorm.Model
	PostID      string  `gorm:"column:post_id; uniqueIndex; primaryKey; type:varchar;" json:"post_id" validate:"required"`
	UsersID     string  `gorm:"column:user_id; type:varchar;" json:"user_id" validate:"required"`
	Users       Users   `gorm:"foriegnKey:UserID;references:user_id" validate:"omitempty,uuid4" json:"user"`
	CompanyName string  `gorm:"column:company_name; type:varchar;" json:"company_name" validate:"required"`
	JobTitle    string  `gorm:"column:job_title; type:varchar;" json:"job_title" validate:"required"`
	Website     string  `gorm:"column:website; type:varchar;" json:"website" validate:"required"`
	JobTypeID   string  `gorm:"column:job_type; type:varchar;" json:"job_type" validate:"required"`
	JobType     JobType `gorm:"references:job_type_id" validate:"omitempty,uuid4" json:"jobtype"`
	CountryID   string  `gorm:"column:country_id; type:varchar;" json:"country_id" validate:"required"`
	Country     Country `gorm:"references:country_id" validate:"omitempty,uuid4" json:"country"`
	Email       string  `gorm:"column:contact_mail; type:varchar;" json:"contact_mail" validate:"required"`
	Description string  `gorm:"column:description; type:varchar;" json:"description" validate:"required"`
}

type Comment struct {
	gorm.Model
	CommentID string `gorm:"comment_id; uniqueIndex; primaryKey; type:varchar;" json:"comment_id" validate:"required"`
	PostID    string `gorm:"post_id; type:varchar;" json:"post_id" validate:"required"`
	Post      Post   `gorm:"references:post_id" validate:"omitempty,uuid4"`
	UsersID   string `gorm:"user_id; type:varchar;" json:"user_id" validate:"required"`
	Users     Users  `gorm:"references:user_id" validate:"omitempty,uuid4"`
	Content   string `gorm:"column:comment; type:varchar;" json:"comment" validate:"required"`
}

type Login struct {
	Name     string `json:"name" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type Claims struct {
	Name    string `json:"username"`
	Email   string `json:"email"`
	UsersID string `json:"user_id"`
	Role    string `json:"role"`
	RoleID  string `json:"role_id"`
}

type PostResponse struct {
	Published   time.Time `json:"CreatedAt" gorm:"column:created_at"`
	PostID      string    `json:"post_id" gorm:"column:post_id;"`
	UserID      string    `json:"user_id" gorm:"column:user_id;"`
	CompanyName string    `json:"company_name" gorm:"column:company_name;"`
	JobTitle    string    `json:"job_title" gorm:"column:job_title;"`
	Website     string    `json:"website" gorm:"column:website;"`
	JobType     JobType   `json:"jobtype" gorm:"column:job_type"`
	Country     Country   `json:"country" gorm:"column:country"`
	Email       string    `json:"contact_mail" gorm:"column:contact_mail"`
	Description string    `json:"description" gorm:"column:description"`
}

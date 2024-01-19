package models

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	PostID       string         `gorm:"column:post_id; uniqueIndex;primaryKey; type:varchar" json:"post_id"  validate:"required"`
	Title        string         `gorm:"column:title; type:varchar" json:"title"  validate:"required"`
	Content      string         `gorm:"column:content; type:varchar" json:"content"  validate:"required"`
	Excerpt      string         `gorm:"column:excerpt; type:varchar" json:"excerpt"  validate:"required"`
	Status       string         `gorm:"column:status; type:varchar" json:"status" validate:"required"`
	CategoryID   pq.StringArray `gorm:"column:category_id; type:varchar[]" json:"category_id"  validate:"required"`
	Comments     int            `gorm:"-" json:"comments"`
	PostComments []Comments     `json:"post_comments" gorm:"-"`
}

type Category struct {
	CategoryID  string `gorm:"column:category_id; uniqueIndex;primaryKey; type:varchar" json:"category_id" validate:"required"`
	Name        string `gorm:"column:name; type:varchar" json:"category_name" validate:"required"`
	Description string `gorm:"column:description; type:varchar" json:"description" validate:"required"`
}

type Roles struct {
	RoleID string `gorm:"column:role_id; uniqueIndex;primaryKey; type:varchar" json:"role_id"`
	Role   string `gorm:"column:role; type:varchar" json:"role"`
}

type Users struct {
	UserID   string `gorm:"column:user_id; uniqueIndex;primaryKey; type:varchar" json:"user_id" validate:"required"`
	Email    string `gorm:"column:email; type:varchar" json:"email" validate:"email,required"`
	Name     string `gorm:"column:name; type:varchar" json:"name" validate:"required"`
	Password string `gorm:"column:password; type:varchar" json:"password,omitempty" validate:"required"`
	RolesID  string `gorm:"column:role_id; type :varchar" json:"-"`
	Roles    Roles  `gorm:"references:role_id" json:"-"`
}

type Comments struct {
	gorm.Model
	CommentID string `gorm:"column:comment_id; type:varchar; uniqueIndex;primaryKey;" json:"comment_id" validate:"required"`
	Content   string `gorm:"column:content; type:varchar" json:"content" validate:"required"`
	Website   string `gorm:"column:source; type:varchar" json:"source" validate:"required"`
	UsersID   string `gorm:"column:user_id; type:varchar" json:"user_id" validate:"required"`
	Users     Users  `gorm:"references:user_id;" json:"user" validate:"omitempty,uuid4"`
	PostID    string `gorm:"column:post_id; type:varchar" json:"post_id" validate:"required"`
	Post      Post   `gorm:"references:post_id" json:"-" validate:"omitempty,uuid4"`
}

// These type struct's are used to send and recieve data to the client
type Login struct {
	Name     string `json:"name" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type Claims struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	UsersID string `json:"user_id"`
	RolesID string `json:"role"`
}

type CategoriesCount struct {
	CategoryName string `json:"name"`
	Total        int    `json:"total"`
}

type AllPost struct {
	Post            []Post            `json:"posts" gorm:"posts"`
	CategoriesCount []CategoriesCount `json:"categories"`
	Archieves       []string          `json:"archieves"`
}

type Filter struct {
	Date     string `json:"archieve"`
	Category string `json:"categories"`
}

type Overview struct {
	TotalPost     int64  `json:"total_posts"`
	TotalComments int64  `json:"total_comments"`
	OldestPost    string `json:"first_post"`
}

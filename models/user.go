package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"unique" form:"username" binding:"required"`
	Password string `binding:"required" form:"password"`
}

// type User struct {
// 	gorm.Model
// 	// ID       uint   `json:"id" gorm:"primary_key"` // Add an ID field to the User model
// 	UserID   uint   `json:"user_id" gorm:"primary_key"`  // Add an ID field to the User model
// 	Username string `json:"username" binding:"required"` // Add a username field to the User model
// 	Password string `json:"password" binding:"required"` // Add a password field to the User model
// 	//UserRole can only be "user" or "admin" enum("user", "admin") default("user")
// 	UserRole string `json:"user_role" gorm:"type:ENUM('user', 'admin');default:'user'"` // Add a user_role field to the User model

// }

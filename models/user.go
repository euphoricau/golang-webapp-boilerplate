package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username    string `gorm:"unique" form:"username" binding:"required"`                     // Add a username field to the User model
	Password    string `binding:"required" form:"password"`                                   // Add a password field to the User model
	Email       string `gorm:"unique" form:"email" binding:"required"`                        // Add a email field to the User model
	UserRole    string `gorm:"type:text;default:'user';check:user_role IN ('user', 'admin')"` // Use TEXT with a check constraint
	IsActive    bool   `gorm:"default:true"`                                                  // Add a is_active field to the User model
	IsConfirmed bool   `gorm:"default:false"`                                                 // Add a is_confirmed field to the User model
	Is2FA       bool   `gorm:"default:false"`                                                 // Add a is_2fa field to the User model
}

type LoginRequest struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}

const (
	USER_ROLE_USER  = "user"
	USER_ROLE_ADMIN = "admin"
)

package model

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

//User model
type User struct {
	ID        int64      `json:"id"`
	Email     *string    `json:"email"`
	Name      *string    `json:"name"`
	Gender    *string    `json:"gender"`
	Dob       *time.Time `json:"dob"`
	Password  *string    `json:"password"`
	CreatedAt *time.Time `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt"`
}

// UserUpdatePayload used
type UserUpdatePayload struct {
	ID        *int64     `json:"id"`
	Email     *string    `json:"email"`
	Name      *string    `json:"name"`
	Gender    *string    `json:"gender"`
	Dob       *time.Time `json:"dob"`
	Password  *string    `json:"password"`
	CreatedAt *time.Time `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt"`
}

// LoginPayload struct
type LoginPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

//LoginResponsePayload pay
type LoginResponsePayload struct {
	Token    string `json:"token"`
	UserData User   `json:"userData"`
}

// AuthData struct
type AuthData struct {
	UserID    int64  `json:"userId"`
	UserEmail string `json:"userEmail"`
	jwt.StandardClaims
}

// UserQuery used
type UserQuery struct {
	Page         *int64  `form:"page"`
	ItemsPerPage *int64  `form:"itemsPerPage"`
	SortBy       *string `form:"sortBy"`
	SortOrder    *string `form:"sortOrder"`
}

// id	email	name	gender	dob	password	created_at	updated_at

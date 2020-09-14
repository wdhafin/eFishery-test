package schema

import "github.com/dgrijalva/jwt-go"

// RegisterRequest is
type RegisterRequest struct {
	Phone string `json:"phone" validate:"required,number"`
	Name  string `json:"name" validate:"required"`
	Role  string `json:"role" validate:"required"`
}

// LoginRequest is
type LoginRequest struct {
	Phone    string `json:"phone" validate:"required,number"`
	Password string `json:"password" validate:"required"`
}

// LoginResponse is
type LoginResponse struct {
	Access  string `json:"access"`
	Refresh string `json:"refresh"`
	TTL     string `json:"ttl"`
}

// Token holds information about token that usable by client
type Token struct {
	Access  string `json:"access"`
	Refresh string `json:"refresh"`
	TTL     string `json:"ttl"`
}

// JWTExtract is
type JWTExtract struct {
	Name      string `json:"name"`
	Phone     string `json:"phone"`
	Role      string `json:"role"`
	Timestamp int64  `json:"timestamp"`
}

// JWTMyClaims is
type JWTMyClaims struct {
	Name      string `json:"name"`
	Phone     string `json:"phone"`
	Role      string `json:"role"`
	Timestamp int64  `json:"timestamp"`
	jwt.StandardClaims
}

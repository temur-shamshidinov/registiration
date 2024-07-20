package models

import "github.com/google/uuid"

type SignUp struct {
	Username string `json:"username"`
	Gmail    string `json:"gmail"`
	Password string `json:"password"`
}

type CheckOtp struct {
	Otp   string `json:"otp"`
	Gmail string `json:"gmail"`
}

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserAccount struct {
	UserID   uuid.UUID `json:"user_id"`
	Username string    `json:"username"`
	Gmail    string    `json:"gmail"`
	Password string    `json:"password"`
}

type RespLogIn struct {
	Token string `json:"token"`
}


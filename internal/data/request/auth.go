package request

import "github.com/google/uuid"

type LoginRequest struct {
	UserID uuid.UUID `json:"user_id"`
	IP     string    `json:"ip"`
}

type UpdateTokenRequest struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	IP           string `json:"ip"`
}

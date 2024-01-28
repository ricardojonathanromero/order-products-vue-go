package entities

type RefreshTokenReq struct {
	Token string `json:"token" validate:"required"`
}

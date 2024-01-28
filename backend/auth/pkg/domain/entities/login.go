package entities

type LoginReq struct {
	Username string `json:"username" validate:"required,min=5,max=50"`
	Password string `json:"password" validate:"required,min=8,max=30"`
}

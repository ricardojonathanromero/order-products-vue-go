package entities

type Token struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	Type         string `json:"type"`
	ExpiresIn    int64  `json:"expires_in"`
}

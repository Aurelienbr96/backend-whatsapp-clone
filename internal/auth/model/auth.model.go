package model

type Auth struct {
	RefreshToken string `json:"refreshToken"`
	AccessToken  string `json:"accessToken"`
}

type AuthPayload struct {
	Sub string
}

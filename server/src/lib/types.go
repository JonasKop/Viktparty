package lib

type Weight struct {
	UserID string  `json:"userID"`
	Weight float32 `json:"weight"`
}

type NewWeight struct {
	Weight float32 `json:"weight"`
}

type OIDCConfig struct {
	Issuer   string
	Audience string
}

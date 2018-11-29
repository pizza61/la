package api

type PytankaResponse struct {
	Ro           []ROdpowiedz `json:"ro"`
	CaptchaToken string       `json:"captchatoken"`
}

type CaptchaResponse struct {
	Success  bool    `json:"success"`
	Hostname string  `json:"hostname"`
	Score    float32 `json:"score"`
}

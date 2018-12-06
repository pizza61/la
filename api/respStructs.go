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

type GetPytankaResponse struct {
	GG      Dane   `json:"gg"`
	Captcha string `json:"captcha"`
}

type Config struct {
	CaptchaSite    string `json:"captchaSite"`
	CaptchaSecret  string `json:"captchaSecret"`
	GoogleClientID string `json:"googleClientID"`
	GoogleSecret   string `json:"googleSecret"`
}

type GoogleData struct {
	ID            string `json:"id"`
	Email         string `json:"email"`
	VerifiedEmail bool   `json:"verified_email"`
	Name          string `json:"name"`
	Pic           string `json:"picture"`
}

package event

type RegisterOtpKafka struct {
	Name         string `json:"name"`
	Otp          string `json:"otp"`
	MobileNumber string `json:"mobileNumber"`
}

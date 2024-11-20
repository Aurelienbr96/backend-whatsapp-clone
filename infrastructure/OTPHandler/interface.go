package otphandler

type OTPHandler interface {
	SendOTP(to string) error
	VerififyOTP(phoneNumber string, otp string) error
}

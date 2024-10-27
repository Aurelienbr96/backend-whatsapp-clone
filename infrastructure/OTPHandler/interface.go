package otphandler

type OTPHandler interface {
	SendOTP(to string) error
	VerififyOTP(otp string) error
}

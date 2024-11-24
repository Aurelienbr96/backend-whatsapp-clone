package otphandler

import "github.com/stretchr/testify/mock"

type OTPHandler interface {
	SendOTP(to string) error
	VerifyOTP(phoneNumber string, otp string) error
}

// MockOTPHandler is a mock implementation of the OTPHandler interface.
type MockOTPHandler struct {
	mock.Mock
}

// SendOTP is a mock method for the OTPHandler SendOTP method.
func (m *MockOTPHandler) SendOTP(to string) error {
	args := m.Called(to)
	return args.Error(0)
}

// VerifyOTP is a mock method for the OTPHandler VerifyOTP method.
func (m *MockOTPHandler) VerifyOTP(phoneNumber string, otp string) error {
	args := m.Called(phoneNumber, otp)
	return args.Error(0)
}

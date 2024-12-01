package queue

const CreatedUserSuccessType = "CreatedUserSuccess"

type CreatedUserSuccessPayload struct {
	PhoneNumber string `json:"phoneNumber"`
}

type CreatedUserSuccess struct {
	Type    string                    `json:"type"`
	Payload CreatedUserSuccessPayload `json:"payload"`
}

func NewCreatedUserSuccessMessage(phoneNumber string) *CreatedUserSuccess {
	return &CreatedUserSuccess{
		Type:    CreatedUserSuccessType,
		Payload: CreatedUserSuccessPayload{PhoneNumber: phoneNumber},
	}
}

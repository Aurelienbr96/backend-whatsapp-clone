package usecase

import (
	"github.com/google/uuid"
)

type SyncContactUseCase struct {
	contactRepository IContactRepository
	userRepository    IUserRepository
}

func NewSyncContactUseCase(contactRepository IContactRepository, userRepository IUserRepository) *SyncContactUseCase {
	return &SyncContactUseCase{contactRepository: contactRepository, userRepository: userRepository}
}

type SyncContactUseCaseInput struct {
	OwnerId      string
	PhoneNumbers []string
}

func (s *SyncContactUseCase) Execute(input SyncContactUseCaseInput) error {
	var contactsUuid []uuid.UUID
	contactMap := make(map[string]bool)

	ownerUuid, _ := uuid.Parse(input.OwnerId)

	contactUsers, _ := s.contactRepository.GetContactsByOwnerId(ownerUuid)
	for _, c := range contactUsers {
		contactMap[c.ID] = true
	}

	users, _ := s.userRepository.FindManyByPhoneNumbers(input.PhoneNumbers)

	for _, u := range users {
		if _, exists := contactMap[u.ID.String()]; !exists {
			contactsUuid = append(contactsUuid, u.ID)
		}
	}

	err := s.contactRepository.CreateMany(contactsUuid, ownerUuid)
	return err
}

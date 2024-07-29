package fake

import "fmt"

func NewFakeHashService() *FakeHashService {
	return &FakeHashService{}
}

type FakeHashService struct{}

func (s *FakeHashService) GenerateHash(value string) (string, error) {
	return fmt.Sprintf("%s-hash", value), nil
}

func (s *FakeHashService) VerifyHash(value, hashedValue string) bool {
	return fmt.Sprintf("%s-hash", value) == hashedValue
}

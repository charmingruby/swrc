package bcrypt

import "golang.org/x/crypto/bcrypt"

func NewBcryptService() *BcryptService {
	return &BcryptService{}
}

type BcryptService struct{}

func (h *BcryptService) GenerateHash(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 6)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func (h *BcryptService) VerifyHash(password, passwordHash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password))
	return err == nil
}

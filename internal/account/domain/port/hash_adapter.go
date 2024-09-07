package port

type HashPort interface {
	GenerateHash(value string) (string, error)
	VerifyHash(value, hashedValue string) bool
}

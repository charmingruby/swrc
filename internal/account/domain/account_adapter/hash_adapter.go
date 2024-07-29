package account_adapter

type HashAdapter interface {
	GenerateHash(value string) (string, error)
	VerifyHash(value, hashedValue string) bool
}

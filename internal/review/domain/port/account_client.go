package port

type AccountClient interface {
	AccountExists(accountID string) bool
	ValidAccountExists(accountID string) error
	ValidAccountExistsAndMatchRole(accountID, role string) error
}

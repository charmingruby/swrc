package review_adapter

type AccountClient interface {
	AccountExists(accountID string) bool
	IsTheAccountRole(accountID, role string) bool
}

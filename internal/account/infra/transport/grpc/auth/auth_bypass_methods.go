package auth

var (
	AccountMethodsToBypass = map[string]bool{
		"/proto.AccountService/Register":     true,
		"/proto.AccountService/Authenticate": true,
	}
)

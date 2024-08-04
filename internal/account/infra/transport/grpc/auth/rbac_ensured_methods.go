package auth

import "github.com/charmingruby/swrc/internal/account/domain/entity"

var (
	AccountRBACEnsuredMethods map[string][]string = map[string][]string{
		"/proto.AccountService/VerifyAccount":     {entity.ACCOUNT_ROLE_MANAGER},
		"/proto.AccountService/ManageAccountRole": {entity.ACCOUNT_ROLE_MANAGER},
	}
)

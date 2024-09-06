package account_service

import (
	"github.com/charmingruby/swrc/internal/account/domain/usecase"
	"github.com/charmingruby/swrc/internal/common/infra/security"
	"github.com/charmingruby/swrc/proto/pb"
)

func NewAccountGRPCService(
	accountService usecase.AccountUseCase,
	tokenService security.TokenService) *AccountGRPCService {
	return &AccountGRPCService{
		accountService: accountService,
		tokenService:   tokenService,
	}
}

type AccountGRPCService struct {
	pb.UnimplementedAccountServiceServer
	accountService usecase.AccountUseCase
	tokenService   security.TokenService
}

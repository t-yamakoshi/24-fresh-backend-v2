package resolver

import "github.com/t-yamakoshi/24-fresh-backend-v2/pkg/usecase"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{
	UserUsecase usecase.IFUserUsecase
}

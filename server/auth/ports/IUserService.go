package ports

import "github.com/StrikerSK/go-grpc/proto/auth"

type IUserService interface {
	RegisterUser(request *auth.RegisterRequest) error
}

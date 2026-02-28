package service

type AuthService interface {
	Validate(token string) bool
}

type BasicAuthService struct {
}

func (s *BasicAuthService) Validate(token string) bool {
	if token == "" {
		return false
	}
	return true
}

func NewBasicAuthService() *BasicAuthService {
	return &BasicAuthService{}
}

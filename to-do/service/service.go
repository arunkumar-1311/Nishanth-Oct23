package service

type Service struct{}

type ServiceMethods interface {
	Token
	Claims
	Password
	EmailAndNameValidation
	UUID
}

func AcquireService() ServiceMethods {
	return &Service{}
}

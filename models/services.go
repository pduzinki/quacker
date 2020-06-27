package models

type Services struct {
	Us UserService
}

func NewServices(dialect, connectionInfo string) *Services {
	return &Services{
		Us: NewUserService(dialect, connectionInfo),
	}
}

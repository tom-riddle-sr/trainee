package services

type Services struct {
	Auth    IServicesAuth
	Account IServicesAccount
}

func NewServices(auth IServicesAuth, account IServicesAccount) *Services {
	return &Services{
		Auth:    auth,
		Account: account,
	}
}

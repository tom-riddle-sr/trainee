package services

type Services struct {
	Auth    IServicesAuth
	Account IServicesAccount
	Mongo   IServicesMongo
	Redis   IServicesRedis
}

func NewServices(auth IServicesAuth, account IServicesAccount, mongo IServicesMongo, redis IServicesRedis) *Services {
	return &Services{
		Auth:    auth,
		Account: account,
		Mongo:   mongo,
		Redis:   redis,
	}
}

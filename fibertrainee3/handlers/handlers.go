package handlers

type Handlers struct {
	Auth    IHandlersAuth
	Account IHandlersAccount
	Mongo   IHandlersMongo
	Redis   IHandlersRedis
}

func NewHandlers(auth IHandlersAuth, account IHandlersAccount, mongo IHandlersMongo, redis IHandlersRedis) *Handlers {
	return &Handlers{
		Auth:    auth,
		Account: account,
		Mongo:   mongo,
		Redis:   redis,
	}
}

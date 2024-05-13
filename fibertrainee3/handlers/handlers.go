package handlers

type Handlers struct {
	Auth    IHandlersAuth
	Account IHandlersAccount
}

func NewHandlers(auth IHandlersAuth, account IHandlersAccount) *Handlers {
	return &Handlers{
		Auth:    auth,
		Account: account,
	}
}

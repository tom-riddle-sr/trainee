package input

type CreateAccountData struct {
	Account  string `validate:"required"`
	Password string `validate:"required"`
}

type UpdateAccountData struct {
	IDRequest
	CreateAccountData
}

type IDRequest struct {
	ID int32 `validate:"required"`
}

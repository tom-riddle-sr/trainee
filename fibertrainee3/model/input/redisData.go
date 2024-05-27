package input

type SetRedisData struct {
	Name  string `validate:"required"`
	Phone string `validate:"required"`
}

type RedisKeyRequest struct {
	Key string
}

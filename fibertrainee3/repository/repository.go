package repository

type Repo struct {
	RedisRepo IRedisRepo
	SqlRepo   ISqlRepo
	MongoRepo IMongoRepo
}

func NewRepo(redis IRedisRepo, sql ISqlRepo, mongo IMongoRepo) *Repo {
	return &Repo{
		RedisRepo: redis,
		SqlRepo:   sql,
		MongoRepo: mongo,
	}
}

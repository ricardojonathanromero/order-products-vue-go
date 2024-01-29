package constants

// Auth0 envs
const (
	Auth0Domain       = "AUTH0_DOMAIN"
	Auth0ClientSecret = "AUTH0_CLIENT_SECRET"
	Auth0Audience     = "AUTH0_AUDIENCE"
	Auth0Algorithm    = "AUTH0_ALGORITHM"
)

// mongodb envs
const (
	MongoDBUri        = "MONGODB_URI"
	MongoDBName       = "MONGODB_NAME"
	MongoDBCollection = "MONGODB_COLLECTION"
)

// server envs
const (
	LogLevel              = "LOG_LEVEL"
	CustomBlackList       = "CUSTOM_BLACKLIST"
	ServerPort            = "PORT"
	ServerApiKey          = "SERVER_API_KEY"
	ServerRateLimit       = "SERVER_RATE_LIMIT"
	ServerBurst           = "SERVER_BURST"
	ServerLimitExpireTime = "SERVER_LIMIT_EXPIRE_TIME"
	ServerEnableJaeger    = "SERVER_ENABLE_JAEGER"
)

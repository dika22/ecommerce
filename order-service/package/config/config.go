package config

type Config struct {
	ServerPort                string `env:"SERVER_PORT"`
	AuthHMACExpiry            string `env:"AUTH_HMAC_EXPIRY"`
	NewRelicLicense           string `env:"NEWRELIC_LICENSE"`
	AllowOrigins              string `env:"ALLOW_ORIGINS"`
	AllowHeaders              string `env:"ALLOW_HEADERS"`
	RefreshThresholdSecond    string `env:"REFRESH_THRESHOLD_SECOND"`
	CookieSecure              string `env:"COOKIE_SECURE"`
	InternalAccessKey         string `env:"INTERNAL_ACCESS_KEY"`
	RateLimitThreshold        string `env:"RATE_LIMIT_THRESHOLD"`
	RateLimitBucketLeakSecond string `env:"RATE_LIMIT_BUCKET_LEAK_SECOND"`
	DebugHTTP                 string `env:"DEBUG_HTTP"`
	BaseURLWarehouseService   string `env:"BASE_URL_WAREHOUSE_SERVICE"`
}

func NewConfig() *Config {
	c := &Config{}
	LoadEnv()
	MarshalEnv(c)

	return c
}

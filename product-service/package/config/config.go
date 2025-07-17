package config

type Config struct {
	ServerPort                string `env:"SERVER_PORT"`
	DebugHTTP                 string `env:"DEBUG_HTTP"`
	MailgunBaseURL            string `env:"MAILGUN_BASE_URL"`
	MailgunAPIKey             string `env:"MAILGUN_API_KEY"`
	RateLimitThreshold        string `env:"RATE_LIMIT_THRESHOLD"`
	RateLimitBucketLeakSecond string `env:"RATE_LIMIT_BUCKET_LEAK_SECOND"`
}

func NewConfig() *Config {
	c := &Config{}
	LoadEnv()
	MarshalEnv(c)
	if c.RateLimitThreshold == "" {
		c.RateLimitThreshold = "1000"
	}
	if c.RateLimitBucketLeakSecond == "" {
		c.RateLimitBucketLeakSecond = "1"
	}
	return c
}

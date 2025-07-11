package config

type Config struct {
	ServerPort                string `env:"SERVER_PORT"`
	DebugHTTP                 string `env:"DEBUG_HTTP"`
	RabbitMQURL               string `env:"RABBITMQ_URL"`
}

func NewConfig() *Config {
	c := &Config{}
	LoadEnv()
	MarshalEnv(c)
	return c
}

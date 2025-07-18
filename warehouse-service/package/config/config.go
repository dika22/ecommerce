package config

type Config struct {
	ServerPort                string `env:"SERVER_PORT"`
	DebugHTTP                 string `env:"DEBUG_HTTP"`
	MessageBrokerURL          string `env:"MESSAGE_BROKER_URL"`
	MessageBrokerUsername     string `env:"MESSAGE_BROKER_USERNAME"`
	MessageBrokerPassword     string `env:"MESSAGE_BROKER_PASSWORD"`
	MessageBrokerPort         string `env:"MESSAGE_BROKER_PORT"`
	MessageBrokerQueue        string `env:"MESSAGE_BROKER_QUEUE"`
}

func NewConfig() *Config {
	c := &Config{}
	LoadEnv()
	MarshalEnv(c)
	return c
}

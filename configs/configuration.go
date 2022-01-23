package config

type Configuration struct {
	Server ServerConfiguration
	Login  LoginConfig
}

type LoginConfig struct {
	Token string
}

type ServerConfiguration struct {
	Port string
}

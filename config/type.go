package config

type AppConfig struct {
	Name    string
	Version string
	Api     ApiConfig
	BpInfo  string
	Upgrade string
	Network string
}


type ApiConfig struct {
	Address string
	Port    string
	Debug   bool
}

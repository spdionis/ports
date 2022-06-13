package app

const dbURL = "postgresql://user:password@portsdb/ports"
const listenAddr = "0.0.0.0:8000"

type ConfigManager struct {
	ListenAddr string
	DbURL      string
}

func NewConfig() ConfigManager {
	return ConfigManager{
		DbURL:      dbURL,
		ListenAddr: listenAddr,
	}
}

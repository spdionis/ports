package app

const dbURL = "postgresql://user:password@portsdb/ports"
const listenAddr = "0.0.0.0:8000"
const importBatchSize = 100

type ConfigManager struct {
	ListenAddr      string
	DbURL           string
	ImportBatchSize int
}

func NewConfig() ConfigManager {
	return ConfigManager{
		DbURL:           dbURL,
		ListenAddr:      listenAddr,
		ImportBatchSize: importBatchSize,
	}
}

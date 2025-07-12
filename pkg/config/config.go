package config

type config struct {
	Password string
	Port     string
	DbFile   string
}

const defaultPort = "7540"
const defaultDbFile = "scheduler.db"
const defaultPassword = ""

func New(password, port, dbFile string) *config {
	if port == "" {
		port = defaultPort
	}

	if dbFile == "" {
		dbFile = defaultDbFile
	}

	if password == "" {
		password = defaultPassword
	}

	return &config{
		Password: password,
		Port:     port,
		DbFile:   dbFile,
	}
}

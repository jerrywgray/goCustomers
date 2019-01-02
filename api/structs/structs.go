package structs

// DBConfig holds our database config
type DBConfig struct {
	User     string
	Password string
	Database string
	Host     string
	Port     string
}

// Config holds our config that we'll get from the environment
type Config struct {
	*DBConfig
	DataFilename string
	Env          string
	Port         string
	GinMode      string
	GormMode     string
	Initialize   bool
}

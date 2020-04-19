package config

type Configurations struct {
	Server   ServerConfigurations
	Database DatabaseConfigurations
}

type ServerConfigurations struct {
	Port int
}

type DatabaseConfigurations struct {
	DBHost     string
	DBName     string
	DBUser     string
	DBPassword string
}

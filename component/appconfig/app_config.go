package appconfig

type AppConfig struct {
	Port string
	Env  string

	DBUsername string
	DBPassword string
	DBHost     string
	DBDatabase string

	SecretKey string
}

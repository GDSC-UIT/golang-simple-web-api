package appconfig

type AppConfig struct {
	Port       string
	Env        string
	StaticPath string

	DBUsername string
	DBPassword string
	DBHost     string
	DBDatabase string

	SecretKey string
}

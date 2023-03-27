package application

type DbConfig struct {
	Driver   string `envconfig:"driver"`
	User     string `envconfig:"user"`
	Password string `envconfig:"password"`
	Name     string `envconfig:"name"`
	Host     string `envconfig:"host"`
	Port     string `envconfig:"port"`
}

type Config struct {
	Database DbConfig `envconfig:"db"`
}

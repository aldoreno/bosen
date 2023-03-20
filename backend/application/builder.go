package application

type Services struct {
	dbConfig *DbConfig
}

func NewServices() *Services {
	return &Services{
		dbConfig: &DbConfig{},
	}
}

// TODO: should be able to add multiple database configurations e.g. read, write
func (s Services) AddDbConfig(fn func(config *DbConfig)) {
	fn(s.dbConfig)
}

type Builder struct {
	Services *Services
}

func NewBuilder() *Builder {
	return &Builder{
		Services: NewServices(),
	}
}

func (b *Builder) Build() *Application {
	app := NewApplication(
		WithLogger(),
		WithDatabase(func() *DbConfig {
			return b.Services.dbConfig
		}),
	)
	return app
}

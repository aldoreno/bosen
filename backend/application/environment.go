package application

type Environment struct{}

func (e Environment) IsDevelopment() bool {
	return true
}

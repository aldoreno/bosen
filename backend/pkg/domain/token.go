package domain

type Token struct {
	value string
}

func (t Token) String() string {
	return t.value
}

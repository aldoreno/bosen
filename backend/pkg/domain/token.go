package domain

type Token struct {
	Value string
}

func (t Token) String() string {
	return t.Value
}

package greeting

import "testing"

func TestGreeting(t *testing.T) {
	resultado := Greeting("teste")
	if resultado != "<b>teste</b>" {
		t.Error("O resultado deve ser <b>teste</b>", resultado)
	}

}

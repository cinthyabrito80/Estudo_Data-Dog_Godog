package soma

import (
	"fmt"

	"github.com/DATA-DOG/godog"
)

var resultado int

func faoASomaDe(n1, n2 int) error {
	resultado = n1 + n2
	return nil
}

func oResultadoDeveSer(soma int) error {
	//fmt.Println(resultado)

	if resultado != soma {
		return fmt.Errorf("Erro ao relizar a soma. Esperado: %v Obitido: %v", soma, resultado)
	}
	return nil
}

func FeatureContext(s *godog.Suite) {
	s.Step(`^faço a soma de (\d+) \+ (\d+)$`, faoASomaDe)
	s.Step(`^o resultado deve ser (\d+)$`, oResultadoDeveSer)
}

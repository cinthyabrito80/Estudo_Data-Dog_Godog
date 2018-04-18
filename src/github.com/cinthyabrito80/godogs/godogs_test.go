package godogs

import (
	"fmt"

	"github.com/DATA-DOG/godog"
)

var Godogs int

func queExistemGodogs(available int) error {
	Godogs = available
	return nil
}

func euComo(num int) error {
	if Godogs < num {
		return fmt.Errorf("you cannot eat %d godogs, there are %d available", num, Godogs)
	}
	Godogs -= num
	return nil
}

func deveRestarGodogs(remaining int) error {
	if Godogs != remaining {
		return fmt.Errorf("expected %d godogs to be remaining, but there is %d", remaining, Godogs)
	}
	return nil
}

func FeatureContext(s *godog.Suite) {
	s.Step(`^que existem (\d+) Godogs$`, queExistemGodogs)
	s.Step(`^eu como (\d+)$`, euComo)
	s.Step(`^deve restar (\d+) Godogs$`, deveRestarGodogs)

	s.BeforeScenario(func(interface{}) {
		Godogs = 0 // clean the state before every scenario
	})
}

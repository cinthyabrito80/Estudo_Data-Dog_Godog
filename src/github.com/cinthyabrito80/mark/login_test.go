package mark

import (
	"github.com/DATA-DOG/godog"
	"github.com/cinthyabrito80/mark/support"
	"github.com/tebeka/selenium"
)

var Driver selenium.WebDriver

func queEuAcesseiUmPaginaPrincipal() error {
	Driver.Get("https://marktasks.herokuapp.com")
	return nil
}

func faoLoginComE(email, senha string) (err error) {
	campoEmail, err := Driver.FindElement(selenium.ByID, "login_email")
	if err != nil {
		return
	}
	campoEmail.SendKeys(email)

	campoSenha, err := Driver.FindElement(selenium.ByID, "login_password")
	if err != nil {
		return
	}
	campoSenha.SendKeys(senha)

	botaoLogin, err := Driver.FindElement(selenium.ByID, "btnLogin")
	if err != nil {
		return
	}
	botaoLogin.Click()

	return nil
}

func souAutenticadoComSucesso() error {
	return godog.ErrPending
}

func devoVerASeguinteMensagem(arg1 string) error {
	return godog.ErrPending
}

func FeatureContext(s *godog.Suite) {
	s.Step(`^que eu acessei um pagina principal$`, queEuAcesseiUmPaginaPrincipal)
	s.Step(`^faço login com "([^"]*)" e "([^"]*)"$`, faoLoginComE)
	s.Step(`^sou autenticado com sucesso$`, souAutenticadoComSucesso)
	s.Step(`^devo ver a seguinte mensagem "([^"]*)"$`, devoVerASeguinteMensagem)

	s.BeforeScenario(func(interface{}) {
		Driver = support.WDInit()
	})
}

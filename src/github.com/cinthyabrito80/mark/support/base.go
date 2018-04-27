package support

import (
	"fmt"

	"github.com/tebeka/selenium"
)

var driver selenium.WebDriver

//WDInit retorna uma instancia do WebDriver
func WDInit() selenium.WebDriver {

	var err error
	caps := selenium.Capabilities(map[string]interface{}{"browserName": "chrome"})
	driver, err = selenium.NewRemote(caps, "")

	if err != nil {
		fmt.Println("Erro ao instanciar o driver:", err.Error())
	}

	driver.ResizeWindow("note", 1200, 800)

	return driver

}

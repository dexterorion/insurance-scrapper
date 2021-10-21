package processors

import (
	"fmt"
	"time"

	"github.com/seeker-insurance/seeker-scraper/models"
	"github.com/seeker-insurance/seeker-scraper/outside"
	"github.com/tebeka/selenium"
)

// Nationwide represents the Nationwide structure
type Nationwide struct{}

// Process processes html
func (an Nationwide) Process(zip, state, city string) []models.Agent {
	executor := &outside.SeleniumExecutor{}
	executor.StartSelenium()

	agents := make([]models.Agent, 0)

	url := fmt.Sprintf("https://agency.nationwide.com/search?q=%s", zip)

	// Navigate to the simple playground interface.
	if err := executor.WD.Get(url); err != nil {
		panic(err)
	}

	time.Sleep(time.Second * 3)

	for {
		_, err := executor.WD.FindElement(selenium.ByCSSSelector, ".SpinnerModal--visible")

		if err != nil {
			break
		}

		fmt.Println("Still loading")
		time.Sleep(time.Second)
	}

	htmlAgents, err := executor.WD.FindElements(selenium.ByCSSSelector, ".ResultList-item")
	if err != nil {
		panic(err)
	}

	for _, htmlAgent := range htmlAgents {
		agent := models.Agent{}

		name, err := htmlAgent.FindElement(selenium.ByCSSSelector, "#location-name")
		if err != nil {
			panic(err)
		}

		txtName, err := name.Text()
		if err != nil {
			panic(err)
		}

		phone, err := htmlAgent.FindElement(selenium.ByCSSSelector, "#telephone")
		if err != nil {
			panic(err)
		}

		txtPhone, err := phone.Text()
		if err != nil {
			panic(err)
		}

		address1, err := htmlAgent.FindElement(selenium.ByCSSSelector, ".c-address-street-1")
		if err != nil {
			panic(err)
		}
		txtAddress1, err := address1.Text()
		if err != nil {
			panic(err)
		}

		var txtAddress2 string
		address2, err := htmlAgent.FindElement(selenium.ByCSSSelector, ".c-address-street-2")
		if err == nil {
			txtAddress2, err = address2.Text()
			if err != nil {
				panic(err)
			}
		}

		city, err := htmlAgent.FindElement(selenium.ByCSSSelector, ".c-address-city")
		if err != nil {
			panic(err)
		}
		txtCity, err := city.Text()
		if err != nil {
			panic(err)
		}

		state, err := htmlAgent.FindElement(selenium.ByCSSSelector, ".c-address-state")
		if err != nil {
			panic(err)
		}
		txtState, err := state.Text()
		if err != nil {
			panic(err)
		}

		zip, err := htmlAgent.FindElement(selenium.ByCSSSelector, ".c-address-postal-code")
		if err != nil {
			panic(err)
		}
		txtZip, err := zip.Text()
		if err != nil {
			panic(err)
		}

		agent.Address = fmt.Sprintf("%s %s %s %s %s", txtAddress1, txtAddress2, txtCity, txtState, txtZip)
		agent.Name = txtName
		agent.Phone = txtPhone
		agent.Fax = txtPhone

		agents = append(agents, agent)
	}

	executor.StopSelenium()

	return agents
}

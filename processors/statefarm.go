package processors

import (
	"fmt"
	"time"

	"github.com/dexterorion/insurance-scraper/models"
	"github.com/dexterorion/insurance-scraper/outside"
	"github.com/tebeka/selenium"
)

// StateFarm represents the StateFarm structure
type StateFarm struct{}

// Process processes html
func (an StateFarm) Process(zip, state, city string) []models.Agent {
	executor := &outside.SeleniumExecutor{}
	executor.StartSelenium()

	agents := make([]models.Agent, 0)

	// Navigate to the simple playground interface.
	if err := executor.WD.Get("https://www.statefarm.com/agent/"); err != nil {
		panic(err)
	}

	// Get a reference to the text box containing code.
	elem, err := executor.WD.FindElement(selenium.ByID, "locationText")
	if err != nil {
		panic(err)
	}
	err = elem.SendKeys(zip)
	if err != nil {
		panic(err)
	}

	search, err := executor.WD.FindElement(selenium.ByID, "search")
	if err != nil {
		panic(err)
	}

	search.Click()

	listing, err := executor.WD.FindElement(selenium.ByID, "resultsPage")
	if err != nil {
		panic(err)
	}

	for {
		displayed, err := listing.CSSProperty("display")
		if err != nil {
			panic(err)
		}

		if displayed != "none" {
			break
		}
		time.Sleep(time.Second)
	}

	showAll, err := executor.WD.FindElement(selenium.ByCSSSelector, "a[data-show=all]")
	if err != nil {
		panic(err)
	}
	showAll.Click()

	htmlAgents, err := listing.FindElements(selenium.ByCSSSelector, ".agentRootTemplate")
	if err != nil {
		panic(err)
	}

	for _, htmlAgent := range htmlAgents {
		agent := models.Agent{}
		name, err := htmlAgent.FindElement(selenium.ByCSSSelector, ".agent-name")
		if err != nil {
			panic(err)
		}
		txtName, err := name.Text()
		if err != nil {
			panic(err)
		}

		if txtName == "" {
			continue
		}

		license, err := htmlAgent.FindElement(selenium.ByCSSSelector, ".agent-license")
		if err != nil {
			panic(err)
		}
		txtLicense, err := license.Text()
		if err != nil {
			panic(err)
		}

		address1, err := htmlAgent.FindElement(selenium.ByCSSSelector, ".agent-street-address-1")
		if err != nil {
			panic(err)
		}
		txtAddress1, err := address1.Text()
		if err != nil {
			panic(err)
		}
		address2, err := htmlAgent.FindElement(selenium.ByCSSSelector, ".agent-street-address-2")
		if err != nil {
			panic(err)
		}
		txtAddress2, err := address2.Text()
		if err != nil {
			panic(err)
		}
		cityState, err := htmlAgent.FindElement(selenium.ByCSSSelector, ".agent-city-state-zip")
		if err != nil {
			panic(err)
		}
		txtCityState, err := cityState.Text()
		if err != nil {
			panic(err)
		}

		phone, err := htmlAgent.FindElement(selenium.ByCSSSelector, ".hidden-phone .agent-contact-number")
		if err != nil {
			panic(err)
		}
		txtPhone, err := phone.Text()
		if err != nil {
			panic(err)
		}

		agent.Name = txtName
		agent.Licenses = txtLicense
		agent.Address = fmt.Sprintf("%s %s %s", txtAddress1, txtAddress2, txtCityState)
		agent.Phone = txtPhone
		agent.Fax = txtPhone

		agents = append(agents, agent)
	}

	executor.StopSelenium()

	return agents
}

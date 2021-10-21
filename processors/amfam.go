package processors

import (
	"strings"
	"time"

	"github.com/dexterorion/insurance-scraper/models"
	"github.com/dexterorion/insurance-scraper/outside"
	"github.com/tebeka/selenium"
)

// Amfam represents the Amfam structure
type Amfam struct{}

// Process processes html
func (an Amfam) Process(zip, state, city string) []models.Agent {
	executor := &outside.SeleniumExecutor{}
	executor.StartSelenium()

	agents := make([]models.Agent, 0)

	url := "https://myapps.amfam.com/amfamagentlocator/agentLocator"

	// Navigate to the simple playground interface.
	if err := executor.WD.Get(url); err != nil {
		panic(err)
	}

	time.Sleep(time.Second * 3)

	zipInput, err := executor.WD.FindElement(selenium.ByID, "zip-code")
	if err != nil {
		panic(err)
	}

	zipInput.SendKeys(zip)

	time.Sleep(time.Second * 3)

	search, err := executor.WD.FindElement(selenium.ByCSSSelector, "button[type='submit']")
	if err != nil {
		panic(err)
	}

	search.Click()

	time.Sleep(time.Second * 3)

	refineLink, err := executor.WD.FindElement(selenium.ByCSSSelector, ".refine-link")
	if err != nil {
		panic(err)
	}

	for {
		displayed, err := refineLink.CSSProperty("display")
		if err != nil {
			panic(err)
		}

		if displayed != "none" {
			break
		}
		time.Sleep(time.Second)
	}

	time.Sleep(time.Second * 2)

	htmlAgents, err := executor.WD.FindElements(selenium.ByCSSSelector, "#results-list")
	if err != nil {
		panic(err)
	}

	for _, htmlAgent := range htmlAgents {
		agent := models.Agent{}
		name, err := htmlAgent.FindElement(selenium.ByCSSSelector, ".agent-name-wrapper")
		if err != nil {
			panic(err)
		}

		txtName, err := name.Text()
		if err != nil {
			panic(err)
		}

		arrayName := strings.Split(txtName, "\n")

		phone, err := htmlAgent.FindElement(selenium.ByCSSSelector, "span.icon-call > a")
		if err != nil {
			panic(err)
		}
		txtPhone, err := phone.Text()
		if err != nil {
			panic(err)
		}

		address, err := htmlAgent.FindElement(selenium.ByCSSSelector, "span.icon-call > div")
		if err != nil {
			panic(err)
		}
		txtAddress, err := address.Text()
		if err != nil {
			panic(err)
		}

		if len(arrayName) == 2 {
			agent.Name = strings.ReplaceAll(arrayName[0], "\"", "")
		} else {
			agent.Name = strings.ReplaceAll(txtName, "\"", "")
		}
		agent.Address = strings.ReplaceAll(txtAddress, "\n", " ")
		agent.Phone = txtPhone
		agent.Fax = txtPhone

		agents = append(agents, agent)
	}

	executor.StopSelenium()

	return agents
}

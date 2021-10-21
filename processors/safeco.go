package processors

import (
	"fmt"
	"strings"
	"time"

	"github.com/dexterorion/insurance-scraper/models"
	"github.com/dexterorion/insurance-scraper/outside"
	"github.com/tebeka/selenium"
)

// Safeco represents the Safeco structure
type Safeco struct{}

// Process processes html
func (an Safeco) Process(zip, state, city string) []models.Agent {
	executor := &outside.SeleniumExecutor{}
	executor.StartSelenium()

	agents := make([]models.Agent, 0)

	url := fmt.Sprintf("https://insurance-agent.safeco.com/find-an-insurance-agency/app/search-results-locationSearch=%s", zip)

	// Navigate to the simple playground interface.
	if err := executor.WD.Get(url); err != nil {
		panic(err)
	}

	time.Sleep(time.Second * 3)

	spinner, err := executor.WD.FindElement(selenium.ByCSSSelector, "body > div:nth-child(2)")
	if err != nil {
		panic(err)
	}

	for {
		displayed, err := spinner.CSSProperty("display")

		if err != nil {
			panic(err)
		}

		if displayed == "none" {
			break
		}

		time.Sleep(time.Second)
	}

	htmlAgents, err := executor.WD.FindElements(selenium.ByCSSSelector, ".searchCard")
	if err != nil {
		panic(err)
	}
	var agentLinks []string

	for _, htmlAgent := range htmlAgents {
		link, err := htmlAgent.FindElement(selenium.ByCSSSelector, "#srpgLftAgencyName")
		if err != nil {
			panic(err)
		}

		txtLink, err := link.GetAttribute("href")
		if err != nil {
			panic(err)
		}

		agentLinks = append(agentLinks, txtLink)
	}

	paginator, err := executor.WD.FindElement(selenium.ByCSSSelector, "span[ng-click='nextPage()']")
	if err != nil {
		panic(err)
	}

	// doing four calls

	// one
	paginator.Click()
	time.Sleep(time.Second * 3)
	htmlAgents, err = executor.WD.FindElements(selenium.ByCSSSelector, ".searchCard")
	if err != nil {
		panic(err)
	}
	for _, htmlAgent := range htmlAgents {
		link, err := htmlAgent.FindElement(selenium.ByCSSSelector, "#srpgLftAgencyName")
		if err != nil {
			panic(err)
		}

		txtLink, err := link.GetAttribute("href")
		if err != nil {
			panic(err)
		}

		agentLinks = append(agentLinks, txtLink)
	}

	// two
	paginator.Click()
	time.Sleep(time.Second * 3)
	htmlAgents, err = executor.WD.FindElements(selenium.ByCSSSelector, ".searchCard")
	if err != nil {
		panic(err)
	}
	for _, htmlAgent := range htmlAgents {
		link, err := htmlAgent.FindElement(selenium.ByCSSSelector, "#srpgLftAgencyName")
		if err != nil {
			panic(err)
		}

		txtLink, err := link.GetAttribute("href")
		if err != nil {
			panic(err)
		}

		agentLinks = append(agentLinks, txtLink)
	}

	// three
	paginator.Click()
	time.Sleep(time.Second * 3)
	htmlAgents, err = executor.WD.FindElements(selenium.ByCSSSelector, ".searchCard")
	if err != nil {
		panic(err)
	}
	for _, htmlAgent := range htmlAgents {
		link, err := htmlAgent.FindElement(selenium.ByCSSSelector, "#srpgLftAgencyName")
		if err != nil {
			panic(err)
		}

		txtLink, err := link.GetAttribute("href")
		if err != nil {
			panic(err)
		}

		agentLinks = append(agentLinks, txtLink)
	}
	// four
	paginator.Click()
	time.Sleep(time.Second * 3)
	htmlAgents, err = executor.WD.FindElements(selenium.ByCSSSelector, ".searchCard")
	if err != nil {
		panic(err)
	}
	for _, htmlAgent := range htmlAgents {
		link, err := htmlAgent.FindElement(selenium.ByCSSSelector, "#srpgLftAgencyName")
		if err != nil {
			panic(err)
		}

		txtLink, err := link.GetAttribute("href")
		if err != nil {
			panic(err)
		}

		agentLinks = append(agentLinks, txtLink)
	}

	for k := range agentLinks {
		if err := executor.WD.Get(agentLinks[k]); err != nil {
			panic(err)
		}

		time.Sleep(time.Second * 3)

		spinner, err := executor.WD.FindElement(selenium.ByCSSSelector, "body > div:nth-child(2)")
		if err != nil {
			panic(err)
		}

		for {
			displayed, err := spinner.CSSProperty("display")

			if err != nil {
				panic(err)
			}

			if displayed == "none" {
				break
			}

			time.Sleep(time.Second)
		}

		agent := models.Agent{}

		name, err := executor.WD.FindElement(selenium.ByCSSSelector, "#profileTitle")
		if err != nil {
			panic(err)
		}
		txtName, err := name.Text()
		if err != nil {
			panic(err)
		}
		agent.Name = txtName

		phone, err := executor.WD.FindElement(selenium.ByCSSSelector, "#profilePhone")
		if err != nil {
			panic(err)
		}
		txtPhone, err := phone.Text()
		if err != nil {
			panic(err)
		}
		agent.Phone = txtPhone
		agent.Fax = txtPhone

		addresses, err := executor.WD.FindElements(selenium.ByCSSSelector, ".profileMapAddress")
		if err != nil {
			panic(err)
		}

		stringBAddress := new(strings.Builder)
		for _, addr := range addresses {
			address, err := addr.Text()
			if err != nil {
				panic(err)
			}
			stringBAddress.WriteString(address)
			stringBAddress.WriteString(" ")
		}
		agent.Address = stringBAddress.String()

		agents = append(agents, agent)
	}

	executor.StopSelenium()

	return agents
}

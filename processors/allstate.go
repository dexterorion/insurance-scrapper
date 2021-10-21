package processors

import (
	"fmt"
	"time"

	"github.com/seeker-insurance/seeker-scraper/models"
	"github.com/seeker-insurance/seeker-scraper/outside"
	"github.com/tebeka/selenium"
)

// AllState represents the AllState structure
type AllState struct{}

// Process processes html
func (an AllState) Process(zip, state, city string) []models.Agent {
	executor := &outside.SeleniumExecutor{}
	executor.StartSelenium()

	agents := make([]models.Agent, 0)

	url := fmt.Sprintf("https://agents.allstate.com/locator.html?search=%s&r=20&rating=0&type=4848&type=709&type=710&type=711&type=5387005", zip)

	// Navigate to the simple playground interface.
	if err := executor.WD.Get(url); err != nil {
		executor.StopSelenium()
		panic(err)
	}

	time.Sleep(time.Second * 10)

	listView, err := executor.WD.FindElement(selenium.ByCSSSelector, ".Locator-toggleView--list")
	if err != nil {
		executor.StopSelenium()
		panic(err)
	}

	listView.Click()

	time.Sleep(time.Second * 2)

	content, err := executor.WD.FindElement(selenium.ByCSSSelector, ".Locator-content")
	if err != nil {
		executor.StopSelenium()
		panic(err)
	}

	time.Sleep(time.Second * 2)

	wrapper, err := content.FindElement(selenium.ByCSSSelector, ".Locator-resultsWrapper")
	if err != nil {
		executor.StopSelenium()
		panic(err)
	}

	time.Sleep(time.Second * 2)

	container, err := wrapper.FindElement(selenium.ByCSSSelector, ".Locator-resultsContainer")
	if err != nil {
		executor.StopSelenium()
		panic(err)
	}

	time.Sleep(time.Second * 2)

	locatorResults, err := container.FindElement(selenium.ByCSSSelector, ".Locator-results")
	if err != nil {
		executor.StopSelenium()
		panic(err)
	}

	time.Sleep(time.Second * 2)

	var resultList selenium.WebElement

	counter := 0

	for {
		resultList, err = locatorResults.FindElement(selenium.ByCSSSelector, ".ResultList")
		if err == nil {
			break
		}
		counter++
		fmt.Println("Still nothing")
		time.Sleep(time.Second * 5)

		if counter == 5 {
			executor.StopSelenium()
			fmt.Println("Tried 5 times... finishing. Run it again")
			return nil
		}
	}

	for {
		if err != nil {
			break
		} else {
			displayed, err := resultList.CSSProperty("display")
			if err != nil {
				executor.StopSelenium()
				panic(err)
			}

			if displayed != "none" {
				break
			}
			time.Sleep(time.Second)
			resultList, err = executor.WD.FindElement(selenium.ByCSSSelector, ".ResultList")
		}
	}

	time.Sleep(time.Second * 2)

	htmlAgents, err := executor.WD.FindElements(selenium.ByCSSSelector, ".ResultList-item")
	if err != nil {
		executor.StopSelenium()
		panic(err)
	}

	for _, htmlAgent := range htmlAgents {
		agent := models.Agent{}
		name, err := htmlAgent.FindElement(selenium.ByCSSSelector, ".Teaser-name")
		if err != nil {
			executor.StopSelenium()
			panic(err)
		}
		txtName, err := name.Text()
		if err != nil {
			executor.StopSelenium()
			panic(err)
		}

		if txtName == "" {
			continue
		}

		address1, err := htmlAgent.FindElement(selenium.ByCSSSelector, ".c-address-street-1")
		if err != nil {
			executor.StopSelenium()
			panic(err)
		}
		txtAddress1, err := address1.Text()
		if err != nil {
			executor.StopSelenium()
			panic(err)
		}

		var txtAddress2 string
		address2, err := htmlAgent.FindElement(selenium.ByCSSSelector, ".c-address-street-2")
		if err == nil {
			txtAddress2, err = address2.Text()
			if err != nil {
				executor.StopSelenium()
				panic(err)
			}
		}

		city, err := htmlAgent.FindElement(selenium.ByCSSSelector, ".c-address-city")
		if err != nil {
			executor.StopSelenium()
			panic(err)
		}
		txtCity, err := city.Text()
		if err != nil {
			executor.StopSelenium()
			panic(err)
		}

		state, err := htmlAgent.FindElement(selenium.ByCSSSelector, ".c-address-state")
		if err != nil {
			executor.StopSelenium()
			panic(err)
		}
		txtState, err := state.Text()
		if err != nil {
			executor.StopSelenium()
			panic(err)
		}

		zip, err := htmlAgent.FindElement(selenium.ByCSSSelector, ".c-address-postal-code")
		if err != nil {
			executor.StopSelenium()
			panic(err)
		}
		txtZip, err := zip.Text()
		if err != nil {
			executor.StopSelenium()
			panic(err)
		}

		phone, err := htmlAgent.FindElement(selenium.ByCSSSelector, ".Teaser-phoneText")
		if err != nil {
			executor.StopSelenium()
			panic(err)
		}
		txtPhone, err := phone.Text()
		if err != nil {
			executor.StopSelenium()
			panic(err)
		}

		agent.Name = txtName
		agent.Address = fmt.Sprintf("%s %s %s %s %s", txtAddress1, txtAddress2, txtCity, txtState, txtZip)
		agent.Phone = txtPhone
		agent.Fax = txtPhone

		agents = append(agents, agent)
	}

	executor.StopSelenium()

	return agents
}

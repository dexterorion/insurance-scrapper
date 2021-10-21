package processors

import (
	"fmt"
	"time"

	"github.com/seeker-insurance/seeker-scraper/models"
	"github.com/seeker-insurance/seeker-scraper/outside"
	"github.com/tebeka/selenium"
)

// Liberty represents the Liberty structure
type Liberty struct{}

// Process processes html
func (an Liberty) Process(zip, state, city string) []models.Agent {
	executor := &outside.SeleniumExecutor{}
	executor.StartSelenium()

	agents := make([]models.Agent, 0)

	url := "https://www.libertymutual.com/find-sales-office"

	// Navigate to the simple playground interface.
	if err := executor.WD.Get(url); err != nil {
		panic(err)
	}

	time.Sleep(time.Second * 3)

	cityInput, err := executor.WD.FindElement(selenium.ByID, "office-location-input")
	if err != nil {
		panic(err)
	}

	cityInput.SendKeys(city)

	time.Sleep(time.Second * 3)

	search, err := executor.WD.FindElement(selenium.ByCSSSelector, "#office-search-form .lm-Button")
	if err != nil {
		panic(err)
	}

	search.Click()

	time.Sleep(time.Second * 3)

	for {
		_, err := executor.WD.FindElement(selenium.ByCSSSelector, "main > main > div:nth-child(3)")

		if err == nil {
			break
		}

		fmt.Println("Still loading")
		time.Sleep(time.Second)
	}

	htmlOffices, err := executor.WD.FindElements(selenium.ByCSSSelector, "main > main > div:nth-child(2) > li")
	if err != nil {
		panic(err)
	}

	for _, htmlOffice := range htmlOffices {
		officeLink, err := htmlOffice.FindElement(selenium.ByCSSSelector, ".lm-LinkStandalone")
		if err != nil {
			panic(err)
		}

		link, err := officeLink.GetAttribute("href")
		if err != nil {
			panic(err)
		}

		fmt.Println("aaaaaaaaaaaaaa")
		fmt.Println("aaaaaaaaaaaaaa")
		fmt.Println("aaaaaaaaaaaaaa")
		fmt.Println("aaaaaaaaaaaaaa")
		fmt.Println("aaaaaaaaaaaaaa")
		fmt.Println(link)
		fmt.Println("aaaaaaaaaaaaaa")
		fmt.Println("aaaaaaaaaaaaaa")
		fmt.Println("aaaaaaaaaaaaaa")
		fmt.Println("aaaaaaaaaaaaaa")
		fmt.Println("aaaaaaaaaaaaaa")
	}

	executor.StopSelenium()

	return agents
}

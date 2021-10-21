package processors

import (
	"fmt"

	"github.com/dexterorion/insurance-scraper/models"
	"github.com/gocolly/colly"
)

// Progressive represents the Progressive structure
type Progressive struct{}

// Process processes html
func (an Progressive) Process(zip, state, city string) []models.Agent {
	c := colly.NewCollector()

	agents := make([]models.Agent, 0)

	c.OnHTML(".agent.ghost ", func(e *colly.HTMLElement) {
		agent := &models.Agent{}

		agent.Name = e.ChildText(".name")

		phone := e.ChildText(".phone")
		agent.Phone = phone
		agent.Fax = phone

		agent.Address = e.ChildText(".location")

		agents = append(agents, *agent)
	})

	url := fmt.Sprintf("https://www.progressive.com/agent/find-an-agent/?zipcode=%s&product=AU&state=%s&languageSpoken=E", zip, state)
	c.Visit(url)

	return agents
}

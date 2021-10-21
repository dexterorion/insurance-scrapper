package processors

import (
	"fmt"
	"strings"

	"github.com/dexterorion/insurance-scraper/models"
	"github.com/gocolly/colly"
)

// Farmers represents the Farmers structure
type Farmers struct{}

// Process processes html
func (an Farmers) Process(zip, state, city string) []models.Agent {
	c := colly.NewCollector()

	agents := make([]models.Agent, 0)

	c.OnHTML(".location-list .location-list-results .location .location-details", func(e *colly.HTMLElement) {
		agent := &models.Agent{}

		agent.Name = e.ChildText(".location-title-link")

		phone := e.ChildText(".location-info-phone .hidden-xs.location-info-phone-link")
		if strings.Contains(phone, "Office:") {
			phone = strings.ReplaceAll(phone, "Office:", "")
			phone = strings.TrimSpace(phone)
		}
		agent.Phone = phone
		agent.Fax = phone

		street1 := e.ChildText(".location-info-address .c-AddressRow .c-address-street-1")
		street2 := e.ChildText(".location-info-address .c-AddressRow .c-address-street-2")
		city := e.ChildText(".location-info-address .c-AddressRow .c-address-city")
		state := e.ChildText(".location-info-address .c-AddressRow .c-address-state")
		postalcode := e.ChildText(".location-info-address .c-AddressRow .c-address-postal-code")

		address := fmt.Sprintf("%s %s %s %s %s", street1, street2, city, state, postalcode)
		agent.Address = address

		agents = append(agents, *agent)
	})

	url := fmt.Sprintf("https://agents.farmers.com/amp-2?q=%s", zip)
	c.Visit(url)

	return agents
}

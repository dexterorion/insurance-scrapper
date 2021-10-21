package processors

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly"
	"github.com/seeker-insurance/seeker-scraper/models"
)

// AmericanNational represents the American National structure
type AmericanNational struct{}

// Process processes html
func (an AmericanNational) Process(zip, state, city string) []models.Agent {
	c := colly.NewCollector()

	agents := make([]models.Agent, 0)

	c.OnHTML("#agentTable tr td:nth-child(2)", func(e *colly.HTMLElement) {
		agent := &models.Agent{}
		containsSte := false

		e.ForEach("tr", func(index int, e *colly.HTMLElement) {
			text := strings.TrimSpace(e.Text)

			if index == 0 {
				agent.Name = text
			}

			if index == 1 {
				agent.Address = text
			}

			if index == 2 && strings.Contains(strings.ToLower(text), "ste ") {
				agent.Address = fmt.Sprintf("%s %s", agent.Address, text)
				containsSte = true
			}

			if index == 2 && !containsSte || index == 3 && containsSte {
				agent.Address = fmt.Sprintf("%s %s", agent.Address, text)
			}

			if strings.Contains(text, "Phone:") {
				phone := strings.ReplaceAll(text, "Phone:", "")
				phone = strings.TrimSpace(phone)
				agent.Phone = phone
			}

			if strings.Contains(text, "Email:") {
				email := strings.ReplaceAll(text, "Email:", "")
				email = strings.TrimSpace(email)
				email = strings.ToLower(email)
				agent.Email = email
			}

			if strings.Contains(text, "Licenses:") {
				licenses := strings.ReplaceAll(text, "Licenses:", "")
				licenses = strings.TrimSpace(licenses)
				agent.Licenses = licenses
			}

			if strings.Contains(text, "Fax:") {
				fax := strings.ReplaceAll(text, "Fax:", "")
				fax = strings.TrimSpace(fax)
				agent.Fax = fax
			}
		})

		agents = append(agents, *agent)
	})

	url := fmt.Sprintf("https://client.anpac.info/Agent_Locate/agentlocator/AgentList?type=ZipCode&value=%s&value2=100", zip)
	c.Visit(url)

	return agents
}

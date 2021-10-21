package models

// Agent represents the agent data
type Agent struct {
	Name     string
	Address  string
	Phone    string
	Fax      string
	Licenses string
	Email    string
}

// ToStringArray returns an array of strings
func (a Agent) ToStringArray() []string {
	return []string{a.Name, a.Address, a.Phone, a.Fax, a.Email, a.Licenses}
}

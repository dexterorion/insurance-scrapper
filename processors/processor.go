package processors

import "github.com/dexterorion/insurance-scraper/models"

// Processor represents the processor interface
type Processor interface {
	Process(zip, state, city string) []models.Agent
}

// ProcessorType represents the processor type value
type ProcessorType string

const (
	// ProcessorTypeAmericanNational represents the processor type american national
	ProcessorTypeAmericanNational ProcessorType = "americannational"

	// ProcessorTypeFarmers represents the processor type farmers
	ProcessorTypeFarmers ProcessorType = "farmers"

	// ProcessorTypeAllState represents the processor type allstate
	ProcessorTypeAllState ProcessorType = "allstate"

	// ProcessorTypeProgressive represents the processor type progressive
	ProcessorTypeProgressive ProcessorType = "progressive"

	// ProcessorTypeStateFarm represents the processor type state farm
	ProcessorTypeStateFarm ProcessorType = "statefarm"

	// ProcessorTypeAmfam represents the processor type amfam
	ProcessorTypeAmfam ProcessorType = "amfam"

	// ProcessorTypeLiberty represents the processor type liberty
	ProcessorTypeLiberty ProcessorType = "liberty"

	// ProcessorTypeNationwide represents the processor type nationwide
	ProcessorTypeNationwide ProcessorType = "nationwide"

	// ProcessorTypeTravelers represents the processor type travelers
	ProcessorTypeTravelers ProcessorType = "travelers"

	// ProcessorTypeSafeco represents the processor type safeco
	ProcessorTypeSafeco ProcessorType = "safeco"
)

// ProcessorMap represents the mapping between processor type and the structure
var ProcessorMap map[ProcessorType]Processor = map[ProcessorType]Processor{
	ProcessorTypeAmericanNational: AmericanNational{},
	ProcessorTypeFarmers:          Farmers{},
	ProcessorTypeAllState:         AllState{},
	ProcessorTypeProgressive:      Progressive{},
	ProcessorTypeStateFarm:        StateFarm{},
	ProcessorTypeAmfam:            Amfam{},
	ProcessorTypeLiberty:          Liberty{},
	ProcessorTypeNationwide:       Nationwide{},
	ProcessorTypeTravelers:        Travelers{},
	ProcessorTypeSafeco:           Safeco{},
}

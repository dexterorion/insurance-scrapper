package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/dexterorion/insurance-scraper/models"
	"github.com/dexterorion/insurance-scraper/processors"
)

// Create a new type for a list of Strings
type stringList []string

// Implement the flag.Value interface
func (s *stringList) String() string {
	return fmt.Sprintf("%v", *s)
}

func (s *stringList) Set(value string) error {
	*s = strings.Split(value, ",")
	return nil
}

func main() {
	var processor processors.Processor
	var agents []models.Agent
	var zip string
	var city string
	var state string

	// Subcommands
	runCommand := flag.NewFlagSet("run", flag.ExitOnError)

	// Arguments
	runCommandType := runCommand.String("type", "", "Type to use. (Required)")
	runCommandZip := runCommand.String("zip", "string", "Zip {at least one required: zip, state or city}.")
	runCommandState := runCommand.String("state", "string", "State {at least one required: zip, state or city}.")
	runCommandCity := runCommand.String("city", "string", "City {at least one required: zip, state or city}.")

	// Verify that a subcommand has been provided
	// os.Arg[0] is the main command
	// os.Arg[1] will be the subcommand
	if len(os.Args) < 2 {
		fmt.Println("run subcommad is required")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "run":
		runCommand.Parse(os.Args[2:])
	default:
		flag.PrintDefaults()
		os.Exit(1)
	}

	if runCommand.Parsed() {
		if *runCommandType == "" {
			runCommand.PrintDefaults()
			os.Exit(1)
		}

		if *runCommandZip == "" && *runCommandState == "" && *runCommandCity == "" {
			runCommand.PrintDefaults()
			os.Exit(1)
		}

		processorType := processors.ProcessorType(*runCommandType)
		processor = processors.ProcessorMap[processorType]

		if processor == nil {
			log.Fatal("Invalid processor type", fmt.Errorf("Processor type with value [%s] does not exist", *runCommandType))
		}

		zip = *runCommandZip
		city = *runCommandCity
		state = *runCommandState
	}

	agents = processor.Process(zip, state, city)

	var searchData string
	if zip != "" {
		searchData = zip
	} else {
		if city != "" {
			searchData = city
		} else {
			searchData = state
		}
	}

	file, err := os.Create(fmt.Sprintf("%s-%s-%s.csv", *runCommandType, searchData, time.Now().Format("2006-01-02")))
	checkError("Cannot create file", err)
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	err = writer.Write([]string{"Name", "Address", "Phone", "Fax", "Email", "Licenses"})
	checkError("Cannot write to file", err)

	for _, agent := range agents {
		err = writer.Write(agent.ToStringArray())
		checkError("Cannot write to file", err)
	}
}

func checkError(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
}

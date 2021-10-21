package outside

import (
	"fmt"
	"os"

	"github.com/tebeka/selenium/chrome"

	"github.com/tebeka/selenium"
)

// SeleniumExecutor represents the selenium executor structure
type SeleniumExecutor struct {
	WD      selenium.WebDriver
	Service *selenium.Service
}

// StartSelenium starts selenium
func (se *SeleniumExecutor) StartSelenium() {
	// Start a Selenium WebDriver server instance (if one is not already
	// running).
	const (
		// These paths will be different on your system.
		seleniumPath    = "./outside/selenium-server.jar"
		geckoDriverPath = "./outside/geckodriver-v0.27.0-linux32"
		port            = 4433
	)
	opts := []selenium.ServiceOption{
		selenium.StartFrameBuffer(),           // Start an X frame buffer for the browser to run in.
		selenium.GeckoDriver(geckoDriverPath), // Specify the path to GeckoDriver in order to use Firefox.
		selenium.Output(os.Stderr),            // Output debug information to STDERR.
	}
	// selenium.SetDebug(true)
	service, err := selenium.NewSeleniumService(seleniumPath, port, opts...)
	if err != nil {
		fmt.Println(err.Error())
		panic(err) // panic is used only as an example and is not otherwise recommended.
	}

	// Connect to the WebDriver instance running locally.
	caps := selenium.Capabilities{}
	chrome := chrome.Capabilities{
		Path: "/usr/bin/google-chrome",
		Args: []string{
			"--headless", // <<<
			"--no-sandbox",
			"--user-agent=Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_2) AppleWebKit/604.4.7 (KHTML, like Gecko) Version/11.0.2 Safari/604.4.7",
		},
	}
	caps.AddChrome(chrome)

	wd, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", port))
	if err != nil {
		panic(err)
	}

	se.WD = wd
	se.Service = service
}

// StopSelenium stops selenium execution
func (se *SeleniumExecutor) StopSelenium() {
	se.Service.Stop()
	se.WD.Quit()
}

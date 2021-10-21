# Insurance Scraper

This project provides a funcionality to extract information regarding agent from some selected websites.

## Dependencies

- Golang (v1.12+)
- xvfb
- Java (v11+)
- Google Chrome

## Installation

 - First, you should clone the projece
 
 ```$ git clone github.com/dexterorion/insurance-scrapper.git```
 
 - Install golang dependencies
 
 ```$ go get ```
 
## Usage

  ```go run main.go run [ARGS]```
  
  ### Args
  
  - `--type (required)`: \[ americannational | farmers | allstate | progressive | statefarm | amfam | liberty | nationwide | travelers | safeco \]
  - `--zip (required)`: string
  - `--city (optional)`: string
  - `--state (optional)`: string
  
  One of the three optional arguments should be passed, otherwise it will fail.

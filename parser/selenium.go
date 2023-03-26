package main

import (
	"errors"
	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
)

func initSelenium() (service *selenium.Service, driver selenium.WebDriver) {
	service, err := selenium.NewChromeDriverService(config.chromeDriverPath, 4444)
	if err != nil {
		panic(err)
	}

	caps := selenium.Capabilities{}
	args := []string{
		"window-size=1920x1080",
		"--no-sandbox",
		"--disable-dev-shm-usage",
		"disable-gpu",
	}
	if config.headless {
		args = append(args, "--headless")
	}
	caps.AddChrome(chrome.Capabilities{Args: args})
	driver, err = selenium.NewRemote(caps, "")
	err = driver.ResizeWindow("", 1920, 1080)
	return
}
func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func createNewTabAndSetCurrent(driver selenium.WebDriver) (handle string, err error) {
	handles, err := driver.WindowHandles()
	_, err = driver.ExecuteScript("window.open()", nil)
	handlesWithNew, err := driver.WindowHandles()
	for _, handle = range handlesWithNew {
		if !contains(handles, handle) {
			err = driver.SwitchWindow(handle)
			return
		}
	}
	return "", errors.New("unable to create a new tab")
}

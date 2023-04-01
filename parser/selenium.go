package main

import (
	"errors"
	"fmt"
	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
	"log"
	"time"
)

func getRemote() (driver selenium.WebDriver, err error) {
	if config.localChromeDriver {
		_, err := selenium.NewChromeDriverService("chromedriver", 4444)
		if err != nil {
			panic(err)
		}
	}
	sleep := time.Duration(1) * time.Second
	attempts := 5
	for i := 0; i < attempts; i++ {
		if i > 0 {
			log.Println("retrying after error:", err)
			time.Sleep(sleep)
			sleep *= 2
		}
		caps := selenium.Capabilities{}
		caps.AddChrome(chrome.Capabilities{Args: config.chromeArgs})
		driver, err = selenium.NewRemote(caps, config.chromeUrlPrefix)
		if err == nil {
			return
		}
	}
	err = fmt.Errorf("after %d attempts, last error: %s", attempts, err)
	return
}
func initSelenium() (driver selenium.WebDriver) {
	driver, err := getRemote()
	if err != nil {
		panic(err)
	}
	if err = driver.ResizeWindow("", 1920, 1080); err != nil {
		panic(err)
	}
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

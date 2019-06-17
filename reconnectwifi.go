package main

import (
	"errors"
	"net/http"
	"os/exec"
	"time"

	"github.com/martinlindhe/notify"
)

func restartWifi() error {
	errCount := 0
	cmdOff := exec.Command("/usr/bin/nmcli", "radio", "wifi", "off")
	if err := cmdOff.Run(); err != nil {
		errCount++
	}

	cmdOn := exec.Command("/usr/bin/nmcli", "radio", "wifi", "on")
	if err := cmdOn.Run(); err != nil {
		errCount++
	}

	if errCount > 0 {
		return errors.New("Could not restart NetworkManager")
	} else {
		return nil
	}
}

func main() {
	timeout := time.Duration(3 * time.Second)
	client := http.Client{Timeout: timeout}
	for {
		_, err := client.Get("http://google.com")
		if err != nil {
			notify.Notify("Internet Checker", "Connection failure, restarting wifi...", "", "")
			if err := restartWifi(); err != nil {
				notify.Notify("Internet Checker", "Could not restart wifi, retrying in 6 seconds", "", "")
			}
			time.Sleep(6 * time.Second)
		}
		time.Sleep(time.Second)
	}
}

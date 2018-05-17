package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type amIMole struct {
	IP                    string      `json:"ip"`
	Country               string      `json:"country"`
	City                  interface{} `json:"city"`
	Longitude             float64     `json:"longitude"`
	Latitude              float64     `json:"latitude"`
	MullvadExitIP         bool        `json:"mullvad_exit_ip"`
	MullvadExitIPHostname string      `json:"mullvad_exit_ip_hostname"`
	MullvadServerType     string      `json:"mullvad_server_type"`
	Blacklisted           struct {
		Blacklisted bool `json:"blacklisted"`
		Results     []struct {
			Name        string `json:"name"`
			Link        string `json:"link"`
			Blacklisted bool   `json:"blacklisted"`
		} `json:"results"`
	} `json:"blacklisted"`
	Organization string `json:"organization"`
}

func getJSON() amIMole {
	resp, err := http.Get("https://am.i.mullvad.net/json")
	if err != nil {
		panic(err)
	}

	body, err := ioutil.ReadAll(resp.Body)

	amiMole := amIMole{}
	json.Unmarshal(body, &amiMole)
	// fmt.Println(amiMole)
	return amiMole
}

func main() {
	fmt.Println("Hi! Getting data from am.i.mullvad.net!")
	json := getJSON()
	if json.MullvadExitIP == true {
		fmt.Println("You ARE connected to Mullvad!")
		fmt.Println("IP: ", json.IP)
		fmt.Println("Exit IP Hostname: ", json.MullvadExitIPHostname)
		fmt.Println("Server Type: ", json.MullvadServerType)
	} else {
		fmt.Println("You are NOT connected to Mullvad.")
		fmt.Println("IP: ", json.IP)
	}
	fmt.Println("Country:", json.Country)
	fmt.Println("City", json.City)

	if json.Blacklisted.Blacklisted == true {
		fmt.Println("Your IP IS blacklisted!")
		for _, blacklist := range json.Blacklisted.Results {
			fmt.Println(blacklist.Name, ": ", blacklist.Blacklisted)
			fmt.Println(blacklist.Link)
		}
	} else {
		fmt.Println("Your IP is NOT blacklisted!")
	}

	fmt.Println("Organisation: ", json.Organization)
}

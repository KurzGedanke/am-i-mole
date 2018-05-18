package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
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
	if err != nil {
		panic(err)
	}

	amiMole := amIMole{}
	json.Unmarshal(body, &amiMole)
	return amiMole
}

func printFull(json amIMole) {
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

func printIP(json amIMole) {
	fmt.Println(json.IP)
}

func printConnected(json amIMole) {
	fmt.Println(json.MullvadExitIP)
}

func printCountry(json amIMole) {
	fmt.Println(json.Country)
}

func printCity(json amIMole) {
	fmt.Println(json.City)
}

func printBlacklisted(json amIMole) {
	fmt.Println(json.Blacklisted.Blacklisted)
	if json.Blacklisted.Blacklisted == true {
		for _, blacklist := range json.Blacklisted.Results {
			fmt.Println(blacklist.Name, ": ", blacklist.Blacklisted)
			fmt.Println(blacklist.Link)
		}
	}
}

func printOrganization(json amIMole) {
	fmt.Println(json.Organization)
}

func main() {
	ip := flag.Bool("ip", false, "Prints your current IP.")
	conn := flag.Bool("c", false, "Prints if connected to mullvad.")
	country := flag.Bool("ct", false, "Prints the country connected in.")
	city := flag.Bool("cty", false, "Prints the city connected in.")
	blacklist := flag.Bool("black", false, "Prints if blacklisted.")
	organization := flag.Bool("o", false, "Prints the organization connecte to.")
	flag.Parse()

	json := getJSON()

	if *ip == true {
		printIP(json)
	}
	if *conn == true {
		printConnected(json)
	}
	if *country == true {
		printCountry(json)
	}
	if *city == true {
		printCity(json)
	}
	if *blacklist == true {
		printBlacklisted(json)
	}
	if *organization == true {
		printOrganization(json)
	}

	if len(os.Args) == 1 {
		printFull(json)
	}
}

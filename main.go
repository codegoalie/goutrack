package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os/user"

	"gopkg.in/alecthomas/kingpin.v2"
	"gopkg.in/yaml.v2"
)

var (
	host          = kingpin.Flag("host", "URL for YouTrack installation").Short('h').String()
	login         = kingpin.Flag("login", "login for YouTrack").Short('l').String()
	password      = kingpin.Flag("password", "password for YouTrack").Short('p').String()
	command       = kingpin.Arg("command", "action to perform on YouTrack").Required().String()
	story         = kingpin.Arg("story", "YouTrack story ID").Required().String()
	commandString = kingpin.Arg("commandString", "command string to apply to the story").String()
	commentString = kingpin.Arg("commentString", "add a comment to the story").String()
)

type Config struct {
	Host     string
	Username string
	Password string
}

func main() {
	kingpin.Parse()

	config := readConfigFromFile()
	if *login != "" {
		config.Username = *login
	}
	if *password != "" {
		config.Username = *password
	}
	if *host != "" {
		config.Host = *host
	}

	client := NewYouTrackClient(config.Host, config.Username, config.Password)

	switch *command {
	case "g":
		fmt.Println("Fetching", *story)
		fmt.Println(client.GetIssue(*story))
	case "c":
		fmt.Println("Applying", *commandString, "to", *story)
		fmt.Println(client.CommandIssue(*story, *commandString, *commentString))
	}
}

func readConfigFromFile() Config {
	usr, err := user.Current()
	if err != nil {
		log.Println(err)
		return Config{}
	}
	data, err := ioutil.ReadFile(usr.HomeDir + "/.goutrack")
	if err != nil {
		log.Println(err)
		return Config{}
	}

	config := Config{}
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		log.Printf("error: %v", err)
		return Config{}
	}
	return config
}

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
	login         = kingpin.Flag("login", "login for YouTrack").Short('l').String()
	password      = kingpin.Flag("password", "password for YouTrack").Short('p').String()
	command       = kingpin.Arg("command", "action to perform on YouTrack").Required().String()
	story         = kingpin.Arg("story", "YouTrack story ID").Required().String()
	commandString = kingpin.Arg("commandString", "command string to apply to story").String()
)

type Config struct {
	Username string
	Password string
}

func main() {
	kingpin.Parse()

	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	data, err := ioutil.ReadFile(usr.HomeDir + "/.goutrack")
	if err != nil {
		log.Fatal(err)
	}

	config := Config{}
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	if *login != "" {
		config.Username = *login
	}
	if *password != "" {
		config.Username = *password
	}

	client := NewYouTrackClient(config.Username, config.Password)

	switch *command {
	case "g":
		fmt.Println("Fetching", *story)
		fmt.Println(client.GetIssue(*story))
	case "c":
		fmt.Println("Applying", *commandString, "to", *story)
		fmt.Println(client.CommandIssue(*story, *commandString))
	}
}

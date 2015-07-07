package main

import (
	"fmt"

	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	login         = kingpin.Flag("login", "login for YouTrack").Short('l').String()
	password      = kingpin.Flag("password", "password for YouTrack").Short('p').String()
	command       = kingpin.Arg("command", "action to perform on YouTrack").Required().String()
	story         = kingpin.Arg("story", "YouTrack story ID").Required().String()
	commandString = kingpin.Arg("commandString", "command string to apply to story").String()
)

func main() {
	kingpin.Parse()
	client := NewYouTrackClient(*login, *password)

	switch *command {
	case "g":
		fmt.Println("Fetching", *story)
		fmt.Println(client.GetIssue(*story))
	case "c":
		fmt.Println("Applying", *commandString, "to", *story)
		fmt.Println(client.CommandIssue(*story, *commandString))
	}
}

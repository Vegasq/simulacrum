package utils

import (
	"flag"
	"fmt"
	"os"
)

// Vars represent parsed command line
type Vars struct {
	OldName string
	NewName string
}

// DefineCommandline parse commandline
func DefineCommandline() Vars {
	var projectName = flag.String("old", "",
		"Old project name")
	var newProjectName = flag.String("new", "",
		"New project name")
	flag.Parse()
	if *projectName == "" {
		fmt.Println("Specify original project name in -old")
		os.Exit(1)
	}
	if *newProjectName == "" {
		fmt.Println("Specify new project name in -new")
		os.Exit(1)
	}
	return Vars{OldName: *projectName, NewName: *newProjectName}
}

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
	var projectName = flag.String("project_name", "",
		"help message for project_name")
	var newProjectName = flag.String("new_project_name", "",
		"help message for new_project_name")
	flag.Parse()
	if *projectName == "" {
		fmt.Println("project_name can't be empty")
		os.Exit(1)
	}
	if *newProjectName == "" {
		fmt.Println("new_project_name can't be empty")
		os.Exit(1)
	}
	return Vars{OldName: *projectName, NewName: *newProjectName}
}

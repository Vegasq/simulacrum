package main

import (
	"fmt"
	"project"
	"utils"
)

func main() {
	cmdlnData := utils.DefineCommandline()
	proj := project.GetProject(cmdlnData.OldName, cmdlnData.NewName)
	proj.Clone()
	fmt.Println("Done")
}

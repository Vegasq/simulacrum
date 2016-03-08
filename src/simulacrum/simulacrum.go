package main

import (
	"fmt"
	"project"
	"utils"
)

func main() {
	cmdlnData := utils.DefineCommandline()
	fmt.Println("get_proj")
	proj := project.GetProject(cmdlnData.OldName, cmdlnData.NewName)
	fmt.Println(proj.GetConfig())
	fmt.Println("clone")
	proj.Clone()
	fmt.Println("Done")
}

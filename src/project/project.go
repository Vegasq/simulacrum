package project

import (
	"domains"
	"fmt"
	"networks"
	"utils"
)

// Project represent prefix for networks and domains
type Project struct {
	config utils.GenericConfig
}

// GetConfig returns config associated with project
func (proj Project) GetConfig() utils.GenericConfig {
	return proj.config
}

// Clone networks and domains in project
func (proj Project) Clone() {
	proj.cloneNetworks()
	// proj.cloneDomains()
}

func (proj Project) cloneNetworks() {
	fmt.Println("cloneNetworks")
	fmt.Println(proj.GetConfig())
	fmt.Println(proj.GetConfig().GetDomainXMLPath())
	var nets = networks.GetNetworks(proj.GetConfig())
	fmt.Println("netsclone")
	nets.Clone()
}

func (proj Project) cloneDomains() {
	var doms = domains.GetDomains(
		proj.config.GetOldName(), proj.config.GetNewName(),
		proj.config.GetDomainXMLPath())
	doms.Clone()
}

// GetProject returns Project
func GetProject(oldName string, newName string) Project {
	conf := GetConfig(oldName, newName)
	fmt.Println("Names are updated:")
	fmt.Println(conf.GetNewName() == newName)
	fmt.Println(conf.GetOldName() == oldName)
	proj := Project{config: conf}
	return proj
}

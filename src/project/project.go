package project

import (
	"domains"
	"networks"
)

// Project represent prefix for networks and domains
type Project struct {
	oldName string
	newName string
}

// Clone networks and domains in project
func (proj Project) Clone() {
	proj.cloneNetworks()
	proj.cloneDomains()
}

func (proj Project) cloneNetworks() {
	var nets = networks.GetNetworks(proj.oldName, proj.newName)
	nets.Clone()
}

func (proj Project) cloneDomains() {
	var doms = domains.GetDomains(proj.oldName, proj.newName)
	doms.Clone()
}

// GetProject returns Project
func GetProject(oldName string, newName string) Project {
	proj := Project{oldName: oldName, newName: newName}
	return proj
}

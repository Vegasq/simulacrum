package domains

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// Domains represent slice of exists Network
type Domains struct {
	domains       []Domain
	oldName       string
	newName       string
	domainDirPath string
}

// Init networks
func (doms *Domains) Init() {
	doms.loadDomains()
}

// Clone will call Clone method of domains
func (doms *Domains) Clone() {
	fmt.Println("Domains to clone:")
	fmt.Println(doms.domains)
	for i := 0; i < len(doms.domains); i++ {
		doms.domains[i].Clone()
	}
}

// add will attach extra network to networks
func (doms *Domains) add(dom Domain) {
	doms.domains = append(doms.domains, dom)
}

func (doms *Domains) loadDomains() {
	var files, err = ioutil.ReadDir(doms.domainDirPath)
	if err != nil {
		os.Exit(1)
	}
	fmt.Println("Doms before loading:")
	fmt.Println(doms.domains)
	for i := 0; i < len(files); i++ {
		var xmlFilePath = filepath.Join(doms.domainDirPath, files[i].Name())
		fmt.Println(xmlFilePath)
		if strings.HasPrefix(files[i].Name()+"_", doms.oldName) {
			doms.add(GetDomain(xmlFilePath, doms.oldName, doms.newName))
		}
	}
	fmt.Println("Doms after loading:")
	fmt.Println(doms.domains)
}

// GetDomains build Networks instance
func GetDomains(oldName string, newName string, xmlPath string) Domains {
	doms := Domains{
		domainDirPath: xmlPath, oldName: oldName, newName: newName}
	doms.Init()
	return doms
}

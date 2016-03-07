package networks

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

// Networks represent slice of exists Network
type Networks struct {
	networks       []Network
	oldName        string
	newName        string
	networkDirPath string
}

// Init networks
func (nets *Networks) Init() {
	nets.loadNetworks()
}

// Clone will call Clone method of networks
func (nets *Networks) Clone() {
	fmt.Println("Nets to clone:")
	fmt.Println(nets.networks)
	for i := 0; i < len(nets.networks); i++ {
		nets.networks[i].Clone()
	}
}

// add will attach extra network to networks
func (nets *Networks) add(net Network) {
	nets.networks = append(nets.networks, net)
}

func (nets *Networks) loadNetworks() {
	var files, err = ioutil.ReadDir(nets.networkDirPath)
	if err != nil {
		os.Exit(1)
	}
	fmt.Println("Nets before loading:")
	fmt.Println(nets.networks)
	for i := 0; i < len(files); i++ {
		var xmlFilePath = filepath.Join(nets.networkDirPath, files[i].Name())
		fmt.Println(xmlFilePath)
		nets.add(GetNetwork(xmlFilePath, nets.oldName, nets.newName))
	}
	fmt.Println("Nets after loading:")
	fmt.Println(nets.networks)
}

// GetNetworks build Networks instance
func GetNetworks(oldName string, newName string) Networks {
	nets := Networks{
		networkDirPath: "/etc/libvirt/qemu/networks/",
		oldName:        oldName,
		newName:        newName}
	nets.Init()
	return nets
}

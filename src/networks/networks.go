package networks

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
	"utils"
)

// Networks represent slice of exists Network
type Networks struct {
	networks []Network
	config   utils.GenericConfig
}

// Init networks
func (nets *Networks) Init() {
	fmt.Println("In Init")
	fmt.Println(nets.config.GetDomainXMLPath())
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
	fmt.Println("loadNetworks")
	networkXMLPath, err := filepath.Abs(nets.config.GetNetworkXMLPath())
	if err != nil {
		fmt.Println("Can't get network xml dir")
	}
	var files, _ = ioutil.ReadDir(networkXMLPath)
	fmt.Println("All XMLs from " + networkXMLPath)
	fmt.Println(nets.networks)
	for i := 0; i < len(files); i++ {
		var xmlFilePath = filepath.Join(
			nets.config.GetNetworkXMLPath(), files[i].Name())

		fmt.Println(xmlFilePath)
		if strings.HasPrefix(files[i].Name()+"_", nets.config.GetOldName()) {
			nets.add(GetNetwork(
				xmlFilePath,
				nets.config.GetOldName(),
				nets.config.GetNewName()))
		}
	}
	fmt.Println("Nets after loading:")
	fmt.Println(nets.networks)
}

// GetNetworks build Networks instance
func GetNetworks(config utils.GenericConfig) Networks {
	fmt.Println("Get networks")
	fmt.Println(config)
	fmt.Println(config.GetDomainXMLPath())
	nets := Networks{config: config}
	fmt.Println("init nets")
	nets.Init()
	fmt.Println("return nets")
	return nets
}

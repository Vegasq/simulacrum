package networks

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"utils"
)

// GetNetwork generate Network instance
func GetNetwork(xmlPath string, oldName string, newName string) Network {
	return Network{xmlPath: xmlPath, oldName: oldName, newName: newName}
}

// Network represent single virsh network
type Network struct {
	xmlPath string
	oldName string
	newName string
	netXML  XMLNetwork
}

// Bridge defines bridge interface for network
type Bridge struct {
	Name string `xml:"name,attr"`
}

// Mac defines MAC address
type Mac struct {
	Address string `xml:"address,attr"`
}

// IP defines CIDR
type IP struct {
	Address string `xml:"address,attr"`
}

// XMLNetwork defines xml structure
type XMLNetwork struct {
	XMLName xml.Name `xml:"network"`
	Name    string   `xml:"name"`
	UUID    string   `xml:"uuid"`
	Bridge  Bridge   `xml:"bridge"`
	MAC     Mac      `xml:"mac"`
	IP      IP       `xml:"ip"`
}

// FixStruct replace information in original XML with new one
func (net *Network) FixStruct() {
	net.netXML.UUID = utils.GenUUID()
	net.netXML.Name = strings.Replace(
		net.netXML.Name,
		net.oldName,
		net.newName,
		1)
	net.netXML.Bridge.Name = strings.Replace(
		net.netXML.Bridge.Name,
		"fuel",
		"smlkr",
		1)
	net.netXML.MAC.Address = utils.GenMAC()
	net.netXML.IP.Address = strings.Replace(
		net.netXML.IP.Address,
		".128.",
		".129.",
		1)
}

// Clone network
func (net *Network) Clone() {
	net.netXML = XMLNetwork{}

	// Get full path
	networkXMLPath, err := filepath.Abs(net.xmlPath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Open file
	file, err := os.Open(networkXMLPath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	// Read file
	data, err := ioutil.ReadAll(file)

	// Parse XML
	err = xml.Unmarshal(data, &net.netXML)

	// Replace data
	net.FixStruct()

	newXMLPath := strings.Replace(networkXMLPath, net.oldName, net.newName, 1)
	file, err = os.Create(newXMLPath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	enc := xml.NewEncoder(file)
	enc.Indent("  ", "    ")
	if err := enc.Encode(net.netXML); err != nil {
		fmt.Printf("error: %v\n", err)
	}

	fmt.Println("Network clonned:")
	fmt.Println(net)

}

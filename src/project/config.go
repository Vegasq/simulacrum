package project

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// Config contains configuration
type Config struct {
	DomainXMLPath  string `json:"domainXMLPath"`
	NetworkXMLPath string `json:"networkXMLPath"`
	OldName        string `json:"oldName"`
	NewName        string `json:"newName"`
}

// SetOldName SetOldName
func (conf *Config) SetOldName(name string) {
	fmt.Println("SetOldName: " + name)
	conf.OldName = name
}

// SetNewName SetNewName
func (conf *Config) SetNewName(name string) {
	fmt.Println("SetNewName: " + name)
	conf.NewName = name
}

// GetOldName returns path from config
func (conf Config) GetOldName() string {
	return conf.OldName
}

// GetNewName returns path from config
func (conf Config) GetNewName() string {
	return conf.NewName
}

// GetDomainXMLPath returns path from config
func (conf Config) GetDomainXMLPath() string {
	return conf.DomainXMLPath
}

// GetNetworkXMLPath returns path from config
func (conf Config) GetNetworkXMLPath() string {
	fmt.Println(conf)
	return conf.NetworkXMLPath
}

// GetConfig returns Config
func GetConfig(oldName string, newName string) *Config {
	configName := "conf.json"
	if _, err := os.Stat(configName); os.IsNotExist(err) {
		fmt.Println("conf.json not found")
		ioutil.WriteFile(
			configName,
			[]byte("{\"domainXMLPath\": \"/etc/libvirt/qemu/\",\n"+
				"\"networkXMLPath\": \"/etc/libvirt/qemu/networks/\"}\n"),
			0644)
	}
	file, confErr := os.Open(configName)
	if confErr != nil {
		fmt.Println("Can't open config file")
	}
	decoder := json.NewDecoder(file)

	config := &Config{}
	fmt.Println("Fresh config:")
	fmt.Println(config)
	err := decoder.Decode(&config)
	if err != nil {
		fmt.Println("Error during config file parsing:", err)
	}
	if config.GetOldName() == "" {
		config.SetOldName(oldName)
	}
	if config.GetNewName() == "" {
		config.SetNewName(newName)

		fmt.Println(newName == config.GetNewName())
	}

	return config
}

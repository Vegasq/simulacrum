package domains

// GetDomain generate Network instance
func GetDomain(xmlPath string, oldName string, newName string) Domain {
	return Domain{xmlPath: xmlPath, oldName: oldName, newName: newName}
}

// Domain represent single virsh network
type Domain struct {
	xmlPath string
	oldName string
	newName string
}

// Clone network
func (dom *Domain) Clone() {
	dom.cloneXML()
	dom.cloneStorage()
}

func (dom Domain) cloneStorage() {

}

func (dom Domain) cloneXML() {

}

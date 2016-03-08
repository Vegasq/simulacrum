package utils

// GenericConfig defines Config fields
type GenericConfig interface {
	GetDomainXMLPath() string
	GetNetworkXMLPath() string
	GetOldName() string
	GetNewName() string
	SetOldName(string)
	SetNewName(string)
}

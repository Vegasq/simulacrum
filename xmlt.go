package main

import (
	"encoding/xml"
	"fmt"
	"strings"
)

const XML = `  <network>
      <name>mine_management</name>
      <uuid>c5098218-1a80-1aff-999a-714f5f4a72fd</uuid>
      <bridge name="smlkrbr28"></bridge>
      <mac address="52:54:00:f2:a7:83"></mac>
      <ip address="10.129.17.1"></ip>
  </network>`

type Bridge struct {
	Name string `xml:"name,attr"`
}

func main2() {
	xmlToMap()
}

func xmlToMap() {
	decoder := xml.NewDecoder(strings.NewReader(XML))
	for {
		token, _ := decoder.Token()
		if token == nil {
			break
		}
		switch startElement := token.(type) {
		case xml.StartElement:
			if startElement.Name.Local == "bridge" {
				var br Bridge
				decoder.DecodeElement(&br, &startElement)
				fmt.Println(br)
			}
		}
	}
}

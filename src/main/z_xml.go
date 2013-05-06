// z_xml.go
package z_xml

import (
	"encoding/xml"
	"flag"
	"os"
	"log"
)

const  OpType (
	Extensions = iota	
)

var file = flag.String("xmlFile", "config.xml", "xml file path")
var returnString string

type config struct {
	Type string xml:"extensions>type"
}

func init() {
	flag.Parse()
}

func XmlOps(t int) {
	switch t {
	case Extensions:
		xmlFile, err := os.Open(file)
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
		decoder := xml.NewDecoder(xmlFile)
		for {
			t, _ := decoder.Token()
			if t == nil {
				break
			}
			default: fmt.Println("Not a valid option %s", t)
		}
	}
}
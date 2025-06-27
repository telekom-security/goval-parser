package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/telekom-security/goval-parser/oval"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("usage: %s file1 file2 ...\n", os.Args[0])
		return
	}
	for _, f := range os.Args[1:len(os.Args)] {
		fmt.Println(f + ":")
		oval, err := readOval(f)
		if err != nil {
			log.Fatal(err)
		}
		out := json.NewEncoder(os.Stdout)
		out.SetIndent("", "  ")
		if err := out.Encode(oval); err != nil {
			log.Fatal(err)
		}
		if err := out.Encode(oval.Definitions.Definitions[0].Debian); err != nil {
			log.Fatal(err)
		}
	}
}

// readOval : Read OVAL definitions from file
func readOval(file string) (*oval.Root, error) {
	str, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, fmt.Errorf("Can't open file: %s", err)
	}
	oval := &oval.Root{}
	err = xml.Unmarshal([]byte(str), oval)
	if err != nil {
		return nil, fmt.Errorf("Can't parse XML: %s", err)
	}
	return oval, nil
}

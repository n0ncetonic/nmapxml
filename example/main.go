package main

import (
	"encoding/json"
	"fmt"

	"github.com/n0ncetonic/nmapxml"
)

func main() {
	scanData, err := nmapxml.Readfile("sample.xml")
	if err != nil {
		fmt.Println(err)
		return
	}
	jsonData, err := json.Marshal(scanData)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%+v", string(jsonData))
}

package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
)

const head = `
<?xml version="1.0" encoding="utf-8"?>
<!DOCTYPE TS>
`

type location struct {
	Filename string `xml:"filename,attr"`
	Line     string `xml:"line,attr"`
}
type trans struct {
	Trans string `xml:",chardata"`
	Type  string `xml:"type,attr"`
}
type message struct {
	Locations []location `xml:"location"`
	Source    string     `xml:"source"`
	Trans     trans      `xml:"translation"`
}
type context struct {
	Names    string    `xml:"name"`
	Messages []message `xml:"message"`
}
type ts struct {
	XMLName  xml.Name  `xml:"TS"`
	Version  string    `xml:"version,attr"`
	Language string    `xml:"language,attr"`
	Contexts []context `xml:"context"`
}

//Trans trans
func Trans(inFile, outFile string) {
	// filePath := "trans_zh.tr"
	data, err := ioutil.ReadFile(inFile)
	if err != nil {
		fmt.Println(err)
		return
	}
	q := &ts{}
	err = xml.Unmarshal(data, &q)
	if err != nil {
		fmt.Println(err)
		return
	}
	//TODO  add trans flow

	res, err := xml.MarshalIndent(&q, "", "    ")
	if err != nil {
		fmt.Println(err)
		return
	}
	err = ioutil.WriteFile(outFile, res, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
}

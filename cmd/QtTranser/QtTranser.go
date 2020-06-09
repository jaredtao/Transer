package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"

	"github.com/jaredtao/Transer/services/baidu"
	"github.com/jaredtao/Transer/services/youdao"

	"github.com/jaredtao/Transer/services/transer"
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

//QtTransArgs args
type QtTransArgs struct {
	InputFile  string
	OutputFile string
	API        string
	ID         string
	Secret     string
	TargetLan  string
}

//Trans trans
func Trans(args QtTransArgs) {
	// filePath := "trans_zh.tr"
	data, err := ioutil.ReadFile(args.InputFile)
	if err != nil {
		fmt.Println(err)
		return
	}
	transData := &ts{}
	err = xml.Unmarshal(data, &transData)
	if err != nil {
		fmt.Println(err)
		return
	}
	//TODO  add trans flow
	input := &transer.TransInput{
		ID:     args.ID,
		Secret: args.Secret,
		To:     args.TargetLan,
	}
	alreadyTranNum := 0
	transData.Language = args.TargetLan
	for i := 0; i < len(transData.Contexts); i++ {
		ctx := &transData.Contexts[i]
		for j := 0; j < len(ctx.Messages); j++ {
			msg := &ctx.Messages[j]
			if msg.Trans.Trans != "" {
				input.Query = msg.Trans.Trans
			} else {
				input.Query = msg.Source
			}
			if args.API == "baidu" {
				ans := baidu.Trans(input)
				if ans.Result != "" {
					msg.Trans.Trans = ans.Result
					msg.Trans.Type = ""
				}
			} else if args.API == "youdao" {
				ans := youdao.Trans(input)
				if ans.Result != "" {
					msg.Trans.Trans = ans.Result
					msg.Trans.Type = ""
				}
			}
			alreadyTranNum++
			fmt.Printf("translate %d words already.\r\n", alreadyTranNum)
		}
	}
	fmt.Printf("baidu translate %d words failed.\r\n", baidu.GetFailedCnt())

	res, err := xml.MarshalIndent(&transData, "", "    ")
	if err != nil {
		fmt.Println(err)
		return
	}
	err = ioutil.WriteFile(args.OutputFile, res, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
}

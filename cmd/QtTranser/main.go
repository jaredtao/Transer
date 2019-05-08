package main

import (
	"flag"

	"github.com/jaredtao/Transer/services/baidu"
)

const baiduID = "20190502000293463"
const baiduSecret = "0d2RvCho9XZNEO5GCGNs"

const youdaoID = "1bd659586c52ea1d"
const youdaoSecret = "5ZktXhHfLCpI0KnAdcxx4cPyGJwcVXaV"

var api = flag.String("api", "baidu", "baidu | youdao")
var userID = flag.String("userID", baiduID, "your id")
var secret = flag.String("secret", baiduSecret, "your secret")
var inputFile = flag.String("inputFile", "", "the input file need translate")
var outputFile = flag.String("outputFile", "out.tr", "the output file")
var targetLan = flag.String("targetLang", "zh", "zh | en | ja | ko | fr | es | pt | it | ru | vi | de | ar | id")

func main() {
	flag.Parse()
	args := &QtTransArgs{}
	args.API = *api
	args.ID = *userID
	args.Secret = *secret
	args.InputFile = *inputFile
	args.OutputFile = *outputFile
	args.TargetLan = *targetLan
	if *inputFile == "" || *outputFile == "" {
		flag.PrintDefaults()
		return
	}
	if *api == "baidu" {
		ok, to := baidu.LanConvertFromYouDao(*targetLan)
		if ok {
			args.TargetLan = to
		}
		Trans(*args)
	} else if *api == "youdao" {
		Trans(*args)
	} else {
		flag.PrintDefaults()
	}
}

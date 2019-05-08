package main

import (
	"flag"
	"fmt"

	"github.com/jaredtao/Transer/services/baidu"
	"github.com/jaredtao/Transer/services/transer"
	"github.com/jaredtao/Transer/services/youdao"
)

const baiduID = "20190502000293463"
const baiduSecret = "0d2RvCho9XZNEO5GCGNs"

const youdaoID = "1bd659586c52ea1d"
const youdaoSecret = "5ZktXhHfLCpI0KnAdcxx4cPyGJwcVXaV"

var api = flag.String("api", "baidu", "baidu | youdao")
var userID = flag.String("userID", "20190502000293463", "your id")
var secret = flag.String("secret", "0d2RvCho9XZNEO5GCGNs", "your secret")
var originText = flag.String("text", "Hello World", "the text need translate")
var targetLan = flag.String("targetLang", "zh", "zh | en | ja | ko | fr | es | pt | it | ru | vi | de | ar | id")

func main() {
	flag.Parse()
	input := &transer.TransInput{
		ID:     *userID,
		Secret: *secret,
		Query:  *originText,
		To:     *targetLan,
	}
	var result *transer.TransOutput
	if *api == "baidu" {
		ok, to := baidu.LanConvertFromYouDao(*targetLan)
		if ok {
			input.To = to
		}
		result = baidu.Trans(input)
	} else if *api == "youdao" {
		result = youdao.Trans(input)
	} else {
		flag.PrintDefaults()
	}
	fmt.Println(result.Result)
}

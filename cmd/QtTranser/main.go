package main

const baiduID = "20190502000293463"
const baiduSecret = "0d2RvCho9XZNEO5GCGNs"

const youdaoID = "1bd659586c52ea1d"
const youdaoSecret = "5ZktXhHfLCpI0KnAdcxx4cPyGJwcVXaV"

func main() {
	Trans("trans_zh.tr")
	// input := &transer.TransInput{
	// 	ID:     baiduID,
	// 	Secret: baiduSecret,
	// }
	// input.Query = "黄河远上白云间，一片孤城万仞山"
	// input.To = baidu.En
	// res := baidu.Trans(input)
	// fmt.Println(res.Result)

	// fmt.Println("---------------")

	// input.ID = youdaoID
	// input.Secret = youdaoSecret
	// input.To = youdao.En
	// res = youdao.Trans(input)
	// fmt.Println(res.Result)

}

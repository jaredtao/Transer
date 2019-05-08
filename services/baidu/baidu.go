package baidu

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/jaredtao/Transer/services/transer"
)

const baiduAPI = "https://api.fanyi.baidu.com/api/trans/vip/translate"
const (
	//Auto 自动识别
	Auto = "auto"
	//Zh 中文
	Zh = "zh"
	//En 英语
	En = "en"
	//Yue 粤语
	Yue = "yue"
	//Wyw 文言文
	Wyw = "wyw"
	//Jp 日语
	Jp = "jp"
	//Kor 韩语
	Kor = "kor"
	//Fra 法语
	Fra = "fra"
	//Spa 西班牙语
	Spa = "spa"
	//Th 泰语
	Th = "th"
	//Ara 阿拉伯语
	Ara = "ara"
	//Ru 俄语
	Ru = "ru"
	//Pt 葡萄牙语
	Pt = "pt"
	//De 德语
	De = "de"
	//It 意大利语
	It = "it"
	//El 希腊语
	El = "el"
	//Nl 荷兰语
	Nl = "nl"
	//Pl 波兰语
	Pl = "pl"
	//Bul 保加利亚语
	Bul = "bul"
	//Est 爱沙尼亚语
	Est = "est"
	//Dan 丹麦语
	Dan = "dan"
	//Fin 芬兰语
	Fin = "fin"
	//Cs 捷克语
	Cs = "cs"
	//Rom 罗马尼亚语
	Rom = "rom"
	//Slo 斯洛文尼亚语
	Slo = "slo"
	//Swe 瑞典语
	Swe = "swe"
	//Hu 匈牙利语
	Hu = "hu"
	//Cht 繁体中文
	Cht = "cht"
	//Vie 越南语
	Vie = "vie"
)

//transItem  one Trans item
type transItem struct {
	Src string `json:"src"`
	Dst string `json:"dst"`
}

//result trans result
type result struct {
	From        string      `json:"from"`
	To          string      `json:"to"`
	TransResult []transItem `json:"trans_result"`
}

//Trans trans
func Trans(input *transer.TransInput) *transer.TransOutput {
	output := new(transer.TransOutput)
	second := strconv.FormatInt(time.Now().Unix(), 10)
	salt := second
	code := input.ID + input.Query + salt + input.Secret
	h := md5.New()
	h.Write([]byte(code))
	sign := h.Sum(nil)
	signStr := hex.EncodeToString(sign[:])
	values := make(url.Values)
	values["q"] = []string{input.Query}
	values["from"] = []string{Auto}
	values["to"] = []string{input.To}
	values["appid"] = []string{input.ID}
	values["salt"] = []string{salt}
	values["sign"] = []string{signStr}
	res, err := http.PostForm(baiduAPI, values)
	if err != nil {
		fmt.Println(err)
		return output
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return output
	}
	// fmt.Println(string(body))
	r := &result{}
	err = json.Unmarshal(body, &r)
	if err != nil {
		fmt.Println(err)
		return output
	}
	if len(r.TransResult) > 0 {
		// fmt.Println(r.TransResult[0].Src)
		// fmt.Println(r.TransResult[0].Dst)
		output.Result = r.TransResult[0].Dst
	}
	return output
}

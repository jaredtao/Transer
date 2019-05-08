package youdao

import (
	"crypto/sha256"
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

const youdaoAPI = "https://openapi.youdao.com/api"
const (
	//Auto 自动识别
	Auto = "auto"
	//Zh 中文
	Zh = "zh-CHS"
	//En 英文
	En = "en"
	//Ja 日文
	Ja = "ja"
	//Ko 韩文
	Ko = "ko"
	//Fr 法文
	Fr = "fr"
	//Es 西班牙文
	Es = "es"
	//Pt 葡萄牙文
	Pt = "pt"
	//It 意大利文
	It = "it"
	//Ru 俄文
	Ru = "ru"
	//Vi 越南文
	Vi = "vi"
	//De 德文
	De = "de"
	//Ar 阿拉伯文
	Ar = "ar"
	//Id 印尼文
	Id = "id"
)

//dict  one dict
type dict struct {
	URL string `json:"url"`
}

//result trans result
type result struct {
	TSpeakURL   string   `json:"tSpeakUrl"`
	Query       string   `json:"query"`
	Translation []string `json:"translation"`
	ErrorCode   string   `json:"errorCode"`
	Dict        dict     `json:"dict"`
	WebDict     dict     `json:"webdict"`
	Lang        string   `json:"l"`
	SpeakURL    string   `json:"speakUrl"`
}

//Trans trans
func Trans(input *transer.TransInput) *transer.TransOutput {
	output := new(transer.TransOutput)
	second := strconv.FormatInt(time.Now().Unix(), 10)
	slat := second
	code := input.ID + input.Query + slat + second + input.Secret
	h := sha256.New()
	h.Write([]byte(code))
	sign := h.Sum(nil)
	signStr := hex.EncodeToString(sign[:])

	values := make(url.Values)
	values["q"] = []string{input.Query}
	values["from"] = []string{Auto}
	values["to"] = []string{input.To}
	values["signType"] = []string{"v3"}
	values["appKey"] = []string{input.ID}
	values["salt"] = []string{slat}
	values["sign"] = []string{signStr}
	values["curtime"] = []string{second}
	res, err := http.PostForm(youdaoAPI, values)
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
	r := &result{}
	er := json.Unmarshal(body, &r)
	if err != nil {
		fmt.Println(er)
		return output
	}
	if len(r.Translation) > 0 {
		output.Result = r.Translation[0]
	}
	return output
}

// func YouDao() {
// 	query := "你好"

// 	second := strconv.FormatInt(time.Now().Unix(), 10)
// 	slat := second
// 	fmt.Println(second, " 123")
// 	code := appKey + query + slat + second + appSecret
// 	h := sha256.New()
// 	h.Write([]byte(code))
// 	sign := h.Sum(nil)
// 	signStr := hex.EncodeToString(sign[:])

// 	values := make(url.Values)
// 	values["q"] = []string{query}
// 	values["from"] = []string{"auto"}
// 	values["to"] = []string{"EN"}
// 	values["signType"] = []string{"v3"}
// 	values["appKey"] = []string{appKey}
// 	values["salt"] = []string{slat}
// 	values["sign"] = []string{signStr}
// 	values["curtime"] = []string{second}
// 	res, err := http.PostForm(api, values)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	defer res.Body.Close()
// 	body, err := ioutil.ReadAll(res.Body)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	fmt.Println(string(body))
// }

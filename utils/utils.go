package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/url"
	"runtime"
	"runtime/debug"
	"server/constants"
	"server/proto"
	"sort"
	"strings"
	"time"

	"log"
)

// MarshalStr 将map转换成string
func MarshalStr(m map[string]string) string {
	byt, err := json.Marshal(m)
	if err != nil {
		log.Printf(err.Error())
		return ""
	}
	return string(byt)
}

// Marshal 将map转换成string
func Marshal(m map[string]interface{}) string {
	byt, err := json.Marshal(m)
	if err != nil {
		log.Printf(err.Error())
		return ""
	}
	return string(byt)
}

// Val 从map结构中获取值
func Val(msg map[string]interface{}, key string) string {
	if msg == nil {
		return ""
	}
	val := msg[key]
	if val == nil {
		return ""
	}
	switch val.(type) {
	case string:
		return val.(string)
	case map[string]interface{}:
		return Marshal(val.(map[string]interface{}))
	default:
		log.Printf("util.Val val=%v", val)
		return ""
	}
}

// Unmarshal 将string转换成map
func Unmarshal(str string) map[string]interface{} {
	var data map[string]interface{}
	if err := json.Unmarshal([]byte(str), &data); err != nil {
		log.Printf(err.Error())
		return data
	}
	return data
}

// Map 将数据组装成map对象
func Map(args ...interface{}) map[string]interface{} {
	if len(args)%2 != 0 {
		return nil
	}
	msg := make(map[string]interface{})
	for i := 0; i < len(args)/2; i++ {
		msg[args[2*i].(string)] = args[2*i+1]
	}
	return msg
}

// Map2 将数据组装成map对象
func Map2(args ...interface{}) map[string]string {
	if len(args)%2 != 0 {
		return nil
	}
	msg := make(map[string]string)
	for i := 0; i < len(args)/2; i++ {
		msg[args[2*i].(string)] = args[2*i+1].(string)
	}
	return msg
}

func StructToMap(arg interface{}) map[string]interface{} {
	b, _ := json.Marshal(arg)
	m := Unmarshal(string(b))
	return m
}

func DeleteStructField(arg interface{}, keys []string) map[string]interface{} {
	b, _ := json.Marshal(arg)
	m := Unmarshal(string(b))
	for _, key := range keys {
		delete(m, key)
	}
	return m
}

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

// RandStr 随机数
func RandStr(l int) string {
	bytes := make([]byte, l)
	for i := 0; i < l; i++ {
		bytes[i] = byte(randInt(65, 90))
	}
	return string(bytes)
}

// Recover 抓panic
func Recover(flag string) {
	_, _, l, _ := runtime.Caller(1)
	if err := recover(); err != nil {
		log.Printf("[%s] Recover panic line => %v", flag, l)
		log.Printf("[%s] Recover err => %v", flag, err)
		debug.PrintStack()
	}
}

func ProcessUrlString(url string) []string {
	urls := strings.Split(url, ",")
	for i, s := range urls {
		urls[i] = strings.TrimSpace(s)
	}
	return urls
}

func GenerateNatsUrlString(url string) string {
	var result string
	urls := strings.Split(url, ",")
	length := len(urls)
	for i, s := range urls {
		result += "nats://" + strings.TrimSpace(s)
		if length-1 != i {
			result += ","
		}
	}
	return result

}
func ProcessUrlStringWithHttp(url string) []string {
	urls := strings.Split(url, ",")
	for i, s := range urls {
		urls[i] = "http://" + strings.TrimSpace(s)
	}
	return urls
}
func ProcessUrlStringWithHttps(url string) []string {
	urls := strings.Split(url, ",")
	for i, s := range urls {
		urls[i] = "https://" + strings.TrimSpace(s)
	}
	return urls
}

func CheckSign(data map[string]interface{}, secretKey string) bool {
	if _, ok := data["sign"]; !ok {
		return false
	} else if data["sign"] == "" {
		return false
	}
	keys := make([]string, 0)
	for k, _ := range data {
		if k != "sign" {
			keys = append(keys, k)
		}
	}
	sort.Strings(keys)
	strA := ""
	for i, k := range keys {
		value := ""
		switch data[k].(type) {
		case float32:
			value = fmt.Sprintf("%s=%v", k, int(data[k].(float32)))
		case float64:
			value = fmt.Sprintf("%s=%v", k, int(data[k].(float64)))
		default:
			value = fmt.Sprintf("%s=%v", k, data[k])
		}
		if i == 0 {
			strA += value
		} else {
			strA += "&" + value
		}
	}
	sign := data["sign"]
	signData := strA + "&secret=" + secretKey
	urlencodeStr := UrlEncode(signData)

	if MD5(signData) != sign && MD5(urlencodeStr) != sign {
		return false
	}
	return true

}

func GenerateAccessToken(randStr string) string {
	randomStr := fmt.Sprintf("%s-%d-%d", randStr, time.Now().UnixNano(), rand.Intn(99))
	accessToken := fmt.Sprintf("%s-%s-%d", MD5(randomStr), randStr, time.Now().Add(24*30*60*time.Second).Unix())
	return accessToken
}

func GetUidFromToken(token string) string {
	params := strings.Split(token, "-")
	if len(params) >= 2 {
		return params[1]
	}
	return ""
}

func ParseQueryToMap(query string) map[string]interface{} {
	values, err := url.ParseQuery(query)
	data := make(map[string]interface{})
	if err != nil {
		return nil
	}
	for k, v := range values {
		if len(v) > 1 {
			data[k] = make([]interface{}, 0)
			for i := 0; i < len(v); i++ {
				data[k] = append(data[k].([]interface{}), v[i])
			}
		} else {
			data[k] = v[0]
		}
	}
	return data
}

func GenerateRandNum(width int) string {
	numeric := [10]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	l := len(numeric)
	rand.Seed(time.Now().UnixNano())
	var sb strings.Builder
	for i := 0; i < width; i++ {
		fmt.Fprintf(&sb, "%d", numeric[rand.Intn(l)])
	}
	return sb.String()
}

func GrpcResponse(code int, data string) *proto.Response {
	return &proto.Response{
		Code:    int32(code),
		Message: constants.ErrMsg(code),
		Data:    data,
	}
}

func LoadHTMLTemplate(path string) (string, error) {
	f, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(f), nil
}

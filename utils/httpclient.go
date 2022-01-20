package utils

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	Url "net/url"
	"sort"
	"strconv"
	"strings"
)

type HttpClient struct {
	//请求参数
	ContentType string                 //请求头
	QueryUrl    string                 //请求Url
	Params      map[string]interface{} //请求参数
	Method      string                 //请求方式 GET POST
	PostForm    bool                   //是否提交表单 只有在POST生效

	//返回参数
	Result     []byte
	Err        error
	StatusCode int
	ErrMsg     string
}

func NewHttpClient() *HttpClient {
	return new(HttpClient)
}

func (httpClient *HttpClient) HttpRequest() {
	var resp *http.Response
	var err error
	var client = new(http.Client)

	if StartWith(httpClient.QueryUrl, "https") == false && StartWith(httpClient.QueryUrl, "http") == false {
		httpClient.QueryUrl = "http://" + httpClient.QueryUrl
		if StartWith(httpClient.QueryUrl, "https") {
			tr := &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}
			client = &http.Client{Transport: tr}
		}
	}
	if httpClient.Method == "GET" {
		queryStr := MakeQueryStr(httpClient.Params)
		if strings.HasSuffix(queryStr, "&") {
			strLen := len(queryStr)
			queryStr = queryStr[:strLen-1]
		}
		httpClient.QueryUrl = httpClient.QueryUrl + "?" + queryStr
		fmt.Println(httpClient.QueryUrl)
		resp, err = client.Get(httpClient.QueryUrl)

	} else if httpClient.Method == "POST" {
		if httpClient.PostForm { //提交表单
			data := make(Url.Values)
			for k, v := range httpClient.Params {
				switch v.(type) {
				case int:
					data[k] = []string{strconv.Itoa(v.(int))}
				case int64:
					data[k] = []string{strconv.FormatInt(v.(int64), 10)}
				case string:
					data[k] = []string{v.(string)}
				case float64:
					data[k] = []string{strconv.FormatFloat(v.(float64), 'f', -1, 64)}
				case float32:
					data[k] = []string{strconv.FormatFloat(v.(float64), 'f', -1, 32)}
				default:
					data[k] = []string{v.(string)}
				}
			}
			resp, err = client.PostForm(httpClient.QueryUrl, data)

		} else if httpClient.ContentType == "Content-Type: application/json" { //提交的是json
			param, _ := json.Marshal(httpClient.Params)
			resp, err = client.Post(httpClient.QueryUrl, httpClient.ContentType, strings.NewReader(string(param)))
		} else if httpClient.ContentType == "Content-Type: application/x-www-form-urlencoded" {
			queryStr := MakeQueryStr(httpClient.Params)
			if strings.HasSuffix(queryStr, "&") {
				strLen := len(queryStr)
				queryStr = queryStr[:strLen-1]
			}
			resp, err = client.Post(httpClient.QueryUrl, httpClient.ContentType, strings.NewReader(queryStr))
		} else {
			param, _ := json.Marshal(httpClient.Params)
			resp, err = client.Post(httpClient.QueryUrl, httpClient.ContentType, strings.NewReader(string(param)))
		}
	}

	if err != nil {
		httpClient.Err = err
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	httpClient.Result = body
	httpClient.StatusCode = resp.StatusCode

	if err != nil {
		httpClient.Err = err
		return
	}

	return

}
func StartWith(url, needle string) bool {
	if strings.HasPrefix(url, needle) {
		return true
	}
	return false
}

func MakeQueryStr(params map[string]interface{}) string {
	var keys []string
	for key := range params {
		if key != "sign" {
			keys = append(keys, key)
		}
	}
	sort.Strings(keys)
	queryStr := ""
	for _, v := range keys {
		switch params[v].(type) {
		case int:
			queryStr += v + "=" + strconv.Itoa(params[v].(int)) + "&"
		case int64:
			queryStr += v + "=" + strconv.FormatInt(params[v].(int64), 10) + "&"
		case string:
			queryStr += v + "=" + params[v].(string) + "&"
		case float64:
			queryStr += v + "=" + strconv.FormatFloat(params[v].(float64), 'f', -1, 64) + "&"
		case float32:
			queryStr += v + "=" + strconv.FormatFloat(params[v].(float64), 'f', -1, 32) + "&"
		default:
			queryStr += v + "=" + params[v].(string) + "&"
		}
	}
	return queryStr
}

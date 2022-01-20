package internal

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"server/constants"
	"server/grpcclient"
	"server/log"
	"server/proto"
	"server/services/gateway/conf"
	"server/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

func PostHandler(c *gin.Context) {
	version := c.Param("v")
	service := c.Param("service")
	method := c.Param("method")
	submethod := c.Param("submethod")
	extmethod := c.Param("extmethod")
	contentType := c.GetHeader("Content-Type")

	var data map[string]interface{}
	switch contentType {
	case "application/json":
		rawData, _ := c.GetRawData()
		json.Unmarshal(rawData, &data)
	case "application/x-www-form-urlencoded":
		rawData, _ := c.GetRawData()
		data = utils.ParseQueryToMap(string(rawData))
	default:
		rawData, _ := c.GetRawData()
		json.Unmarshal(rawData, &data)
	}
	log.Debugf(utils.Marshal(data))

	callService(c, version, service, method, submethod, extmethod, data)
}

func RemoteIp(req *http.Request) string {
	remoteAddr := req.RemoteAddr
	var err error

	if ip := req.Header.Get("X-Real-IP"); ip != "" {
		remoteAddr = ip
	} else if ip = req.Header.Get("X-Forwarded-For"); ip != "" {
		remoteAddr = ip
	} else {
		remoteAddr, _, err = net.SplitHostPort(remoteAddr)
		if err != nil {
			remoteAddr = "127.0.0.1"
		}
	}
	if remoteAddr == "::1" {
		remoteAddr = "127.0.0.1"
	}

	return remoteAddr
}

func PutHandler(c *gin.Context) {

}

func GetHandler(c *gin.Context) {
	version := c.Param("v")
	service := c.Param("service")
	method := c.Param("method")
	submethod := c.Param("submethod")
	extmethod := c.Param("extmethod")
	path := c.Request.RequestURI
	query := strings.Split(path, "?")
	data := make(map[string]interface{})
	if len(query) >= 2 {
		data = utils.ParseQueryToMap(query[1])
	}
	log.Debugf(utils.Marshal(data))

	callService(c, version, service, method, submethod, extmethod, data)

}

func DeleteHandler(c *gin.Context) {

}

func getJsonMap(c *gin.Context) map[string]interface{} {
	data, err := c.GetRawData()
	if err != nil {
		log.Errorf("read no data from raw")
		return nil
	}
	if len(data) == 0 {
		return nil
	}
	var params map[string]interface{}
	err = json.Unmarshal(data, &params)
	if err != nil {
		return nil
	}
	return params
}

func response(c *gin.Context, statusCode int, resultCode int, resultMsg string, data map[string]interface{}) {
	if data == nil {
		c.JSON(statusCode, gin.H{
			"code":    resultCode,
			"message": resultMsg,
		})
	} else {
		c.JSON(statusCode, gin.H{
			"code":    resultCode,
			"message": resultMsg,
			"data":    data,
		})
	}

}

func callService(c *gin.Context, version, service, method, submethod, extmethod string, data map[string]interface{}) bool {
	if data == nil {
		data = map[string]interface{}{}
	}

	token := c.GetHeader("token")
	tokenKey := ""
	if service != "admin" {
		tokenKey = token
	} else {
		token = utils.InterfaceToString(data["token"])
		tokenKey = fmt.Sprintf("admin-token:%s", token)
	}
	params := strings.Split(token, "-")
	var uid string
	if len(params) == 3 {
		uid = params[1]
	}
	//todo token 验证暂时放开
	if uid != "" && token != "" {
		verifyUid := redisClient.Get(tokenKey)
		if verifyUid != uid {
			c.JSON(200, gin.H{
				"code":    constants.ERR_ACCESS_TOKEN,
				"message": "token error",
			})
			return false
		}
	}
	//todo sign验证暂时放开
	if utils.InterfaceToString(data["sign"]) != "" {
		if !utils.CheckSign(data, conf.Global.Secret) {
			c.JSON(200, gin.H{
				"code":    constants.ERR_SIGN,
				"message": "check sign error",
			})
			return false
		}
	}
	data["ip"] = RemoteIp(c.Request) //ger real ip
	data["token"] = token
	requestData, err := json.Marshal(data)
	if err != nil {
		c.JSON(200, gin.H{
			"code":    constants.ERR_EMPTY_DATA,
			"message": "empty data",
		})
		return false
	}
	header := proto.Header{
		Token: token,
		Ip:    RemoteIp(c.Request),
	}
	req := &proto.Request{
		Version:   version,
		Method:    method,
		Submethod: submethod,
		Extmethod: extmethod,
		Data:      string(requestData),
		Header:    &header,
	}
	r, err, ok := grpcclient.CallService(version, service, req)
	if ok {
		if err != nil {
			c.JSON(200, gin.H{
				"code":    constants.ERR_SREVICE_UNAVAILABLE,
				"message": "Service Unavailable",
			})
		} else {
			var data map[string]interface{}
			err := json.Unmarshal([]byte(r.Data), &data)
			if err != nil {
				var tempData []map[string]interface{}
				err := json.Unmarshal([]byte(r.Data), &tempData)
				if err != nil {
					c.JSON(200, gin.H{
						"code":    r.Code,
						"message": r.Message,
						"data":    r.Data,
					})
				} else {
					c.JSON(200, gin.H{
						"code":    r.Code,
						"message": r.Message,
						"data":    tempData,
					})
				}
			} else {
				c.JSON(200, gin.H{
					"code":    r.Code,
					"message": r.Message,
					"data":    data,
				})
			}
		}
	} else {
		c.JSON(200, gin.H{
			"code":    constants.ERR_SREVICE_UNKNOWN,
			"message": "Unknown service",
		})
	}
	return true
}

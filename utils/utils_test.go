package utils

import (
	"encoding/json"
	"testing"
)

func TestCheckSign(t *testing.T) {
	str := `{"age":51,"birthday":"1970-01-01","describe":"","experience":"","firstname":"","gender":1,"lastname":"","nickname":"polo","register_channel":"1","register_city":"深圳","register_platform":"ios","register_province":"广东","register_version":"1.0.0","sign":"1fa725160eca4163d1f969e22f0230fa","uid":1064,"user_type":1}`
	var data map[string]interface{}
	json.Unmarshal([]byte(str), &data)
	CheckSign(data, "348aef0470a3be44939b39ff92fbe417")

}

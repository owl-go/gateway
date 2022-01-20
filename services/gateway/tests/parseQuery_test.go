package tests

import (
	"fmt"
	"server/utils"
	"testing"
)

func TestParseQuery(t *testing.T) {
	data := utils.ParseQueryToMap("id=1&name=2&age=12&age=23")
	//for k, v := range values {
	//	fmt.Println(k, v)
	//}
	//
	//if err != nil {
	//	log.Fatal(err)
	//}
	//s := struct {
	//	Name string
	//}{
	//	Name: "hellp",
	//}
	fmt.Println(data)

}

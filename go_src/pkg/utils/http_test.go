package utils

import (
	"fmt"
	"testing"
)

func TestGetRequest(t *testing.T) {
	http := HttpRequest{}
	str, err := http.Request("GET", "https://www.baidu.com", nil).ParseBytes()
	fmt.Print(string(str))
	if err != nil {
		t.Error("请求失败")
	}
}

func ExampleRequest() {

}
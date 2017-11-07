package wxpay

import (
	"fmt"
	"testing"

	"github.com/fatih/structs"
)

func Test_Pay(t *testing.T) {
	reqDto := reqPayDto{}
	reqDto.AppId = "xiao.xinmiao"
	s := structs.New(reqDto)
	s.TagName = "json"
	param := s.Map()
	fmt.Printf("\n%+v", param)
	fmt.Printf("\n%+v", param["base"])

}

/**
 * Created by huQg on 2018/5/17,017.
 */

package myJson

import (
	"fmt"
	"testing"
)

func TestXYZ(t *testing.T) {
	fmt.Printf("testxyz\n")

	myJson := NewMS_tgC_MyJson()
	myJson.SetObjs("name", "huqing")
	myJson.SetObjs("age", 30)
	myJson.SetObjs("borthday", "19870702")

	myJsonTmp := NewMS_tgC_MyJson()
	myJsonTmp.Append("s1")
	myJsonTmp.Append("s2")
	myJsonTmp.Append("s3")

	myJson.SetObjs("school", myJsonTmp)

	str := FormatMyJson(myJson)
	fmt.Println("res1 = ", str)

	str = FormatMyJson2(myJson)
	fmt.Println("res2 = ", str)
}

const (
	textJs = `{"ip":"127.0.0.1", "port":6379, "enable": true}`
)

func Test2(t *testing.T) {
	mjs, err := ParseFromBuffer([]byte(textJs))
	if err != nil {
		fmt.Println(err)
		return
	}

	str, err := mjs.AsString("ip")
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println(str)
	}

	port, err := mjs.AsInt("port")
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println(port)
	}

	benable, err := mjs.AsBool("enable")
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println(benable)
	}
}

/**
 * Created by huQg on 2018/5/17,017.
 */

package myJson_test

import (
	"fmt"
	. "myself/myJson"
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

	str, _ := FormatMyJson(myJson)
	fmt.Println("res1 = ", str)

	str, _ = FormatMyJson2(myJson)
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

func Test3(t *testing.T) {
	mjs1 := NewMS_tgC_MyJson()
	mjs2 := NewMS_tgC_MyJson()

	mjs1.SetObjs("a", 1)
	mjs2.SetObjs("b", 3)

	mjs1.Append(mjs2)

	sjs, _ := FormatMyJson(mjs1)
	fmt.Println(sjs)
}

func TestArray(t *testing.T) {
	strAry := "[\"1\",\"2\",\"3\",\"4\",\"a\",\"b\",\"c\",\"d\",\"e\",\"f\"]"

	mJs, err := ParseFromBuffer([]byte(strAry))
	if err != nil {
		fmt.Println(err)
		return
	}

	var sRes []string
	res, err := mJs.AsArray("")
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println(res)
	}

	err = mJs.AsArrayWithResult("", &sRes)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println(sRes)
	}
}

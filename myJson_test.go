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

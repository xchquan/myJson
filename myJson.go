/**
 * Created by huQg on 2018/5/17,017.
 */

package myJson

import (
	"encoding/json"
	"fmt"
	"os"
)

type MSType int

const (
	MS_TYPE_Array = iota
	MS_TYPE_MAP
	MS_TYPE_STRING
	MS_TYPE_Int
	MS_TYPE_Int32
	MS_TYPE_Int64
	MS_TYPE_Int_C
	MS_TYPE_Float32
	MS_TYPE_Float64
	MS_TYPE_Float_C
	MS_TYPE_UNKNOW
)

type MS_tgC_MyJson struct {
	jsVal interface{}
}

/// new struct for MS_tgC_MyJson
func NewMS_tgC_MyJson() *MS_tgC_MyJson {
	mjs := &MS_tgC_MyJson{}
	return mjs
}

/// parse from buffer
func (self *MS_tgC_MyJson) ParseFromBuffer(buf []byte) error {
	err := json.Unmarshal(buf, &self.jsVal)
	return err
}

/// parse from file
func (self *MS_tgC_MyJson) ParseFromFile(strFile string) error {
	file, err := os.Open(strFile)
	if nil != err {
		return err
	}
	defer file.Close()

	/// get length for file
	stat, err := file.Stat()
	if nil != err {
		return err
	}
	var flen = stat.Size()
	buf := make([]byte, int(flen+1))
	resLen, err := file.Read(buf)
	if nil != err {
		return err
	}

	buf = buf[0:resLen]
	return self.ParseFromBuffer(buf)
}

/// get data's type
func checkDataType(itf interface{}) MSType {
	if nil == itf {
		return MS_TYPE_UNKNOW
	}
	switch itf.(type) {
	case []interface{}:
		return MS_TYPE_Array
	case map[string]interface{}:
		return MS_TYPE_MAP
	case string:
		return MS_TYPE_STRING
	case int:
		return MS_TYPE_Int
	case int32:
		return MS_TYPE_Int32
	case int64:
		return MS_TYPE_Int64
	case float32:
		return MS_TYPE_Float32
	case float64:
		return MS_TYPE_Float64
	default:
		return MS_TYPE_UNKNOW
	}

	return MS_TYPE_UNKNOW
}

func FormatMyJson(myJson *MS_tgC_MyJson) string {
	if nil == myJson {
		return ""
	}

	buf, err := json.Marshal(myJson.jsVal)
	if nil != err {
		return ""
	} else {
		return string(buf)
	}
}

/// new myjson object
func OrgMyJson(itf interface{}) *MS_tgC_MyJson {
	myjson := NewMS_tgC_MyJson()
	myjson.jsVal = itf
	return myjson
}

func FormatMyJson2(myJson *MS_tgC_MyJson) string {
	if nil == myJson {
		return ""
	}
	return fmt.Sprintf("%v", myJson)
}

///////////////////////////////////////////////////////////////////////////////////////////////////
/// get data from MyJson
func (self *MS_tgC_MyJson) ExistOfTag(sTag string) error {

	mstype := checkDataType(self.jsVal)

	switch mstype {
	case MS_TYPE_UNKNOW:
		return fmt.Errorf("val's type is not define")
	case MS_TYPE_MAP:
		break
	default:
		return fmt.Errorf("json's format isn't map")
	}

	mapval := self.jsVal.(map[string]interface{})
	if _, ok := mapval[sTag]; !ok {
		return fmt.Errorf("tag[%s] isn't exist", sTag)
	}

	return nil
}

func (self *MS_tgC_MyJson) IsArray(sTag string) error {
	mstype := checkDataType(self.jsVal)

	if mstype == MS_TYPE_Array {
		return nil
	}

	return fmt.Errorf("not array")
}

func (self *MS_tgC_MyJson) IsObjs(sTag string) error {
	mstype := checkDataType(self.jsVal)

	if mstype == MS_TYPE_MAP {
		return nil
	}

	return fmt.Errorf("not objs")
}

func (self *MS_tgC_MyJson) IsString(sTag string) error {
	mstype := checkDataType(self.jsVal)

	if mstype == MS_TYPE_STRING {
		return nil
	}

	return fmt.Errorf("not string")
}

func (self *MS_tgC_MyJson) IsInt(sTag string) error {
	mstype := checkDataType(self.jsVal)

	if mstype == MS_TYPE_Int || mstype == MS_TYPE_Int32 || mstype == MS_TYPE_Int64 {
		return nil
	}

	return fmt.Errorf("not int")
}

func (self *MS_tgC_MyJson) IsFloat(sTag string) error {
	mstype := checkDataType(self.jsVal)

	if mstype == MS_TYPE_Float32 || mstype == MS_TYPE_Float64 {
		return nil
	}

	return fmt.Errorf("not float")
}

/// get special type's data
func (self *MS_tgC_MyJson) AsArray(sTag string) ([]interface{}, error) {
	err := self.IsArray(sTag)
	if nil != err {
		return nil, err
	}

	res := self.jsVal.([]interface{})
	return res, nil
}

func (self *MS_tgC_MyJson) AsMap(sTag string) (map[string]interface{}, error) {
	err := self.IsObjs(sTag)
	if nil != err {
		return nil, err
	}

	res := self.jsVal.(map[string]interface{})
	return res, nil
}

func (self *MS_tgC_MyJson) AsString(sTag string) (string, error) {
	err := self.IsString(sTag)
	if nil != err {
		return "", err
	}

	res := self.jsVal.(string)
	return res, nil
}

func (self *MS_tgC_MyJson) AsInt(sTag string) (int64, error) {
	mstype := checkDataType(self.jsVal)

	switch mstype {
	case MS_TYPE_Int:
		res := self.jsVal.(int)
		return int64(res), nil
	case MS_TYPE_Int32:
		res := self.jsVal.(int32)
		return int64(res), nil
	case MS_TYPE_Int64:
		res := self.jsVal.(int64)
		return res, nil
	default:
		return 0, fmt.Errorf("type isn't int")
	}
}

func (self *MS_tgC_MyJson) AsFloat(sTag string) (float64, error) {
	mstype := checkDataType(self.jsVal)

	switch mstype {
	case MS_TYPE_Float32:
		res := self.jsVal.(float32)
		return float64(res), nil
	case MS_TYPE_Float64:
		res := self.jsVal.(float64)
		return res, nil
	default:
		return 0, fmt.Errorf("type isn't float")
	}
}

///////////////////////////////////////////////////////////////////////////////////////////////////
/// orgnization data for MyJson
func (self *MS_tgC_MyJson) NewArray() {
	var itf []interface{}
	self.jsVal = itf
}

func (self *MS_tgC_MyJson) NewMap() {
	self.jsVal = make(map[string]interface{})
}

func (self *MS_tgC_MyJson) Append(itfval interface{}) error {
	if nil == itfval {
		return fmt.Errorf("interface is nil")
	}

	mstype := checkDataType(self.jsVal)

	if mstype == MS_TYPE_UNKNOW {
		self.NewArray()
	} else if mstype != MS_TYPE_Array {
		return fmt.Errorf("type isn't array")
	}

	if itfJson, ok := itfval.(*MS_tgC_MyJson); ok {
		itf, _ := self.jsVal.([]interface{})
		self.jsVal = append(itf, itfJson.jsVal)
	} else if itfJson, ok := itfval.(MS_tgC_MyJson); ok {
		itf, _ := self.jsVal.([]interface{})
		self.jsVal = append(itf, itfJson.jsVal)
	} else {
		itf, _ := self.jsVal.([]interface{})
		self.jsVal = append(itf, itfval)
	}
	return nil
}

func (self *MS_tgC_MyJson) SetObjs(sTag string, itfval interface{}) error {
	if nil == itfval {
		return fmt.Errorf("interface is nil")
	}

	mstype := checkDataType(self.jsVal)

	if mstype == MS_TYPE_UNKNOW {
		self.NewMap()
	} else if mstype != MS_TYPE_MAP {
		return fmt.Errorf("type isn't array")
	}

	if itfJson, ok := itfval.(*MS_tgC_MyJson); ok {
		itf, _ := self.jsVal.(map[string]interface{})
		itf[sTag] = itfJson.jsVal
	} else if itfJson, ok := itfval.(MS_tgC_MyJson); ok {
		itf, _ := self.jsVal.(map[string]interface{})
		itf[sTag] = itfJson.jsVal
	} else {
		itf, _ := self.jsVal.(map[string]interface{})
		itf[sTag] = itfval
	}
	return nil
}

///////////////////////////////////////////////////////////////////////////////////////////////////
/// control
func (self *MS_tgC_MyJson) Clear() {
	self.jsVal = nil
}

func (self *MS_tgC_MyJson) IsNil() error {
	if nil == self.jsVal {
		return nil
	} else {
		return fmt.Errorf("items exist")
	}
}

/**
 * Created by huQg on 2018/5/17,017.
 */

package myJson

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
)

type MSType int

const (
	MS_TYPE_Array MSType = iota
	MS_TYPE_MAP
	MS_TYPE_STRING
	MS_TYPE_BOOL
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
func (slf *MS_tgC_MyJson) ParseFromBuffer(buf []byte) error {
	err := json.Unmarshal(buf, &slf.jsVal)
	return err
}

/// parse from file
func (slf *MS_tgC_MyJson) ParseFromFile(strFile string) error {

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
	return slf.ParseFromBuffer(buf)
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
	case bool:
		return MS_TYPE_BOOL
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

/// parse file
func ParseFromFile(strFile string) (*MS_tgC_MyJson, error) {
	myjson := NewMS_tgC_MyJson()
	err := myjson.ParseFromFile(strFile)
	return myjson, err
}

/// parse buffer
func ParseFromBuffer(buf []byte) (*MS_tgC_MyJson, error) {
	myjson := NewMS_tgC_MyJson()
	err := myjson.ParseFromBuffer(buf)
	return myjson, err
}

/// format string
func FormatMyJson(myJson *MS_tgC_MyJson) (string, error) {
	if nil == myJson {
		return "", fmt.Errorf("point of myjson is nil")
	}

	buf, err := json.Marshal(myJson.jsVal)
	if nil != err {
		return "", err
	}

	return string(buf), nil
}

/// new myjson object
func OrgMyJson(itf interface{}) *MS_tgC_MyJson {
	myjson := NewMS_tgC_MyJson()
	myjson.jsVal = itf
	return myjson
}

func FormatMyJson2(myJson *MS_tgC_MyJson) (string, error) {
	if nil == myJson {
		return "", fmt.Errorf("point of myjson is nil")
	}
	return fmt.Sprintf("%v", myJson), nil
}

///////////////////////////////////////////////////////////////////////////////////////////////////
/// get data from MyJson
func (slf *MS_tgC_MyJson) ExistOfTag(sTag string) (interface{}, error) {

	mstype := checkDataType(slf.jsVal)

	switch mstype {
	case MS_TYPE_UNKNOW:
		return nil, fmt.Errorf("val's type is not define")
	case MS_TYPE_MAP:
		break
	default:
		return nil, fmt.Errorf("json's format isn't map")
	}

	mapval := slf.jsVal.(map[string]interface{})
	res, ok := mapval[sTag]

	if !ok {
		return nil, fmt.Errorf("tag[%s] isn't exist", sTag)
	}

	return res, nil
}

func (slf *MS_tgC_MyJson) isArrayBase() error {
	mstype := checkDataType(slf.jsVal)
	if mstype == MS_TYPE_Array {
		return nil
	}
	return fmt.Errorf("not array")
}

func (slf *MS_tgC_MyJson) IsArray(sTag string) error {
	if len(sTag) == 0 {
		return slf.isArrayBase()
	}

	itf, err := slf.ExistOfTag(sTag)
	if err != nil {
		return err
	}

	mstype := checkDataType(itf)

	if mstype == MS_TYPE_Array {
		return nil
	}

	return fmt.Errorf("not array")
}

func (slf *MS_tgC_MyJson) IsObjs(sTag string) error {
	itf, err := slf.ExistOfTag(sTag)
	if err != nil {
		return err
	}

	mstype := checkDataType(itf)

	if mstype == MS_TYPE_MAP {
		return nil
	}

	return fmt.Errorf("not objs")
}

func (slf *MS_tgC_MyJson) IsBool(sTag string) error {
	itf, err := slf.ExistOfTag(sTag)
	if err != nil {
		return err
	}

	mstype := checkDataType(itf)

	if mstype == MS_TYPE_BOOL {
		return nil
	}

	return fmt.Errorf("not bool")
}

func (slf *MS_tgC_MyJson) IsString(sTag string) error {
	itf, err := slf.ExistOfTag(sTag)
	if err != nil {
		return err
	}

	mstype := checkDataType(itf)

	if mstype == MS_TYPE_STRING {
		return nil
	}

	return fmt.Errorf("not string")
}

func (slf *MS_tgC_MyJson) IsNumberic(sTag string) error {
	return slf.IsInt(sTag)
}

func (slf *MS_tgC_MyJson) IsInt(sTag string) error {
	itf, err := slf.ExistOfTag(sTag)
	if err != nil {
		return err
	}

	mstype := checkDataType(itf)

	if mstype == MS_TYPE_Int || mstype == MS_TYPE_Int32 || mstype == MS_TYPE_Int64 || mstype == MS_TYPE_Float32 || mstype == MS_TYPE_Float64 {
		return nil
	}

	return fmt.Errorf("not int")
}

func (slf *MS_tgC_MyJson) IsFloat(sTag string) error {
	itf, err := slf.ExistOfTag(sTag)
	if err != nil {
		return err
	}

	mstype := checkDataType(itf)

	if mstype == MS_TYPE_Float32 || mstype == MS_TYPE_Float64 {
		return nil
	}

	return fmt.Errorf("not float")
}

func (slf *MS_tgC_MyJson) asArrayBase() ([]interface{}, error) {
	if err := slf.isArrayBase(); err != nil {
		return nil, err
	}

	res := slf.jsVal.([]interface{})
	return res, nil
}

/// get special type's data
func (slf *MS_tgC_MyJson) AsArray(sTag string) ([]interface{}, error) {
	if len(sTag) == 0 {
		return slf.asArrayBase()
	}

	itf, err := slf.ExistOfTag(sTag)
	if err != nil {
		return nil, err
	}

	mstype := checkDataType(itf)

	if mstype != MS_TYPE_Array {
		return nil, fmt.Errorf("not array")
	}

	res := itf.([]interface{})
	return res, nil
}

func (slf *MS_tgC_MyJson) AsArrayWithResult(sTag string, result interface{}) error {

	var res []interface{}
	var err error

	if len(sTag) == 0 {
		res, err = slf.asArrayBase()
	} else {
		res, err = slf.AsArray(sTag)
	}

	if err != nil {
		return err
	}

	resultv := reflect.ValueOf(result)
	if resultv.Kind() != reflect.Ptr || resultv.Elem().Kind() != reflect.Slice {
		panic("result argument must be a slice address")
	}

	slicev := resultv.Elem()
	elemt := slicev.Type().Elem()

	for _, v := range res {
		elemtp := reflect.New(elemt)
		elemtp.Elem().Set(reflect.ValueOf(v))
		slicev = reflect.Append(slicev, elemtp.Elem())
	}
	resultv.Elem().Set(slicev)
	return nil
}

func (slf *MS_tgC_MyJson) AsMap(sTag string) (map[string]interface{}, error) {
	itf, err := slf.ExistOfTag(sTag)
	if err != nil {
		return nil, err
	}

	mstype := checkDataType(itf)

	if mstype != MS_TYPE_MAP {
		return nil, fmt.Errorf("not objs")
	}

	res := itf.(map[string]interface{})
	return res, nil
}

func (slf *MS_tgC_MyJson) AsBool(sTag string) (bool, error) {
	itf, err := slf.ExistOfTag(sTag)
	if err != nil {
		return false, err
	}

	mstype := checkDataType(itf)
	if mstype != MS_TYPE_BOOL {
		return false, fmt.Errorf("not bool")
	}

	res := itf.(bool)
	return res, nil
}

func (slf *MS_tgC_MyJson) AsString(sTag string) (string, error) {
	itf, err := slf.ExistOfTag(sTag)
	if err != nil {
		return "", err
	}

	mstype := checkDataType(itf)

	if mstype != MS_TYPE_STRING {
		return "", fmt.Errorf("not string")
	}

	res := itf.(string)
	return res, nil
}

func (slf *MS_tgC_MyJson) AsInt(sTag string) (int64, error) {
	itf, err := slf.ExistOfTag(sTag)
	if err != nil {
		return 0, err
	}

	mstype := checkDataType(itf)

	switch mstype {
	case MS_TYPE_Int:
		res := itf.(int)
		return int64(res), nil
	case MS_TYPE_Int32:
		res := itf.(int32)
		return int64(res), nil
	case MS_TYPE_Int64:
		res := itf.(int64)
		return res, nil
	case MS_TYPE_Float32:
		res := itf.(float32)
		return int64(res), nil
	case MS_TYPE_Float64:
		res := itf.(float64)
		return int64(res), nil
	default:
		return 0, fmt.Errorf("type isn't int")
	}
}

func (slf *MS_tgC_MyJson) AsFloat(sTag string) (float64, error) {
	itf, err := slf.ExistOfTag(sTag)
	if err != nil {
		return 0, err
	}

	mstype := checkDataType(itf)

	switch mstype {
	case MS_TYPE_Float32:
		res := itf.(float32)
		return float64(res), nil
	case MS_TYPE_Float64:
		res := itf.(float64)
		return res, nil
	default:
		return 0, fmt.Errorf("type isn't float")
	}
}

///////////////////////////////////////////////////////////////////////////////////////////////////
/// orgnization data for MyJson
func (slf *MS_tgC_MyJson) NewArray() {
	var itf []interface{}
	slf.jsVal = itf
}

func (slf *MS_tgC_MyJson) NewMap() {
	slf.jsVal = make(map[string]interface{})
}

func (slf *MS_tgC_MyJson) Append(itfval interface{}) error {
	if nil == itfval {
		return fmt.Errorf("interface is nil")
	}

	mstype := checkDataType(slf.jsVal)

	if mstype == MS_TYPE_UNKNOW {
		slf.NewArray()
	} else if mstype != MS_TYPE_Array {
		return fmt.Errorf("type isn't array")
	}

	if itfJson, ok := itfval.(*MS_tgC_MyJson); ok {
		itf, _ := slf.jsVal.([]interface{})
		slf.jsVal = append(itf, itfJson.jsVal)
	} else if itfJson, ok := itfval.(MS_tgC_MyJson); ok {
		itf, _ := slf.jsVal.([]interface{})
		slf.jsVal = append(itf, itfJson.jsVal)
	} else {
		itf, _ := slf.jsVal.([]interface{})
		slf.jsVal = append(itf, itfval)
	}
	return nil
}

func (slf *MS_tgC_MyJson) SetObjs(sTag string, itfval interface{}) error {
	if nil == itfval {
		return fmt.Errorf("interface is nil")
	}

	mstype := checkDataType(slf.jsVal)

	if mstype == MS_TYPE_UNKNOW {
		slf.NewMap()
	} else if mstype != MS_TYPE_MAP {
		return fmt.Errorf("type isn't array")
	}

	if itfJson, ok := itfval.(*MS_tgC_MyJson); ok {
		itf, _ := slf.jsVal.(map[string]interface{})
		itf[sTag] = itfJson.jsVal
	} else if itfJson, ok := itfval.(MS_tgC_MyJson); ok {
		itf, _ := slf.jsVal.(map[string]interface{})
		itf[sTag] = itfJson.jsVal
	} else {
		itf, _ := slf.jsVal.(map[string]interface{})
		itf[sTag] = itfval
	}
	return nil
}

///////////////////////////////////////////////////////////////////////////////////////////////////
/// control
func (slf *MS_tgC_MyJson) Clear() {
	slf.jsVal = nil
}

func (slf *MS_tgC_MyJson) IsNil() error {
	if nil == slf.jsVal {
		return nil
	} else {
		return fmt.Errorf("items exist")
	}
}

package main

import (
	"fmt"
	"strconv"
	"strings"

	"customVariableExample/configVar"
)

func main() {
	var err error

	var exBool bool
	err = configVar.SetConfig("ON", false, &exBool, &configVar.SwitchType{})
	if err != nil {
		fmt.Printf("輸出結果: %v (%T), 錯誤提示: %v\n", exBool, exBool, err)
	} else {
		fmt.Printf("輸出結果: %v (%T)\n", exBool, exBool)
	} //輸出結果: true (bool)

	var exInt int8
	err = configVar.SetConfig("60.1", "90", &exInt, &configVar.SecondsInADay{})
	if err != nil {
		fmt.Printf("輸出結果: %v (%T), 錯誤提示: %v\n", exInt, exInt, err)
	} else {
		fmt.Printf("輸出結果: %v (%T)\n", exInt, exInt)
	} //輸出結果: 60 (int8)

	var exUint uint64
	err = configVar.SetConfig("-99.9", "11", &exUint, &customStruct{})
	if err != nil {
		fmt.Printf("輸出結果: %v (%T), 錯誤提示: %v\n", exUint, exUint, err)
	} else {
		fmt.Printf("輸出結果: %v (%T)\n", exUint, exUint)
	} //輸出結果: 11 (uint64), 錯誤提示: input value > '-99.9' > 不可低於10 strconv.ParseUint: parsing "-99": invalid syntax ; default value > <nil>

	var exSlice []uint32
	exSlice = append(exSlice, 123)
	err = configVar.SetConfig("456", "111", &exSlice, &configVar.AddStringSlice{})
	if err != nil {
		fmt.Printf("輸出結果: %v (%T), 錯誤提示: %v\n", exSlice, exSlice, err)
	} else {
		fmt.Printf("輸出結果: %v (%T)\n", exSlice, exSlice)
	} //輸出結果: [123 456] ([]uint32)

	var exCustomTypeSlice []interface{}
	exCustomTypeSlice = append(exCustomTypeSlice, 456)
	err = configVar.SetConfig("Test", "Def", &exCustomTypeSlice, &configVar.AddStringSlice{})
	if err != nil {
		fmt.Printf("輸出結果: %v (%T), 錯誤提示: %v\n", exCustomTypeSlice, exCustomTypeSlice, err)
	} else {
		fmt.Printf("輸出結果: %v (%T)\n", exCustomTypeSlice, exCustomTypeSlice)
	} //輸出結果: [456 Test] ([]interface {})

	exStringMap := make(map[string]interface{})
	exStringMap["原始Key"] = "原始Value"
	err = configVar.SetConfig("VVVVVVVVVVVVV", exStringMap, &exStringMap, &stringMapType{SetKey: "KeyTest"})
	if err != nil {
		fmt.Printf("輸出結果: %v (%T), 錯誤提示: %v\n", exStringMap, exStringMap, err)
	} else {
		fmt.Printf("輸出結果: %v (%T)\n", exStringMap, exStringMap)
	} //輸出結果: map[KeyTest:VVVVVVVVVVVVV 原始Key:原始Value] (map[string]interface {})

	var exCustomTypeMap interface{}
	err = configVar.SetConfig("testKey|t1|t2|t3|t4", nil, &exCustomTypeMap, &customTypeMap{})
	if err != nil {
		fmt.Printf("輸出結果: %v (%T), 錯誤提示: %v\n", exCustomTypeMap, exCustomTypeMap, err)
	} else {
		exCustomTypeMapAssertion := exCustomTypeMap.(map[string]*CustomTypeValue)
		for index, value := range exCustomTypeMapAssertion {
			fmt.Printf("輸出結果: %v %v (%T)\n", index, value, value)
		}
	} //輸出結果: testKey &{t1 t2 t3 t4 } (*main.CustomTypeValue)

}

type customStruct struct{}

func (_ customStruct) GetValue(inputValue string) (output interface{}, err error) {

	var cache int
	if index := strings.Index(inputValue, "."); index < 1 { //無條件捨去小數，直接直接取整數
		cache, err = strconv.Atoi(inputValue)
	} else {
		cache, err = strconv.Atoi(inputValue[:index])
	}

	if err != nil {
		return cache, err
	} else if cache < 10 {
		return cache, fmt.Errorf("'%v' > 不可低於10", inputValue)
	} else {
		return cache, nil
	}
}

type stringMapType struct {
	SetKey string
}

func (mp *stringMapType) GetValue(inputValue string) (output interface{}, err error) {
	if len(inputValue) < 10 {
		return nil, fmt.Errorf("'%v' > 字串長度不可低於10", inputValue)
	}

	cache := make(map[string]interface{})
	cache[mp.SetKey] = inputValue
	return cache, nil
}

type customTypeMap map[string]*CustomTypeValue

type CustomTypeValue struct {
	f1 string
	f2 string
	f3 string
	f4 string
	f5 string
}

func (mp customTypeMap) GetValue(inputValue string) (output interface{}, err error) {
	cacheSlice := strings.Split(inputValue, "|")
	if len(cacheSlice) < 5 {
		return nil, fmt.Errorf("'%v' > ' | ' 分隔號數量不可低於5", inputValue)
	} else if cacheSlice[1] == "" || cacheSlice[2] == "" || cacheSlice[3] == "" || cacheSlice[4] == "" {
		return nil, fmt.Errorf("'%v' > 部份數值不可為空", inputValue)
	}

	//cacheMap := make(map[string]interface{})
	cacheMap := make(map[string]*CustomTypeValue)
	cacheMap[cacheSlice[0]] = &CustomTypeValue{f1: cacheSlice[1], f2: cacheSlice[2], f3: cacheSlice[3], f4: cacheSlice[4]}
	return cacheMap, nil
}

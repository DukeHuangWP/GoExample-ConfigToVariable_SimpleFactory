package customVar

import (
	"fmt"
	"net"
	"net/url"
	"strconv"
	"strings"
)

type BaseSlice struct{} //僅轉換Golnag slice ([]int,[]uint,[]float,[]string,[]bool,[]interface) ，無做任何檢查或處理
func (_ BaseSlice) GetValue(inputValue string) (output interface{}, err error) {
	return inputValue, nil
}

type IntType struct{} //純轉換Golnag  int，無做任何檢查或處理
func (_ IntType) GetValue(inputValue string) (output interface{}, err error) {
	_, err = strconv.Atoi(inputValue)
	if err != nil {
		return nil, err
	}
	return inputValue, nil
}

type Int8Type struct{} //純轉換Golnag int8，無做任何檢查或處理
func (_ Int8Type) GetValue(inputValue string) (output interface{}, err error) {
	_, err = strconv.ParseInt(inputValue, 10, 8)
	if err != nil {
		return nil, err
	}
	return inputValue, nil
}

type Int16Type struct{} //純轉換Golnag int16，無做任何檢查或處理
func (_ Int16Type) GetValue(inputValue string) (output interface{}, err error) {
	_, err = strconv.ParseInt(inputValue, 10, 16)
	if err != nil {
		return nil, err
	}
	return inputValue, nil
}

type Int32Type struct{} //純轉換Golnag int32，無做任何檢查或處理
func (_ Int32Type) GetValue(inputValue string) (output interface{}, err error) {
	_, err = strconv.ParseInt(inputValue, 10, 32)
	if err != nil {
		return nil, err
	}
	return inputValue, nil
}

type Int64Type struct{} //純轉換Golnag int64，無做任何檢查或處理
func (_ Int64Type) GetValue(inputValue string) (output interface{}, err error) {
	_, err = strconv.ParseInt(inputValue, 10, 64)
	if err != nil {
		return nil, err
	}
	return inputValue, nil
}

type UintType struct{} //純轉換Golnag uint，無做任何檢查或處理
func (_ UintType) GetValue(inputValue string) (output interface{}, err error) {
	_, err = strconv.ParseUint(inputValue, 10, 32)
	if err != nil {
		return nil, err
	}
	return inputValue, nil
}

type Uint8Type struct{} //純轉換Golnag uint8，無做任何檢查或處理
func (_ Uint8Type) GetValue(inputValue string) (output interface{}, err error) {
	_, err = strconv.ParseUint(inputValue, 10, 8)
	if err != nil {
		return nil, err
	}
	return inputValue, nil
}

type Uint16Type struct{} //純轉換Golnag uint16，無做任何檢查或處理
func (_ Uint16Type) GetValue(inputValue string) (output interface{}, err error) {
	_, err = strconv.ParseUint(inputValue, 10, 16)
	if err != nil {
		return nil, err
	}
	return inputValue, nil
}

type Uint32Type struct{} //純轉換Golnag uint32，無做任何檢查或處理
func (_ Uint32Type) GetValue(inputValue string) (output interface{}, err error) {
	_, err = strconv.ParseUint(inputValue, 10, 32)
	if err != nil {
		return nil, err
	}
	return inputValue, nil
}

type Uint64Type struct{} //純轉換Golnag uint64，無做任何檢查或處理
func (_ Uint64Type) GetValue(inputValue string) (output interface{}, err error) {
	_, err = strconv.ParseUint(inputValue, 10, 16)
	if err != nil {
		return nil, err
	}
	return inputValue, nil
}

type Float32Type struct{} //純轉換Golnag float32，無做任何檢查或處理
func (_ Float32Type) GetValue(inputValue string) (output interface{}, err error) {
	_, err = strconv.ParseFloat(inputValue, 32)
	if err != nil {
		return nil, err
	}
	return inputValue, nil
}

type Float64Type struct{} //純轉換Golnag float64，無做任何檢查或處理
func (_ Float64Type) GetValue(inputValue string) (output interface{}, err error) {
	_, err = strconv.ParseFloat(inputValue, 64)
	if err != nil {
		return nil, err
	}
	return inputValue, nil
}

type StringType struct{} //純轉換Golnagstring，無做任何檢查或處理
func (_ StringType) GetValue(inputValue string) (output interface{}, err error) {
	return inputValue, nil
}

type SwitchType struct{} //bool，"ON","On","on" 回傳true,其餘參照strconv.ParseBool
func (_ SwitchType) GetValue(inputValue string) (output interface{}, err error) {
	if inputValue == "on" || inputValue == "On" || inputValue == "ON" || inputValue == "開" {
		return true, nil
	} else {
		cache, _ := strconv.ParseBool(inputValue)
		return cache, nil
	}
}

type ValidHostURL struct{} //string，host網域判斷(不包含Scheme)
func (_ ValidHostURL) GetValue(inputValue string) (output interface{}, err error) {
	_, err = url.Parse(inputValue)
	if err != nil { //不合法網址
		return nil, fmt.Errorf("'%v' > 不合法的網址", inputValue)
	} else {
		return inputValue, nil
	}
}

type ValidWebURL struct{} //string，Web網址判斷
func (_ ValidWebURL) GetValue(inputValue string) (output interface{}, err error) {
	cache, err := url.Parse(inputValue)
	if err != nil { //不合法網址
		return nil, fmt.Errorf("'%v' > 不合法的網址", inputValue)
	} else {
		if cache.Scheme != "" && cache.Host != "" { //合法且有效網址
			return inputValue, nil
		} else { //合法但無效網址
			return nil, fmt.Errorf("'%v' > 合法但無效網址", inputValue)
		}
	}
}

type ValidIPnPort struct{} //string，IPv4可帶port，例如:"192.168.0.1:1234"
func (_ ValidIPnPort) GetValue(inputValue string) (output interface{}, err error) {

	var portString string
	var ipString string
	portIndex := strings.LastIndex(inputValue, ":")
	if portIndex > 1 {
		portUInt, err := strconv.ParseUint(inputValue[portIndex+1:], 10, 16) //0~65535
		if err != nil {
			return "", fmt.Errorf("不合法的Port")
		} //0~65535
		portString = ":" + fmt.Sprint(portUInt)
		ipString = inputValue[:portIndex]
	} else {
		ipString = inputValue
	}

	ipNet := net.ParseIP(ipString)
	if ipNet == nil {
		return "", fmt.Errorf("不合法的IP")
	}

	if ipNet.To4() == nil || ipNet.To16() == nil { //ipv4 or ipv6
		return "", fmt.Errorf("不合法的IPv4,IPv6")
	}
	return ipNet.String() + portString, nil
}

func StringToNumber(number string) (int, error) {
	if index := strings.Index(number, "."); index < 1 {
		cache, err := strconv.Atoi(number)
		return cache, err
	} else {
		cache, err := strconv.Atoi(number[:index])
		return cache, err
	}
}

type MilliSecsInADay struct{} //int毫秒，通常用於較精密計時，不可超過一小時或低於1ms
func (_ MilliSecsInADay) GetValue(inputValue string) (output interface{}, err error) {
	cache, err := StringToNumber(inputValue)
	if err != nil {
		return nil, err
	}
	if cache < 1 { //不可低於1ms
		return nil, fmt.Errorf("'%v' > 不可低於1ms", inputValue)
	}
	if cache > 3600000 { //不可超過1小時
		return nil, fmt.Errorf("'%v' > 不可超過1小時", inputValue)
	}
	return cache, nil
}

type MinutesInADay struct{} //int分鐘，不可超過一天或低於1分鐘
func (_ MinutesInADay) GetValue(inputValue string) (output interface{}, err error) {
	cache, err := StringToNumber(inputValue)
	if err != nil {
		return nil, err
	}
	if cache < 1 { //不可低於1ms
		return nil, fmt.Errorf("'%v' > 不可低於1ms", inputValue)
	}
	if cache > 3600000 { //不可超過1小時
		return nil, fmt.Errorf("'%v' > 不可超過1小時", inputValue)
	}
	return cache, nil
}

type HoursInADay struct{} //int小時，不可超過一天或低於1小時
func (_ HoursInADay) GetValue(inputValue string) (output interface{}, err error) {
	cache, err := StringToNumber(inputValue)
	if err != nil {
		return nil, err
	}
	if cache < 1 { //不可低於1小時
		return nil, fmt.Errorf("'%v' > 不可低於1小時", inputValue)
	}
	if cache > 24 { //不可超過一天
		return nil, fmt.Errorf("'%v' > 不可超過一天", inputValue)
	}
	return cache, nil
}

type DaysInAWeek struct{} //int天數，不可超過七天或低於一天
func (_ DaysInAWeek) GetValue(inputValue string) (output interface{}, err error) {
	cache, err := StringToNumber(inputValue)
	if err != nil {
		return nil, err
	}
	if cache < 1 { //不可低於一天
		return nil, fmt.Errorf("'%v' > 不可低於一天", inputValue)
	}

	if cache > 7 { //不可超過七天
		return nil, fmt.Errorf("'%v' > 不可超過七天", inputValue)
	}
	return cache, nil
}

type DaysInAMonth struct{} //int天數，不可超過31天或低於一天
func (_ DaysInAMonth) GetValue(inputValue string) (output interface{}, err error) {
	cache, err := StringToNumber(inputValue)
	if err != nil {
		return nil, err
	}
	if cache < 1 { //不可低於一天
		return nil, fmt.Errorf("'%v' > 不可低於一天", inputValue)
	}

	if cache > 30 { //不可超過31天
		return nil, fmt.Errorf("'%v' > 不可超過31天", inputValue)
	}
	return cache, nil
}

type DaysInHalfAYear struct{} //int天數，不可超過180天或低於30天
func (_ DaysInHalfAYear) GetValue(inputValue string) (output interface{}, err error) {
	cache, err := StringToNumber(inputValue)
	if err != nil {
		return nil, err
	}
	if cache < 30 { //不可低於30天
		return nil, fmt.Errorf("'%v' > 不可低於30天", inputValue)
	}

	if cache > 180 { //不可超過180天
		return nil, fmt.Errorf("'%v' > 不可超過180天", inputValue)
	}
	return cache, nil
}

type SecondsInADay struct{} //輸出單位秒，不可超過一天或低於2秒
func (_ SecondsInADay) GetValue(inputValue string) (output interface{}, err error) {
	cache, err := StringToNumber(inputValue) //無條件捨去小數，直接直接取整數
	if err != nil {
		return cache, err
	} else if cache < 2 {
		return cache, fmt.Errorf("'%v' > 不可低於2秒", inputValue)
	} else if cache > 86400 {
		return cache, fmt.Errorf("'%v' > 不可超過一天", inputValue)
	} else {
		return cache, nil
	}
}

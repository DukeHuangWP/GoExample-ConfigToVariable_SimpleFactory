package customVar

import (
	"fmt"
	"net"
	"net/url"
	"strconv"
	"strings"
)

type AddStringSlice struct{}

func (_ AddStringSlice) GetValue(inputValue string) (output interface{}, err error) {
	return inputValue, nil
}

type IntType struct{} //純轉換Golnag  int，無做任何處理
func (_ IntType) GetValue(inputValue string) (output interface{}, err error) {
	return inputValue, nil
}

type Int8Type struct{} //純轉換Golnag int8，無做任何處理
func (_ Int8Type) GetValue(inputValue string) (output interface{}, err error) {
	return inputValue, nil
}

type Int16Type struct{} //純轉換Golnag int16，無做任何處理
func (_ Int16Type) GetValue(inputValue string) (output interface{}, err error) {
	return inputValue, nil
}

type Int32Type struct{} //純轉換Golnag int32，無做任何處理
func (_ Int32Type) GetValue(inputValue string) (output interface{}, err error) {
	return inputValue, nil
}

type Int64Type struct{} //純轉換Golnag int64，無做任何處理
func (_ Int64Type) GetValue(inputValue string) (output interface{}, err error) {
	return inputValue, nil
}

type UintType struct{} //純轉換Golnag uint，無做任何處理
func (_ UintType) GetValue(inputValue string) (output interface{}, err error) {
	return inputValue, nil
}

type Uint8Type struct{} //純轉換Golnag uint8，無做任何處理
func (_ Uint8Type) GetValue(inputValue string) (output interface{}, err error) {
	return inputValue, nil
}

type Uint16Type struct{} //純轉換Golnag uint16，無做任何處理
func (_ Uint16Type) GetValue(inputValue string) (output interface{}, err error) {
	return inputValue, nil
}

type Uint32Type struct{} //純轉換Golnag uint32，無做任何處理
func (_ Uint32Type) GetValue(inputValue string) (output interface{}, err error) {
	return inputValue, nil
}

type Uint64Type struct{} //純轉換Golnag uint64，無做任何處理
func (_ Uint64Type) GetValue(inputValue string) (output interface{}, err error) {
	return inputValue, nil
}

type Float32Type struct{} //純轉換Golnag float32，無做任何處理
func (_ Float32Type) GetValue(inputValue string) (output interface{}, err error) {
	return inputValue, nil
}

type Float64Type struct{} //純轉換Golnag float64，無做任何處理
func (_ Float64Type) GetValue(inputValue string) (output interface{}, err error) {
	return inputValue, nil
}

type StringType struct{} //純轉換Golnagstring，無做任何處理
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

type HostURL struct{} //string，網址判斷
func (_ HostURL) GetValue(inputValue string) (output interface{}, err error) {
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

type HostIP struct{} //string，IPv4可帶port，例如:"192.168.0.1:1234"
func (_ HostIP) GetValue(inputValue string) (output interface{}, err error) {
	//暫時不考慮ipv6
	cache := strings.Split(inputValue, ":")
	count := len(cache)
	if count == 1 {
		address := net.ParseIP(inputValue)
		if address == nil { //不合法IP
			return nil, fmt.Errorf("'%v' > 不合法的IP", inputValue)
		} else {
			return address.String(), nil
		}
	} else if count == 2 {
		port, err := strconv.ParseUint(cache[1], 10, 32)
		if err == nil {
			address := net.ParseIP(cache[0])
			if address != nil || port != 0 { //不合法IP
				return address.String() + ":" + fmt.Sprint(port), nil
			}
		}
		return nil, fmt.Errorf("'%v' > 不合法的Port", inputValue)
	} else {
		return nil, fmt.Errorf("'%v' > 不合法的IP:Port", inputValue)
	}
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

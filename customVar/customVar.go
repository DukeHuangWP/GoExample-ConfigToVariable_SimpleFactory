package customVar

import (
	"fmt"
	"strconv"
)

//自訂型別介面
type CustomTypes interface {
	GetValue(string) (interface{}, error) //func設為全域，提供使用者可任意對介面繼承、方法重載
}

/*
將輸入值轉換成自訂型別並寫入到目標變數
@param inputValue: 外部輸入變數(須先轉成字串)
@param defaultValue: 預設值(當自定型別轉換發生錯誤時則輸出此值)
@param outputPointer:  寫入輸出變數的指標
@param Custom:  自訂型別，可OOP繼承與方法重載
@return error : 錯誤提示
*/
func SetConfig(inputValue string, defaultValue, outputPointer interface{}, Custom CustomTypes) (err error) {
	return configToVariable(inputValue, defaultValue, outputPointer, Custom)
}

/*
將特定值值轉換成自訂型別並寫入到目標變數
@param inputValue: 特定值
@param Custom:  自訂型別，可OOP繼承與方法重載
@return error : 錯誤提示
*/
func SetValue(inputValue, outputPointer interface{}, Custom CustomTypes) (err error) {
	err = configToVariable(fmt.Sprint(inputValue), inputValue, outputPointer, Custom)
	return err
}

func configToVariable(inputValue string, defaultValue, outputPointer interface{}, Custom CustomTypes) (err error) {
	outputVal, errVal := Custom.GetValue(inputValue)

	switch outputPointer.(type) {
	case *string:
		outputPtr := outputPointer.(*string)
		if errVal == nil {
			*outputPtr = fmt.Sprint(outputVal)
			return nil
		} else {
			*outputPtr = fmt.Sprint(defaultValue)
			return errVal
		}
	case *bool:
		outputPtr := outputPointer.(*bool)
		outputType, errType := strconv.ParseBool(fmt.Sprint(outputVal))
		if errVal == nil && errType == nil {
			*outputPtr = outputType
			return nil
		} else {
			outputDef, errDef := strconv.ParseBool(fmt.Sprint(defaultValue))
			*outputPtr = outputDef
			if errVal == nil && errDef == nil {
				return errType
			} else {
				return fmt.Errorf("input value > %v %v ; default value > %v", errVal, errType, errDef)
			}
		}

	case *int:
		outputPtr := outputPointer.(*int)
		outputType, errType := strconv.Atoi(fmt.Sprint(outputVal))
		if errVal == nil && errType == nil {
			*outputPtr = outputType
			return nil
		} else {
			outputDef, errDef := strconv.Atoi(fmt.Sprint(defaultValue))
			*outputPtr = outputDef
			if errVal == nil && errDef == nil {
				return errType
			} else {
				return fmt.Errorf("input value > %v %v ; default value > %v", errVal, errType, errDef)
			}
		}
	case *int8:
		outputPtr := outputPointer.(*int8)
		outputType, errType := strconv.ParseInt(fmt.Sprint(outputVal), 10, 8)
		if errVal == nil && errType == nil {
			*outputPtr = int8(outputType)
			return nil
		} else {
			outputDef, errDef := strconv.ParseInt(fmt.Sprint(defaultValue), 10, 8)
			*outputPtr = int8(outputDef)
			if errVal == nil && errDef == nil {
				return errType
			} else {
				return fmt.Errorf("input value > %v %v ; default value > %v", errVal, errType, errDef)
			}
		}
	case *int16:
		outputPtr := outputPointer.(*int16)
		outputType, errType := strconv.ParseInt(fmt.Sprint(outputVal), 10, 16)
		if errVal == nil && errType == nil {
			*outputPtr = int16(outputType)
			return nil
		} else {
			outputDef, errDef := strconv.ParseInt(fmt.Sprint(defaultValue), 10, 16)
			*outputPtr = int16(outputDef)
			if errVal == nil && errDef == nil {
				return errType
			} else {
				return fmt.Errorf("input value > %v %v ; default value > %v", errVal, errType, errDef)
			}
		}
	case *int32:
		outputPtr := outputPointer.(*int32)
		outputType, errType := strconv.ParseInt(fmt.Sprint(outputVal), 10, 32)
		if errVal == nil && errType == nil {
			*outputPtr = int32(outputType)
			return nil
		} else {
			outputDef, errDef := strconv.ParseInt(fmt.Sprint(defaultValue), 10, 32)
			*outputPtr = int32(outputDef)
			if errVal == nil && errDef == nil {
				return errType
			} else {
				return fmt.Errorf("input value > %v %v ; default value > %v", errVal, errType, errDef)
			}
		}
	case *int64:
		outputPtr := outputPointer.(*int64)
		outputType, errType := strconv.ParseInt(fmt.Sprint(outputVal), 10, 64)
		if errVal == nil && errType == nil {
			*outputPtr = outputType
			return nil
		} else {
			outputDef, errDef := strconv.ParseInt(fmt.Sprint(defaultValue), 10, 64)
			*outputPtr = outputDef
			if errVal == nil && errDef == nil {
				return errType
			} else {
				return fmt.Errorf("input value > %v %v ; default value > %v", errVal, errType, errDef)
			}
		}

	case *uint:
		outputPtr := outputPointer.(*uint)
		outputType, errType := strconv.ParseUint(fmt.Sprint(outputVal), 10, 32)
		if errVal == nil && errType == nil {
			*outputPtr = uint(outputType)
			return nil
		} else {
			outputDef, errDef := strconv.ParseUint(fmt.Sprint(defaultValue), 10, 32)
			*outputPtr = uint(outputDef)
			if errVal == nil && errDef == nil {
				return errType
			} else {
				return fmt.Errorf("input value > %v %v ; default value > %v", errVal, errType, errDef)
			}
		}
	case *uint8:
		outputPtr := outputPointer.(*uint8)
		outputType, errType := strconv.ParseUint(fmt.Sprint(outputVal), 10, 8)
		if errVal == nil && errType == nil {
			*outputPtr = uint8(outputType)
			return nil
		} else {
			outputDef, errDef := strconv.ParseUint(fmt.Sprint(defaultValue), 10, 8)
			*outputPtr = uint8(outputDef)
			if errVal == nil && errDef == nil {
				return errType
			} else {
				return fmt.Errorf("input value > %v %v ; default value > %v", errVal, errType, errDef)
			}
		}
	case *uint16:
		outputPtr := outputPointer.(*uint16)
		outputType, errType := strconv.ParseUint(fmt.Sprint(outputVal), 10, 16)
		if errVal == nil && errType == nil {
			*outputPtr = uint16(outputType)
			return nil
		} else {
			outputDef, errDef := strconv.ParseUint(fmt.Sprint(defaultValue), 10, 16)
			*outputPtr = uint16(outputDef)
			if errVal == nil && errDef == nil {
				return errType
			} else {
				return fmt.Errorf("input value > %v %v ; default value > %v", errVal, errType, errDef)
			}
		}
	case *uint32:
		outputPtr := outputPointer.(*uint32)
		outputType, errType := strconv.ParseUint(fmt.Sprint(outputVal), 10, 32)
		if errVal == nil && errType == nil {
			*outputPtr = uint32(outputType)
			return nil
		} else {
			outputDef, errDef := strconv.ParseUint(fmt.Sprint(defaultValue), 10, 32)
			*outputPtr = uint32(outputDef)
			if errVal == nil && errDef == nil {
				return errType
			} else {
				return fmt.Errorf("input value > %v %v ; default value > %v", errVal, errType, errDef)
			}
		}
	case *uint64:
		outputPtr := outputPointer.(*uint64)
		outputType, errType := strconv.ParseUint(fmt.Sprint(outputVal), 10, 64)
		if errVal == nil && errType == nil {
			*outputPtr = outputType
			return nil
		} else {
			outputDef, errDef := strconv.ParseUint(fmt.Sprint(defaultValue), 10, 64)
			*outputPtr = outputDef
			if errVal == nil && errDef == nil {
				return errType
			} else {
				return fmt.Errorf("input value > %v %v ; default value > %v", errVal, errType, errDef)
			}
		}

	case *float32:
		outputPtr := outputPointer.(*float32)
		outputType, errType := strconv.ParseFloat(fmt.Sprint(outputVal), 32)
		if errVal == nil && errType == nil {
			*outputPtr = float32(outputType)
			return nil
		} else {
			outputDef, errDef := strconv.ParseFloat(fmt.Sprint(defaultValue), 32)
			*outputPtr = float32(outputDef)
			if errVal == nil && errDef == nil {
				return errType
			} else {
				return fmt.Errorf("input value > %v %v ; default value > %v", errVal, errType, errDef)
			}
		}
	case *float64:
		outputPtr := outputPointer.(*float64)
		outputType, errType := strconv.ParseFloat(fmt.Sprint(outputVal), 64)
		if errVal == nil && errType == nil {
			*outputPtr = outputType
			return nil
		} else {
			outputDef, errDef := strconv.ParseFloat(fmt.Sprint(defaultValue), 64)
			*outputPtr = outputDef
			if errVal == nil && errDef == nil {
				return errType
			} else {
				return fmt.Errorf("input value > %v %v ; default value > %v", errVal, errType, errDef)
			}
		}
	case *[]int:
		outputPtr := outputPointer.(*[]int)
		if errVal == nil {
			outputSliceType, ok := outputVal.([]int)
			if ok {
				(*outputPtr) = append(*outputPtr, outputSliceType...)
				return nil
			}
		} else {
			outputSliceType, ok := defaultValue.([]int)
			if ok {
				(*outputPtr) = append(*outputPtr, outputSliceType...)
				return errVal
			}
		}
		outputType, errType := strconv.Atoi(fmt.Sprint(outputVal))
		if errVal == nil && errType == nil {
			(*outputPtr) = append(*outputPtr, outputType)
			return nil
		} else {
			outputDef, errDef := strconv.Atoi(fmt.Sprint(defaultValue))
			(*outputPtr) = append(*outputPtr, outputDef)
			if errVal == nil && errDef == nil {
				return errType
			} else {
				return fmt.Errorf("input value > %v %v ; default value > %v", errVal, errType, errDef)
			}
		}
	case *[]int8:
		outputPtr := outputPointer.(*[]int8)
		if errVal == nil {
			outputSliceType, ok := outputVal.([]int8)
			if ok {
				(*outputPtr) = append(*outputPtr, outputSliceType...)
				return nil
			}
		} else {
			outputSliceType, ok := defaultValue.([]int8)
			if ok {
				(*outputPtr) = append(*outputPtr, outputSliceType...)
				return errVal
			}
		}
		outputType, errType := strconv.ParseInt(fmt.Sprint(outputVal), 10, 8)
		if errVal == nil && errType == nil {
			(*outputPtr) = append(*outputPtr, int8(outputType))
			return nil
		} else {
			outputDef, errDef := strconv.ParseInt(fmt.Sprint(defaultValue), 10, 8)
			(*outputPtr) = append(*outputPtr, int8(outputDef))
			if errVal == nil && errDef == nil {
				return errType
			} else {
				return fmt.Errorf("input value > %v %v ; default value > %v", errVal, errType, errDef)
			}
		}
	case *[]int16:
		outputPtr := outputPointer.(*[]int16)
		if errVal == nil {
			outputSliceType, ok := outputVal.([]int16)
			if ok {
				(*outputPtr) = append(*outputPtr, outputSliceType...)
				return nil
			}
		} else {
			outputSliceType, ok := defaultValue.([]int16)
			if ok {
				(*outputPtr) = append(*outputPtr, outputSliceType...)
				return errVal
			}
		}
		outputType, errType := strconv.ParseInt(fmt.Sprint(outputVal), 10, 16)
		if errVal == nil && errType == nil {
			(*outputPtr) = append(*outputPtr, int16(outputType))
			return nil
		} else {
			outputDef, errDef := strconv.ParseInt(fmt.Sprint(defaultValue), 10, 16)
			(*outputPtr) = append(*outputPtr, int16(outputDef))
			if errVal == nil && errDef == nil {
				return errType
			} else {
				return fmt.Errorf("input value > %v %v ; default value > %v", errVal, errType, errDef)
			}
		}
	case *[]int32:
		outputPtr := outputPointer.(*[]int32)
		if errVal == nil {
			outputSliceType, ok := outputVal.([]int32)
			if ok {
				(*outputPtr) = append(*outputPtr, outputSliceType...)
				return nil
			}
		} else {
			outputSliceType, ok := defaultValue.([]int32)
			if ok {
				(*outputPtr) = append(*outputPtr, outputSliceType...)
				return errVal
			}
		}
		outputType, errType := strconv.ParseInt(fmt.Sprint(outputVal), 10, 32)
		if errVal == nil && errType == nil {
			(*outputPtr) = append(*outputPtr, int32(outputType))
			return nil
		} else {
			outputDef, errDef := strconv.ParseInt(fmt.Sprint(defaultValue), 10, 32)
			(*outputPtr) = append(*outputPtr, int32(outputDef))
			if errVal == nil && errDef == nil {
				return errType
			} else {
				return fmt.Errorf("input value > %v %v ; default value > %v", errVal, errType, errDef)
			}
		}
	case *[]int64:
		outputPtr := outputPointer.(*[]int64)
		if errVal == nil {
			outputSliceType, ok := outputVal.([]int64)
			if ok {
				(*outputPtr) = append(*outputPtr, outputSliceType...)
				return nil
			}
		} else {
			outputSliceType, ok := defaultValue.([]int64)
			if ok {
				(*outputPtr) = append(*outputPtr, outputSliceType...)
				return errVal
			}
		}
		outputType, errType := strconv.ParseInt(fmt.Sprint(outputVal), 10, 64)
		if errVal == nil && errType == nil {
			(*outputPtr) = append(*outputPtr, outputType)
			return nil
		} else {
			outputDef, errDef := strconv.ParseInt(fmt.Sprint(defaultValue), 10, 64)
			(*outputPtr) = append(*outputPtr, outputDef)
			if errVal == nil && errDef == nil {
				return errType
			} else {
				return fmt.Errorf("input value > %v %v ; default value > %v", errVal, errType, errDef)
			}
		}
	case *[]uint:
		outputPtr := outputPointer.(*[]uint)
		if errVal == nil {
			outputSliceType, ok := outputVal.([]uint)
			if ok {
				(*outputPtr) = append(*outputPtr, outputSliceType...)
				return nil
			}
		} else {
			outputSliceType, ok := defaultValue.([]uint)
			if ok {
				(*outputPtr) = append(*outputPtr, outputSliceType...)
				return errVal
			}
		}
		outputType, errType := strconv.ParseUint(fmt.Sprint(outputVal), 10, 32)
		if errVal == nil && errType == nil {
			(*outputPtr) = append(*outputPtr, uint(outputType))
			return nil
		} else {
			outputDef, errDef := strconv.ParseUint(fmt.Sprint(defaultValue), 10, 32)
			(*outputPtr) = append(*outputPtr, uint(outputDef))
			if errVal == nil && errDef == nil {
				return errType
			} else {
				return fmt.Errorf("input value > %v %v ; default value > %v", errVal, errType, errDef)
			}
		}
	case *[]uint8:
		outputPtr := outputPointer.(*[]uint8)
		if errVal == nil {
			outputSliceType, ok := outputVal.([]uint8)
			if ok {
				(*outputPtr) = append(*outputPtr, outputSliceType...)
				return nil
			}
		} else {
			outputSliceType, ok := defaultValue.([]uint8)
			if ok {
				(*outputPtr) = append(*outputPtr, outputSliceType...)
				return errVal
			}
		}
		outputType, errType := strconv.ParseUint(fmt.Sprint(outputVal), 10, 8)
		if errVal == nil && errType == nil {
			(*outputPtr) = append(*outputPtr, uint8(outputType))
			return nil
		} else {
			outputDef, errDef := strconv.ParseUint(fmt.Sprint(defaultValue), 10, 8)
			(*outputPtr) = append(*outputPtr, uint8(outputDef))
			if errVal == nil && errDef == nil {
				return errType
			} else {
				return fmt.Errorf("input value > %v %v ; default value > %v", errVal, errType, errDef)
			}
		}
	case *[]uint16:
		outputPtr := outputPointer.(*[]uint16)
		if errVal == nil {
			outputSliceType, ok := outputVal.([]uint16)
			if ok {
				(*outputPtr) = append(*outputPtr, outputSliceType...)
				return nil
			}
		} else {
			outputSliceType, ok := defaultValue.([]uint16)
			if ok {
				(*outputPtr) = append(*outputPtr, outputSliceType...)
				return errVal
			}
		}
		outputType, errType := strconv.ParseUint(fmt.Sprint(outputVal), 10, 16)
		if errVal == nil && errType == nil {
			(*outputPtr) = append(*outputPtr, uint16(outputType))
			return nil
		} else {
			outputDef, errDef := strconv.ParseUint(fmt.Sprint(defaultValue), 10, 16)
			(*outputPtr) = append(*outputPtr, uint16(outputDef))
			if errVal == nil && errDef == nil {
				return errType
			} else {
				return fmt.Errorf("input value > %v %v ; default value > %v", errVal, errType, errDef)
			}
		}
	case *[]uint32:
		outputPtr := outputPointer.(*[]uint32)
		if errVal == nil {
			outputSliceType, ok := outputVal.([]uint32)
			if ok {
				(*outputPtr) = append(*outputPtr, outputSliceType...)
				return nil
			}
		} else {
			outputSliceType, ok := defaultValue.([]uint32)
			if ok {
				(*outputPtr) = append(*outputPtr, outputSliceType...)
				return errVal
			}
		}
		outputType, errType := strconv.ParseUint(fmt.Sprint(outputVal), 10, 32)
		if errVal == nil && errType == nil {
			(*outputPtr) = append(*outputPtr, uint32(outputType))
			return nil
		} else {
			outputDef, errDef := strconv.ParseUint(fmt.Sprint(defaultValue), 10, 32)
			(*outputPtr) = append(*outputPtr, uint32(outputDef))
			if errVal == nil && errDef == nil {
				return errType
			} else {
				return fmt.Errorf("input value > %v %v ; default value > %v", errVal, errType, errDef)
			}
		}
	case *[]uint64:
		outputPtr := outputPointer.(*[]uint64)
		if errVal == nil {
			outputSliceType, ok := outputVal.([]uint64)
			if ok {
				(*outputPtr) = append(*outputPtr, outputSliceType...)
				return nil
			}
		} else {
			outputSliceType, ok := defaultValue.([]uint64)
			if ok {
				(*outputPtr) = append(*outputPtr, outputSliceType...)
				return errVal
			}
		}
		outputType, errType := strconv.ParseUint(fmt.Sprint(outputVal), 10, 64)
		if errVal == nil && errType == nil {
			(*outputPtr) = append(*outputPtr, outputType)
			return nil
		} else {
			outputDef, errDef := strconv.ParseUint(fmt.Sprint(defaultValue), 10, 64)
			(*outputPtr) = append(*outputPtr, outputDef)
			if errVal == nil && errDef == nil {
				return errType
			} else {
				return fmt.Errorf("input value > %v %v ; default value > %v", errVal, errType, errDef)
			}
		}

	case *[]float32:
		outputPtr := outputPointer.(*[]float32)
		if errVal == nil {
			outputSliceType, ok := outputVal.([]float32)
			if ok {
				(*outputPtr) = append(*outputPtr, outputSliceType...)
				return nil
			}
		} else {
			outputSliceType, ok := defaultValue.([]float32)
			if ok {
				(*outputPtr) = append(*outputPtr, outputSliceType...)
				return errVal
			}
		}
		outputType, errType := strconv.ParseFloat(fmt.Sprint(outputVal), 32)
		if errVal == nil && errType == nil {
			(*outputPtr) = append(*outputPtr, float32(outputType))
			return nil
		} else {
			outputDef, errDef := strconv.ParseFloat(fmt.Sprint(defaultValue), 32)
			(*outputPtr) = append(*outputPtr, float32(outputDef))
			if errVal == nil && errDef == nil {
				return errType
			} else {
				return fmt.Errorf("input value > %v %v ; default value > %v", errVal, errType, errDef)
			}
		}
	case *[]float64:
		outputPtr := outputPointer.(*[]float64)
		if errVal == nil {
			outputSliceType, ok := outputVal.([]float64)
			if ok {
				(*outputPtr) = append(*outputPtr, outputSliceType...)
				return nil
			}
		} else {
			outputSliceType, ok := defaultValue.([]float64)
			if ok {
				(*outputPtr) = append(*outputPtr, outputSliceType...)
				return errVal
			}
		}
		outputType, errType := strconv.ParseFloat(fmt.Sprint(outputVal), 64)
		if errVal == nil && errType == nil {
			(*outputPtr) = append(*outputPtr, outputType)
			return nil
		} else {
			outputDef, errDef := strconv.ParseFloat(fmt.Sprint(defaultValue), 64)
			(*outputPtr) = append(*outputPtr, outputDef)
			if errVal == nil && errDef == nil {
				return errType
			} else {
				return fmt.Errorf("input value > %v %v ; default value > %v", errVal, errType, errDef)
			}
		}
	case *[]string:
		outputPtr := outputPointer.(*[]string)
		if errVal == nil {
			outputType, ok := outputVal.(string)
			if ok {
				(*outputPtr) = append(*outputPtr, outputType)
				return nil
			}
			outputSliceType, ok := outputVal.([]string)
			if ok {
				(*outputPtr) = append(*outputPtr, outputSliceType...)
				return nil
			}
			return fmt.Errorf("input value error type > %v (%T) ;", outputVal, outputVal)
		} else {
			outputStringDef, ok := defaultValue.(string)
			if ok {
				(*outputPtr) = append(*outputPtr, outputStringDef)
				return errVal
			}
			outputSliceType, ok := defaultValue.([]string)
			if ok {
				(*outputPtr) = append(*outputPtr, outputSliceType...)
				return errVal
			}
			return fmt.Errorf("input value > %v ; default error type > %v (%T)", errVal, defaultValue, defaultValue)
		}
	case *[]bool:
		outputPtr := outputPointer.(*[]bool)
		if errVal == nil {
			outputType, ok := outputVal.(bool)
			if ok {
				(*outputPtr) = append(*outputPtr, outputType)
				return nil
			}
			outputSliceType, ok := outputVal.([]bool)
			if ok {
				(*outputPtr) = append(*outputPtr, outputSliceType...)
				return nil
			}

			return fmt.Errorf("input value error type > %v (%T) ;", outputVal, outputVal)
		} else {
			outputStringDef, ok := defaultValue.(bool)
			if ok {
				(*outputPtr) = append(*outputPtr, outputStringDef)
				return errVal
			}
			outputSliceType, ok := defaultValue.([]bool)
			if ok {
				(*outputPtr) = append(*outputPtr, outputSliceType...)
				return errVal
			}
			return fmt.Errorf("input value > %v ; default error type > %v (%T)", errVal, defaultValue, defaultValue)
		}
	case *[]interface{}:
		outputPtr := outputPointer.(*[]interface{})
		if errVal == nil {
			outputType, ok := outputVal.(interface{})
			if ok {
				(*outputPtr) = append(*outputPtr, outputType)
				return nil
			}
			outputSliceType, ok := outputVal.([]interface{})
			if ok {
				(*outputPtr) = append(*outputPtr, outputSliceType...)
				return nil
			}
			return fmt.Errorf("input value error type > %v (%T) ;", outputVal, outputVal)
		} else {
			outputType, ok := defaultValue.(interface{})
			if ok {
				(*outputPtr) = append(*outputPtr, outputType)
				return errVal
			}
			outputSliceType, ok := defaultValue.([]interface{})
			if ok {
				(*outputPtr) = append(*outputPtr, outputSliceType...)
				return errVal
			}
			return fmt.Errorf("input value > %v ; default error type > %v (%T)", errVal, defaultValue, defaultValue)
		}

	case *map[string]interface{}:
		var ok bool
		outputPtr := outputPointer.(*map[string]interface{})
		if len(*outputPtr) <= 0 {
			(*outputPtr) = make(map[string]interface{})
			if errVal == nil {
				*outputPtr, ok = outputVal.(map[string]interface{})
				if ok {
					return nil
				} else {
					return fmt.Errorf("程序輸出變數 %v(%T).GetValue() 輸出應為(map[string]interface{},err)，不准許(%T)", (*outputPtr), Custom, outputVal)
				}
			} else {
				*outputPtr, ok = defaultValue.(map[string]interface{})
				if ok {
					return errVal
				} else {
					return fmt.Errorf("程序輸出變數之defaultValue預設值 %v(%T).GetValue() 輸出應為(map[string]interface{},err)，不准許(%T)", (*outputPtr), Custom, defaultValue)
				}
			}

		} else {
			var cache map[string]interface{}
			if errVal == nil {
				cache, ok = outputVal.(map[string]interface{})
				if ok {
					for index, value := range cache {
						(*outputPtr)[index] = value
					}
					return nil
				} else {
					return fmt.Errorf("程序輸出變數 (%T).GetValue() 輸出應為(map[string]interface{},err)，不准許(%T)", Custom, outputVal)
				}
			} else {
				cache, ok = defaultValue.(map[string]interface{})
				if ok {
					for index, value := range cache {
						(*outputPtr)[index] = value
					}
					return errVal
				} else {
					return fmt.Errorf("程序輸出變數之defaultValue預設值 (%T).GetValue() 輸出應為(map[string]interface{},err)，不准許(%T)", Custom, defaultValue)
				}
			}
		}

	case *interface{}:
		outputPtr := outputPointer.(*interface{})
		if errVal == nil {
			(*outputPtr) = outputVal
			return nil
		} else {
			(*outputPtr) = defaultValue
			return errVal
		}

	default:
		return fmt.Errorf("程序未設置 outputPointer > '%v' (%v)型別的格式判斷", outputPointer, outputPointer)
	}
}

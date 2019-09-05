package array

import "reflect"

func InArray(val interface{}, array interface{}) (exists bool, index int) {
	exists = false
	index = -1

	switch reflect.TypeOf(array).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(array)

		for i := 0; i < s.Len(); i++ {
			if reflect.DeepEqual(val, s.Index(i).Interface()) == true {
				index = i
				exists = true
				return
			}
		}
	}

	return
}

func Remove(array interface{}, value interface{}) interface{} {
	switch reflect.TypeOf(array).Kind() {
	case reflect.Slice:
		isExist, index := InArray(value, array)
		if isExist {
			switch reflect.TypeOf(reflect.ValueOf(array).Index(0).Interface()).Kind() {
			case reflect.Bool:
				array = append(array.([]bool)[:index], array.([]bool)[(index+1):]...)
			case reflect.Int:
				array = append(array.([]int)[:index], array.([]int)[(index+1):]...)
			case reflect.Int8:
				array = append(array.([]int8)[:index], array.([]int8)[(index+1):]...)
			case reflect.Int16:
				array = append(array.([]int16)[:index], array.([]int16)[(index+1):]...)
			case reflect.Int32:
				array = append(array.([]int32)[:index], array.([]int32)[(index+1):]...)
			case reflect.Int64:
				array = append(array.([]int64)[:index], array.([]int64)[(index+1):]...)
			case reflect.Uint:
				array = append(array.([]uint)[:index], array.([]uint)[(index+1):]...)
			case reflect.Uint8:
				array = append(array.([]uint8)[:index], array.([]uint8)[(index+1):]...)
			case reflect.Uint16:
				array = append(array.([]uint16)[:index], array.([]uint16)[(index+1):]...)
			case reflect.Uint32:
				array = append(array.([]uint32)[:index], array.([]uint32)[(index+1):]...)
			case reflect.Uint64:
				array = append(array.([]uint64)[:index], array.([]uint64)[(index+1):]...)
			case reflect.Float32:
				array = append(array.([]float32)[:index], array.([]float32)[(index+1):]...)
			case reflect.Float64:
				array = append(array.([]float64)[:index], array.([]float64)[(index+1):]...)
			case reflect.String:
				array = append(array.([]string)[:index], array.([]string)[(index+1):]...)

			}

		}
	}

	return array
}

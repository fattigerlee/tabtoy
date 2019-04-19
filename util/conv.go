package util

import (
	"strconv"
	"strings"
)

func StringToPrimitive(str string, value interface{}) (error, bool) {
	switch raw := value.(type) {
	case *int32:
		v, err := strconv.ParseInt(str, 10, 32)
		if err != nil {
			return err, false
		}

		*raw = int32(v)
	case *int64:
		v, err := strconv.ParseInt(str, 10, 64)
		if err != nil {
			return err, false
		}

		*raw = v
	case *uint32:
		v, err := strconv.ParseUint(str, 10, 32)
		if err != nil {
			return err, false
		}

		*raw = uint32(v)
	case *uint64:
		v, err := strconv.ParseUint(str, 10, 64)
		if err != nil {
			return err, false
		}

		*raw = v
	case *string:
		*raw = str
	case *bool:

		var v bool
		var err error

		switch str {
		case "是":
			v = true
		case "否", "":
			v = false
		default:
			v, err = strconv.ParseBool(str)
			if err != nil {
				return err, false
			}
		}

		*raw = v
	case *float32:
		v, err := strconv.ParseFloat(str, 32)
		if err != nil {
			return err, false
		}

		*raw = float32(v)
	case *float64:
		v, err := strconv.ParseFloat(str, 64)
		if err != nil {
			return err, false
		}

		*raw = float64(v)

	default:
		return nil, false
	}

	return nil, true
}

// 下划线转驼峰
func UnderlineToCamel(str string) string {
	data := make([]byte, 0, len(str))
	size := len(str)
	for i := 0; i < size; i++ {
		d := str[i]
		if i == 0 && d >= 'a' && d <= 'z' {
			d = d - 32
		}

		if d == '_' && i+1 < size {
			i = i + 1
			d = str[i]
			if d >= 'a' && d <= 'z' {
				d = str[i] - 32
			}
		}

		if d != '_' {
			data = append(data, d)
		}
	}
	return string(data[:])
}

// 驼峰转下划线
func CamelToUnderline(str string) string {
	data := make([]byte, 0, len(str)*2)
	size := len(str)

	// 全大写,直接转小写
	upper := true
	for i := 0; i < size; i++ {
		if str[i] < 'A' || str[i] > 'Z' {
			upper = false
			break
		}
	}

	if upper {
		return strings.ToLower(str)
	}

	for i := 0; i < size-1; i++ {
		if str[i] >= 'A' && str[i] <= 'Z' &&
			str[i+1] >= 'A' && str[i+1] <= 'Z' {
			data = append(data, str[i])
			continue
		}

		if i > 0 && str[i] >= 'A' && str[i] <= 'Z' {
			data = append(data, '_', str[i])
			continue
		}
		data = append(data, str[i])
	}
	data = append(data, str[size-1])

	return strings.ToLower(string(data[:]))
}
